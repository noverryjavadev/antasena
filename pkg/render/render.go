package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// RenderTemplate renders a template
func RenderTemplate(w http.ResponseWriter, tmpl string) {

	// create template cache
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatalln(err)
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	// render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

	//parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl, "./templates/base.layout.tmpl")
	//err := parsedTemplate.Execute(w, nil)
	//if err != nil {
	//	fmt.Println("error parsing template:", err)
	//}
}

//var tc = make(map[string]*template.Template)
//
//func RenderTemplate(w http.ResponseWriter, t string) {
//	var tmpl *template.Template
//	var err error
//
//	//	 check to see if we already have a template in our cache
//	_, inMap := tc[t]
//	if !inMap {
//		// need to create the template
//		log.Println("creating template and adding to cache")
//		err = createTemplateCache(t)
//		if err != nil {
//			log.Println("error creating template cache:", err)
//			return
//		}
//	} else {
//		log.Println("using cached template")
//	}
//
//	tmpl = tc[t]
//
//	err = tmpl.Execute(w, nil)
//}
//
//func createTemplateCache(t string) error {
//	templates := []string{
//		"./templates/" + t, "./templates/base.layout.tmpl",
//	}
//
//	tmp, err := template.ParseFiles(templates...)
//	if err != nil {
//		return err
//	}
//
//	tc[t] = tmp
//	return nil
//}

func createTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Parse(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

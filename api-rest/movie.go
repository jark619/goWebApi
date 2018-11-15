package main

type Movie struct{
	Name string     `json:"nombre"`
	Year int        `json:"anio"`
	Director string `json:"director"`
}

type Movies []Movie
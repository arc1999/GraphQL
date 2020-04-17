package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"net/http"
)
type Cab struct {
	Cab_id int `json:"cab_id"`
	Cab_location string `json:"cab_location"`
	Cab_lat float64 `json:"cab_lat"`
	Cab_long float64 `json:"cab_long"`
}
var Cabs =[]Cab{
	{
		Cab_id:1,
		Cab_location: "Wakad",
		Cab_lat:5.12,
		Cab_long:8.01,
	},
	{
		Cab_id:2,
		Cab_location: "Ballewadi",
		Cab_lat:15.12,
		Cab_long:38.01,
	},
	{
		Cab_id:3,
		Cab_location: "Ballewadi",
		Cab_lat:34.3,
		Cab_long:50.40,
	},
	{
		Cab_id:4,
		Cab_location: "Ballewadi",
		Cab_lat:23.42,
		Cab_long:87.51,
	},

}
var cabType = graphql.NewObject(
	graphql.ObjectConfig{
		Name:        "Cab",
		Fields:      graphql.Fields{
			"Cab_id" : &graphql.Field{
							Type:graphql.Int,
						},
			"Cab_location": &graphql.Field{
							Type:graphql.String,
							},
			"Cab_lat":&graphql.Field{
							Type:graphql.Float,
							},
			"Cab_long": &graphql.Field{
							Type:graphql.Float,
						},
		},

	},
)
var queryType=graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields:graphql.Fields{

					"cab":&graphql.Field{
						Type:cabType,
						Description:"Fetch cab with location",
						Args:graphql.FieldConfigArgument{
							"Cab_location": &graphql.ArgumentConfig{
								Type:      graphql.String,
							},
						},
						Resolve: func(p graphql.ResolveParams) ( interface{},  error) {
							location,ok:=p.Args["Cab_location"].(string)
							if ok {
								var loc_cabs []Cab
								for _,cab := range Cabs{
									if string(cab.Cab_location)==location{
										loc_cabs=append(loc_cabs,cab)
									}
								}
								return loc_cabs ,nil
							}
							return nil,nil
						},
					},
					"list": &graphql.Field{
						Type:              graphql.NewList(cabType),
						Description:       "Get all cabs",
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return Cabs,nil
						},
					},

		},

	},)
var schema,_ =graphql.NewSchema(
	 graphql.SchemaConfig{
	 	Query:queryType,
	 	},
)
func executeQuery(query string ,schema graphql.Schema) *graphql.Result{
	result:= graphql.Do(graphql.Params{
		Schema:         schema,
		RequestString: query,

	})
	if len(result.Errors)> 0{
		fmt.Printf("errors:%v",result.Errors)
	}
	return result
}
func main() {
	http.HandleFunc("/cab", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

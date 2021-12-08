package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	Information string             `json:"info,omitempty" bson:"info,omitempty"`
	ListOfInfo  []ListOfInfo       `json:"listInfo,omitempty" bson:"listInfo,omitempty"`
}
type ListOfInfo struct {
	Info string `json:"info,omitempty" bson:"info,omitempty"`
	Desc string `json:"desc,omitempty" bson:"desc,omitempty"`
}

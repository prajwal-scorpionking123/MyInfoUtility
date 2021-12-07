package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title       string             `json:"title,omitempty" bson:"title,omitempty"`
	Category    string             `json:"category,omitempty" bson:"category,omitempty"`
	Type        string             `json:"type,omitempty" bson:"type,omitempty"`
	Information string             `json:"info,omitempty" bson:"info,omitempty"`
}

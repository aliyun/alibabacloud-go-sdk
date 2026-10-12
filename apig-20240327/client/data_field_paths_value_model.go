// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataFieldPathsValue interface {
	dara.Model
	String() string
	GoString() string
	SetFieldKey(v string) *DataFieldPathsValue
	GetFieldKey() *string
	SetJsonPath(v string) *DataFieldPathsValue
	GetJsonPath() *string
	SetSource(v string) *DataFieldPathsValue
	GetSource() *string
	SetCategory(v string) *DataFieldPathsValue
	GetCategory() *string
}

type DataFieldPathsValue struct {
	// The identifier key of the template field, such as input_token.
	//
	// example:
	//
	// input_token
	FieldKey *string `json:"fieldKey,omitempty" xml:"fieldKey,omitempty"`
	// The JSONPath expression of the template field in the request or response message, such as usage.prompt_tokens.
	//
	// example:
	//
	// usage.prompt_tokens
	JsonPath *string `json:"jsonPath,omitempty" xml:"jsonPath,omitempty"`
	// The value source of the field, which indicates the part of the message from which the field is read, such as response_body.
	//
	// example:
	//
	// response_body
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// The statistical category to which the field belongs.
	//
	// example:
	//
	// usage
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
}

func (s DataFieldPathsValue) String() string {
	return dara.Prettify(s)
}

func (s DataFieldPathsValue) GoString() string {
	return s.String()
}

func (s *DataFieldPathsValue) GetFieldKey() *string {
	return s.FieldKey
}

func (s *DataFieldPathsValue) GetJsonPath() *string {
	return s.JsonPath
}

func (s *DataFieldPathsValue) GetSource() *string {
	return s.Source
}

func (s *DataFieldPathsValue) GetCategory() *string {
	return s.Category
}

func (s *DataFieldPathsValue) SetFieldKey(v string) *DataFieldPathsValue {
	s.FieldKey = &v
	return s
}

func (s *DataFieldPathsValue) SetJsonPath(v string) *DataFieldPathsValue {
	s.JsonPath = &v
	return s
}

func (s *DataFieldPathsValue) SetSource(v string) *DataFieldPathsValue {
	s.Source = &v
	return s
}

func (s *DataFieldPathsValue) SetCategory(v string) *DataFieldPathsValue {
	s.Category = &v
	return s
}

func (s *DataFieldPathsValue) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataResourceValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DataResourceValue
	GetName() *string
	SetType(v string) *DataResourceValue
	GetType() *string
	SetContent(v string) *DataResourceValue
	GetContent() *string
	SetMetadata(v map[string]interface{}) *DataResourceValue
	GetMetadata() map[string]interface{}
}

type DataResourceValue struct {
	// The name.
	//
	// example:
	//
	// agentspec-example
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The type.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The content.
	//
	// example:
	//
	// Sample content
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// The metadata.
	Metadata map[string]interface{} `json:"metadata,omitempty" xml:"metadata,omitempty"`
}

func (s DataResourceValue) String() string {
	return dara.Prettify(s)
}

func (s DataResourceValue) GoString() string {
	return s.String()
}

func (s *DataResourceValue) GetName() *string {
	return s.Name
}

func (s *DataResourceValue) GetType() *string {
	return s.Type
}

func (s *DataResourceValue) GetContent() *string {
	return s.Content
}

func (s *DataResourceValue) GetMetadata() map[string]interface{} {
	return s.Metadata
}

func (s *DataResourceValue) SetName(v string) *DataResourceValue {
	s.Name = &v
	return s
}

func (s *DataResourceValue) SetType(v string) *DataResourceValue {
	s.Type = &v
	return s
}

func (s *DataResourceValue) SetContent(v string) *DataResourceValue {
	s.Content = &v
	return s
}

func (s *DataResourceValue) SetMetadata(v map[string]interface{}) *DataResourceValue {
	s.Metadata = v
	return s
}

func (s *DataResourceValue) Validate() error {
	return dara.Validate(s)
}

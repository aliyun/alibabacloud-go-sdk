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
	Name     *string                `json:"Name,omitempty" xml:"Name,omitempty"`
	Type     *string                `json:"Type,omitempty" xml:"Type,omitempty"`
	Content  *string                `json:"Content,omitempty" xml:"Content,omitempty"`
	Metadata map[string]interface{} `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSourcePluginsValue interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *DataSourcePluginsValue
	GetName() *string
	SetFromFields(v string) *DataSourcePluginsValue
	GetFromFields() *string
	SetParameters(v map[string]*string) *DataSourcePluginsValue
	GetParameters() map[string]*string
}

type DataSourcePluginsValue struct {
	Name       *string            `json:"name,omitempty" xml:"name,omitempty"`
	FromFields *string            `json:"fromFields,omitempty" xml:"fromFields,omitempty"`
	Parameters map[string]*string `json:"parameters,omitempty" xml:"parameters,omitempty"`
}

func (s DataSourcePluginsValue) String() string {
	return dara.Prettify(s)
}

func (s DataSourcePluginsValue) GoString() string {
	return s.String()
}

func (s *DataSourcePluginsValue) GetName() *string {
	return s.Name
}

func (s *DataSourcePluginsValue) GetFromFields() *string {
	return s.FromFields
}

func (s *DataSourcePluginsValue) GetParameters() map[string]*string {
	return s.Parameters
}

func (s *DataSourcePluginsValue) SetName(v string) *DataSourcePluginsValue {
	s.Name = &v
	return s
}

func (s *DataSourcePluginsValue) SetFromFields(v string) *DataSourcePluginsValue {
	s.FromFields = &v
	return s
}

func (s *DataSourcePluginsValue) SetParameters(v map[string]*string) *DataSourcePluginsValue {
	s.Parameters = v
	return s
}

func (s *DataSourcePluginsValue) Validate() error {
	return dara.Validate(s)
}

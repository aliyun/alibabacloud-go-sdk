// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataToolSpecToolsMetaValue interface {
	dara.Model
	String() string
	GoString() string
	SetEnabled(v bool) *DataToolSpecToolsMetaValue
	GetEnabled() *bool
	SetTemplates(v map[string]interface{}) *DataToolSpecToolsMetaValue
	GetTemplates() map[string]interface{}
}

type DataToolSpecToolsMetaValue struct {
	// example:
	//
	// true
	Enabled   *bool                  `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	Templates map[string]interface{} `json:"Templates,omitempty" xml:"Templates,omitempty"`
}

func (s DataToolSpecToolsMetaValue) String() string {
	return dara.Prettify(s)
}

func (s DataToolSpecToolsMetaValue) GoString() string {
	return s.String()
}

func (s *DataToolSpecToolsMetaValue) GetEnabled() *bool {
	return s.Enabled
}

func (s *DataToolSpecToolsMetaValue) GetTemplates() map[string]interface{} {
	return s.Templates
}

func (s *DataToolSpecToolsMetaValue) SetEnabled(v bool) *DataToolSpecToolsMetaValue {
	s.Enabled = &v
	return s
}

func (s *DataToolSpecToolsMetaValue) SetTemplates(v map[string]interface{}) *DataToolSpecToolsMetaValue {
	s.Templates = v
	return s
}

func (s *DataToolSpecToolsMetaValue) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iServiceChanges interface {
	dara.Model
	String() string
	GoString() string
	SetMerge(v map[string]interface{}) *ServiceChanges
	GetMerge() map[string]interface{}
}

type ServiceChanges struct {
	// example:
	//
	// {}: 不进行修改
	Merge map[string]interface{} `json:"merge,omitempty" xml:"merge,omitempty"`
}

func (s ServiceChanges) String() string {
	return dara.Prettify(s)
}

func (s ServiceChanges) GoString() string {
	return s.String()
}

func (s *ServiceChanges) GetMerge() map[string]interface{} {
	return s.Merge
}

func (s *ServiceChanges) SetMerge(v map[string]interface{}) *ServiceChanges {
	s.Merge = v
	return s
}

func (s *ServiceChanges) Validate() error {
	return dara.Validate(s)
}

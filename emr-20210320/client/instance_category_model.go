// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstanceCategory interface {
	dara.Model
	String() string
	GoString() string
	SetDefaultValue(v string) *InstanceCategory
	GetDefaultValue() *string
	SetKeys(v []*string) *InstanceCategory
	GetKeys() []*string
	SetValues(v []*string) *InstanceCategory
	GetValues() []*string
}

type InstanceCategory struct {
	// 默认值。
	//
	// example:
	//
	// CLUSTER
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// example:
	//
	// null
	Keys []*string `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// example:
	//
	// null
	Values []*string `json:"Values,omitempty" xml:"Values,omitempty" type:"Repeated"`
}

func (s InstanceCategory) String() string {
	return dara.Prettify(s)
}

func (s InstanceCategory) GoString() string {
	return s.String()
}

func (s *InstanceCategory) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *InstanceCategory) GetKeys() []*string {
	return s.Keys
}

func (s *InstanceCategory) GetValues() []*string {
	return s.Values
}

func (s *InstanceCategory) SetDefaultValue(v string) *InstanceCategory {
	s.DefaultValue = &v
	return s
}

func (s *InstanceCategory) SetKeys(v []*string) *InstanceCategory {
	s.Keys = v
	return s
}

func (s *InstanceCategory) SetValues(v []*string) *InstanceCategory {
	s.Values = v
	return s
}

func (s *InstanceCategory) Validate() error {
	return dara.Validate(s)
}

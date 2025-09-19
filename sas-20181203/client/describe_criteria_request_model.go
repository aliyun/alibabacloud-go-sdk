// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCriteriaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v string) *DescribeCriteriaRequest
	GetMachineTypes() *string
	SetSupportAutoTag(v bool) *DescribeCriteriaRequest
	GetSupportAutoTag() *bool
	SetValue(v string) *DescribeCriteriaRequest
	GetValue() *string
}

type DescribeCriteriaRequest struct {
	// The type of the asset. Valid values:
	//
	// 	- Set the value to **ecs**, which specifies to query all Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// ecs
	MachineTypes *string `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty"`
	// Specifies whether the keyword that you specify for fuzzy search can be automatically matched. Default value: **false**. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	SupportAutoTag *bool `json:"SupportAutoTag,omitempty" xml:"SupportAutoTag,omitempty"`
	// The keyword that you specify for fuzzy search when you query the asset.
	//
	// example:
	//
	// 47.96
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeCriteriaRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCriteriaRequest) GoString() string {
	return s.String()
}

func (s *DescribeCriteriaRequest) GetMachineTypes() *string {
	return s.MachineTypes
}

func (s *DescribeCriteriaRequest) GetSupportAutoTag() *bool {
	return s.SupportAutoTag
}

func (s *DescribeCriteriaRequest) GetValue() *string {
	return s.Value
}

func (s *DescribeCriteriaRequest) SetMachineTypes(v string) *DescribeCriteriaRequest {
	s.MachineTypes = &v
	return s
}

func (s *DescribeCriteriaRequest) SetSupportAutoTag(v bool) *DescribeCriteriaRequest {
	s.SupportAutoTag = &v
	return s
}

func (s *DescribeCriteriaRequest) SetValue(v string) *DescribeCriteriaRequest {
	s.Value = &v
	return s
}

func (s *DescribeCriteriaRequest) Validate() error {
	return dara.Validate(s)
}

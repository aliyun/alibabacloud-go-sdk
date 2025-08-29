// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateResourceRuleItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *UpdateResourceRuleItemRequest
	GetDescription() *string
	SetInstanceId(v string) *UpdateResourceRuleItemRequest
	GetInstanceId() *string
	SetMaxValue(v float64) *UpdateResourceRuleItemRequest
	GetMaxValue() *float64
	SetMinValue(v float64) *UpdateResourceRuleItemRequest
	GetMinValue() *float64
	SetName(v string) *UpdateResourceRuleItemRequest
	GetName() *string
	SetValue(v float64) *UpdateResourceRuleItemRequest
	GetValue() *float64
}

type UpdateResourceRuleItemRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string  `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaxValue   *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	MinValue   *float64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// This parameter is required.
	Name  *string  `json:"Name,omitempty" xml:"Name,omitempty"`
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s UpdateResourceRuleItemRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateResourceRuleItemRequest) GoString() string {
	return s.String()
}

func (s *UpdateResourceRuleItemRequest) GetDescription() *string {
	return s.Description
}

func (s *UpdateResourceRuleItemRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *UpdateResourceRuleItemRequest) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *UpdateResourceRuleItemRequest) GetMinValue() *float64 {
	return s.MinValue
}

func (s *UpdateResourceRuleItemRequest) GetName() *string {
	return s.Name
}

func (s *UpdateResourceRuleItemRequest) GetValue() *float64 {
	return s.Value
}

func (s *UpdateResourceRuleItemRequest) SetDescription(v string) *UpdateResourceRuleItemRequest {
	s.Description = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetInstanceId(v string) *UpdateResourceRuleItemRequest {
	s.InstanceId = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetMaxValue(v float64) *UpdateResourceRuleItemRequest {
	s.MaxValue = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetMinValue(v float64) *UpdateResourceRuleItemRequest {
	s.MinValue = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetName(v string) *UpdateResourceRuleItemRequest {
	s.Name = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) SetValue(v float64) *UpdateResourceRuleItemRequest {
	s.Value = &v
	return s
}

func (s *UpdateResourceRuleItemRequest) Validate() error {
	return dara.Validate(s)
}

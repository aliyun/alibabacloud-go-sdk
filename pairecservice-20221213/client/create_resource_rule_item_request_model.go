// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateResourceRuleItemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateResourceRuleItemRequest
	GetDescription() *string
	SetInstanceId(v string) *CreateResourceRuleItemRequest
	GetInstanceId() *string
	SetMaxValue(v float64) *CreateResourceRuleItemRequest
	GetMaxValue() *float64
	SetMinValue(v float64) *CreateResourceRuleItemRequest
	GetMinValue() *float64
	SetName(v string) *CreateResourceRuleItemRequest
	GetName() *string
	SetValue(v float64) *CreateResourceRuleItemRequest
	GetValue() *float64
}

type CreateResourceRuleItemRequest struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	MaxValue *float64 `json:"MaxValue,omitempty" xml:"MaxValue,omitempty"`
	// This parameter is required.
	MinValue *float64 `json:"MinValue,omitempty" xml:"MinValue,omitempty"`
	// This parameter is required.
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// This parameter is required.
	Value *float64 `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateResourceRuleItemRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateResourceRuleItemRequest) GoString() string {
	return s.String()
}

func (s *CreateResourceRuleItemRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateResourceRuleItemRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateResourceRuleItemRequest) GetMaxValue() *float64 {
	return s.MaxValue
}

func (s *CreateResourceRuleItemRequest) GetMinValue() *float64 {
	return s.MinValue
}

func (s *CreateResourceRuleItemRequest) GetName() *string {
	return s.Name
}

func (s *CreateResourceRuleItemRequest) GetValue() *float64 {
	return s.Value
}

func (s *CreateResourceRuleItemRequest) SetDescription(v string) *CreateResourceRuleItemRequest {
	s.Description = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetInstanceId(v string) *CreateResourceRuleItemRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetMaxValue(v float64) *CreateResourceRuleItemRequest {
	s.MaxValue = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetMinValue(v float64) *CreateResourceRuleItemRequest {
	s.MinValue = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetName(v string) *CreateResourceRuleItemRequest {
	s.Name = &v
	return s
}

func (s *CreateResourceRuleItemRequest) SetValue(v float64) *CreateResourceRuleItemRequest {
	s.Value = &v
	return s
}

func (s *CreateResourceRuleItemRequest) Validate() error {
	return dara.Validate(s)
}

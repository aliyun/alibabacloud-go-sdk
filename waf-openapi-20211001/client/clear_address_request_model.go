// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iClearAddressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ClearAddressRequest
	GetInstanceId() *string
	SetResourceManagerResourceGroupId(v string) *ClearAddressRequest
	GetResourceManagerResourceGroupId() *string
	SetRuleId(v int64) *ClearAddressRequest
	GetRuleId() *int64
}

type ClearAddressRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// waf-cn-7mz****hw0u
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// rg-acfm***q
	ResourceManagerResourceGroupId *string `json:"ResourceManagerResourceGroupId,omitempty" xml:"ResourceManagerResourceGroupId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 12345678
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
}

func (s ClearAddressRequest) String() string {
	return dara.Prettify(s)
}

func (s ClearAddressRequest) GoString() string {
	return s.String()
}

func (s *ClearAddressRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ClearAddressRequest) GetResourceManagerResourceGroupId() *string {
	return s.ResourceManagerResourceGroupId
}

func (s *ClearAddressRequest) GetRuleId() *int64 {
	return s.RuleId
}

func (s *ClearAddressRequest) SetInstanceId(v string) *ClearAddressRequest {
	s.InstanceId = &v
	return s
}

func (s *ClearAddressRequest) SetResourceManagerResourceGroupId(v string) *ClearAddressRequest {
	s.ResourceManagerResourceGroupId = &v
	return s
}

func (s *ClearAddressRequest) SetRuleId(v int64) *ClearAddressRequest {
	s.RuleId = &v
	return s
}

func (s *ClearAddressRequest) Validate() error {
	return dara.Validate(s)
}

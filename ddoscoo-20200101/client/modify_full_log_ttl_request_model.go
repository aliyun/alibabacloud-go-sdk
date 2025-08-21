// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFullLogTtlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *ModifyFullLogTtlRequest
	GetResourceGroupId() *string
	SetTtl(v int32) *ModifyFullLogTtlRequest
	GetTtl() *int32
}

type ModifyFullLogTtlRequest struct {
	// The ID of the resource group to which the instance belongs in Resource Management. This parameter is empty by default, which indicates that the instance belongs to the default resource group.
	//
	// example:
	//
	// default
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The log storage duration of a website. Valid values: **7*	- to **180**. Unit: days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30
	Ttl *int32 `json:"Ttl,omitempty" xml:"Ttl,omitempty"`
}

func (s ModifyFullLogTtlRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFullLogTtlRequest) GoString() string {
	return s.String()
}

func (s *ModifyFullLogTtlRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyFullLogTtlRequest) GetTtl() *int32 {
	return s.Ttl
}

func (s *ModifyFullLogTtlRequest) SetResourceGroupId(v string) *ModifyFullLogTtlRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyFullLogTtlRequest) SetTtl(v int32) *ModifyFullLogTtlRequest {
	s.Ttl = &v
	return s
}

func (s *ModifyFullLogTtlRequest) Validate() error {
	return dara.Validate(s)
}

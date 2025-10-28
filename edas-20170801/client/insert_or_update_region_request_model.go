// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertOrUpdateRegionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDebugEnable(v bool) *InsertOrUpdateRegionRequest
	GetDebugEnable() *bool
	SetDescription(v string) *InsertOrUpdateRegionRequest
	GetDescription() *string
	SetId(v int64) *InsertOrUpdateRegionRequest
	GetId() *int64
	SetMseInstanceId(v string) *InsertOrUpdateRegionRequest
	GetMseInstanceId() *string
	SetRegionName(v string) *InsertOrUpdateRegionRequest
	GetRegionName() *string
	SetRegionTag(v string) *InsertOrUpdateRegionRequest
	GetRegionTag() *string
	SetRegistryType(v string) *InsertOrUpdateRegionRequest
	GetRegistryType() *string
}

type InsertOrUpdateRegionRequest struct {
	// Specifies whether to enable remote debugging. Valid values:
	//
	// 	- true: enables remote debugging.
	//
	// 	- false: disables remote debugging.
	//
	// example:
	//
	// true
	DebugEnable *bool `json:"DebugEnable,omitempty" xml:"DebugEnable,omitempty"`
	// The description of the namespace. The description can be up to 128 characters in length.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to create or modify the namespace. If this parameter is left empty or is set to 0, the namespace is created. Otherwise, the namespace is modified.
	//
	// example:
	//
	// 0
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the MSE registry.
	//
	// example:
	//
	// mse_prepaid_public_cn-tl32n******
	MseInstanceId *string `json:"MseInstanceId,omitempty" xml:"MseInstanceId,omitempty"`
	// The name of the namespace. The name can be up to 63 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_region
	RegionName *string `json:"RegionName,omitempty" xml:"RegionName,omitempty"`
	// The ID of the namespace.
	//
	// 	- The ID of a custom namespace is in the `Region ID:Namespace identifier` format. Example: cn-beijing:tdy218.
	//
	// 	- The ID of the default namespace is in the `region ID` format. Example: cn-beijing.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing:test
	RegionTag *string `json:"RegionTag,omitempty" xml:"RegionTag,omitempty"`
	// The type of the registry.
	//
	// 	- default: the shared registry of Enterprise Distributed Application Service (EDAS)
	//
	// 	- exclusive_mse: a Microservices Engine (MSE) registry
	//
	// example:
	//
	// default
	RegistryType *string `json:"RegistryType,omitempty" xml:"RegistryType,omitempty"`
}

func (s InsertOrUpdateRegionRequest) String() string {
	return dara.Prettify(s)
}

func (s InsertOrUpdateRegionRequest) GoString() string {
	return s.String()
}

func (s *InsertOrUpdateRegionRequest) GetDebugEnable() *bool {
	return s.DebugEnable
}

func (s *InsertOrUpdateRegionRequest) GetDescription() *string {
	return s.Description
}

func (s *InsertOrUpdateRegionRequest) GetId() *int64 {
	return s.Id
}

func (s *InsertOrUpdateRegionRequest) GetMseInstanceId() *string {
	return s.MseInstanceId
}

func (s *InsertOrUpdateRegionRequest) GetRegionName() *string {
	return s.RegionName
}

func (s *InsertOrUpdateRegionRequest) GetRegionTag() *string {
	return s.RegionTag
}

func (s *InsertOrUpdateRegionRequest) GetRegistryType() *string {
	return s.RegistryType
}

func (s *InsertOrUpdateRegionRequest) SetDebugEnable(v bool) *InsertOrUpdateRegionRequest {
	s.DebugEnable = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetDescription(v string) *InsertOrUpdateRegionRequest {
	s.Description = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetId(v int64) *InsertOrUpdateRegionRequest {
	s.Id = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetMseInstanceId(v string) *InsertOrUpdateRegionRequest {
	s.MseInstanceId = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetRegionName(v string) *InsertOrUpdateRegionRequest {
	s.RegionName = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetRegionTag(v string) *InsertOrUpdateRegionRequest {
	s.RegionTag = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) SetRegistryType(v string) *InsertOrUpdateRegionRequest {
	s.RegistryType = &v
	return s
}

func (s *InsertOrUpdateRegionRequest) Validate() error {
	return dara.Validate(s)
}

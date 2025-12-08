// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCenterPolicyListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcademicProxy(v string) *DescribeCenterPolicyListRequest
	GetAcademicProxy() *string
	SetBusinessType(v int32) *DescribeCenterPolicyListRequest
	GetBusinessType() *int32
	SetModelLibrary(v string) *DescribeCenterPolicyListRequest
	GetModelLibrary() *string
	SetPageNumber(v int32) *DescribeCenterPolicyListRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeCenterPolicyListRequest
	GetPageSize() *int32
	SetPolicyGroupId(v []*string) *DescribeCenterPolicyListRequest
	GetPolicyGroupId() []*string
	SetPortProxy(v string) *DescribeCenterPolicyListRequest
	GetPortProxy() *string
	SetResourceType(v string) *DescribeCenterPolicyListRequest
	GetResourceType() *string
	SetScope(v string) *DescribeCenterPolicyListRequest
	GetScope() *string
}

type DescribeCenterPolicyListRequest struct {
	AcademicProxy *string `json:"AcademicProxy,omitempty" xml:"AcademicProxy,omitempty"`
	// The business type.
	//
	// Valid values:
	//
	// 	- 1: public cloud
	//
	// 	- 8: commercial edition.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	BusinessType *int32  `json:"BusinessType,omitempty" xml:"BusinessType,omitempty"`
	ModelLibrary *string `json:"ModelLibrary,omitempty" xml:"ModelLibrary,omitempty"`
	// The page number.\\
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The IDs of the cloud computer policies.
	PolicyGroupId []*string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty" type:"Repeated"`
	PortProxy     *string   `json:"PortProxy,omitempty" xml:"PortProxy,omitempty"`
	// The resource type.
	//
	// Valid values:
	//
	// 	- app: cloud applications.
	//
	// 	- desktop: cloud computers.
	//
	// This parameter is required.
	//
	// example:
	//
	// desktop
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The effective scope of the cloud computer policy.
	//
	// Valid values:
	//
	// 	- IP: The policy applies to specific IP addresses.
	//
	// 	- GLOBAL: The policy applies globally.
	//
	// example:
	//
	// GLOBAL
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
}

func (s DescribeCenterPolicyListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCenterPolicyListRequest) GoString() string {
	return s.String()
}

func (s *DescribeCenterPolicyListRequest) GetAcademicProxy() *string {
	return s.AcademicProxy
}

func (s *DescribeCenterPolicyListRequest) GetBusinessType() *int32 {
	return s.BusinessType
}

func (s *DescribeCenterPolicyListRequest) GetModelLibrary() *string {
	return s.ModelLibrary
}

func (s *DescribeCenterPolicyListRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeCenterPolicyListRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeCenterPolicyListRequest) GetPolicyGroupId() []*string {
	return s.PolicyGroupId
}

func (s *DescribeCenterPolicyListRequest) GetPortProxy() *string {
	return s.PortProxy
}

func (s *DescribeCenterPolicyListRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeCenterPolicyListRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeCenterPolicyListRequest) SetAcademicProxy(v string) *DescribeCenterPolicyListRequest {
	s.AcademicProxy = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetBusinessType(v int32) *DescribeCenterPolicyListRequest {
	s.BusinessType = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetModelLibrary(v string) *DescribeCenterPolicyListRequest {
	s.ModelLibrary = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetPageNumber(v int32) *DescribeCenterPolicyListRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetPageSize(v int32) *DescribeCenterPolicyListRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetPolicyGroupId(v []*string) *DescribeCenterPolicyListRequest {
	s.PolicyGroupId = v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetPortProxy(v string) *DescribeCenterPolicyListRequest {
	s.PortProxy = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetResourceType(v string) *DescribeCenterPolicyListRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) SetScope(v string) *DescribeCenterPolicyListRequest {
	s.Scope = &v
	return s
}

func (s *DescribeCenterPolicyListRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeServiceLinkedRoleStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeServiceLinkedRoleStatusRequest
	GetRegionId() *string
	SetSdkRequest(v *DescribeServiceLinkedRoleStatusRequestSdkRequest) *DescribeServiceLinkedRoleStatusRequest
	GetSdkRequest() *DescribeServiceLinkedRoleStatusRequestSdkRequest
}

type DescribeServiceLinkedRoleStatusRequest struct {
	// The ID of the region in which the instance resides.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The Security Center SDK request.
	SdkRequest *DescribeServiceLinkedRoleStatusRequestSdkRequest `json:"SdkRequest,omitempty" xml:"SdkRequest,omitempty" type:"Struct"`
}

func (s DescribeServiceLinkedRoleStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeServiceLinkedRoleStatusRequest) GetSdkRequest() *DescribeServiceLinkedRoleStatusRequestSdkRequest {
	return s.SdkRequest
}

func (s *DescribeServiceLinkedRoleStatusRequest) SetRegionId(v string) *DescribeServiceLinkedRoleStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusRequest) SetSdkRequest(v *DescribeServiceLinkedRoleStatusRequestSdkRequest) *DescribeServiceLinkedRoleStatusRequest {
	s.SdkRequest = v
	return s
}

func (s *DescribeServiceLinkedRoleStatusRequest) Validate() error {
	if s.SdkRequest != nil {
		if err := s.SdkRequest.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeServiceLinkedRoleStatusRequestSdkRequest struct {
	// The service-linked role. Default value: **AliyunServiceRoleForSas**. Valid values:
	//
	// - **AliyunServiceRoleForSas**: the service-linked role for Security Center (sas). Security Center uses this role to access your resources in other Alibaba Cloud services.
	//
	// - **AliyunServiceRoleForSasCspm**: the service-linked role for Security Center - CSPM (sas-cspm). sas-cspm uses this role to access your resources in other Alibaba Cloud services.
	//
	// example:
	//
	// AliyunServiceRoleForSas
	ServiceLinkedRole *string `json:"ServiceLinkedRole,omitempty" xml:"ServiceLinkedRole,omitempty"`
}

func (s DescribeServiceLinkedRoleStatusRequestSdkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeServiceLinkedRoleStatusRequestSdkRequest) GoString() string {
	return s.String()
}

func (s *DescribeServiceLinkedRoleStatusRequestSdkRequest) GetServiceLinkedRole() *string {
	return s.ServiceLinkedRole
}

func (s *DescribeServiceLinkedRoleStatusRequestSdkRequest) SetServiceLinkedRole(v string) *DescribeServiceLinkedRoleStatusRequestSdkRequest {
	s.ServiceLinkedRole = &v
	return s
}

func (s *DescribeServiceLinkedRoleStatusRequestSdkRequest) Validate() error {
	return dara.Validate(s)
}

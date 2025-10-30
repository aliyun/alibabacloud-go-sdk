// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcEndpointServiceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *GetVpcEndpointServiceAttributeRequest
	GetRegionId() *string
	SetServiceId(v string) *GetVpcEndpointServiceAttributeRequest
	GetServiceId() *string
}

type GetVpcEndpointServiceAttributeRequest struct {
	// The region ID of the endpoint service.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/120468.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The endpoint service ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// epsrv-hp3vpx8yqxblby3i****
	ServiceId *string `json:"ServiceId,omitempty" xml:"ServiceId,omitempty"`
}

func (s GetVpcEndpointServiceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointServiceAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointServiceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcEndpointServiceAttributeRequest) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetVpcEndpointServiceAttributeRequest) SetRegionId(v string) *GetVpcEndpointServiceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeRequest) SetServiceId(v string) *GetVpcEndpointServiceAttributeRequest {
	s.ServiceId = &v
	return s
}

func (s *GetVpcEndpointServiceAttributeRequest) Validate() error {
	return dara.Validate(s)
}

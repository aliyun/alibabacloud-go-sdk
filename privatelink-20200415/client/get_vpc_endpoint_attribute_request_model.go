// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVpcEndpointAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndpointId(v string) *GetVpcEndpointAttributeRequest
	GetEndpointId() *string
	SetRegionId(v string) *GetVpcEndpointAttributeRequest
	GetRegionId() *string
}

type GetVpcEndpointAttributeRequest struct {
	// The ID of the endpoint whose attributes you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// ep-hp33b2e43fays7s8****
	EndpointId *string `json:"EndpointId,omitempty" xml:"EndpointId,omitempty"`
	// The region ID of the endpoint whose attributes you want to query.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/448570.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-huhehaote
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s GetVpcEndpointAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetVpcEndpointAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetVpcEndpointAttributeRequest) GetEndpointId() *string {
	return s.EndpointId
}

func (s *GetVpcEndpointAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetVpcEndpointAttributeRequest) SetEndpointId(v string) *GetVpcEndpointAttributeRequest {
	s.EndpointId = &v
	return s
}

func (s *GetVpcEndpointAttributeRequest) SetRegionId(v string) *GetVpcEndpointAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *GetVpcEndpointAttributeRequest) Validate() error {
	return dara.Validate(s)
}

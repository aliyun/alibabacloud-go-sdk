// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransformEipSegmentToPublicIpAddressPoolRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *TransformEipSegmentToPublicIpAddressPoolRequest
	GetClientToken() *string
	SetDescription(v string) *TransformEipSegmentToPublicIpAddressPoolRequest
	GetDescription() *string
	SetInstanceId(v string) *TransformEipSegmentToPublicIpAddressPoolRequest
	GetInstanceId() *string
	SetName(v string) *TransformEipSegmentToPublicIpAddressPoolRequest
	GetName() *string
	SetRegionId(v string) *TransformEipSegmentToPublicIpAddressPoolRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *TransformEipSegmentToPublicIpAddressPoolRequest
	GetResourceGroupId() *string
}

type TransformEipSegmentToPublicIpAddressPoolRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among all requests. The token can contain only ASCII characters.
	//
	// >  If you do not specify this parameter, the system automatically uses the **request ID*	- as the **client token**. The **request ID*	- is different for each request.
	//
	// example:
	//
	// 02fb3da4-130e-11****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the IP address pool.
	//
	// The description must be 0 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the contiguous EIP group to be migrated.
	//
	// This parameter is required.
	//
	// example:
	//
	// eipsg-2zett8ba055tbsxme****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the IP address pool.
	//
	// The name must be 0 to 128 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// AddressPoolName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the region to which the contiguous EIP group belongs. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the address pool belongs.
	//
	// example:
	//
	// rg-acfmxazb4pcdvf****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s TransformEipSegmentToPublicIpAddressPoolRequest) String() string {
	return dara.Prettify(s)
}

func (s TransformEipSegmentToPublicIpAddressPoolRequest) GoString() string {
	return s.String()
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) GetDescription() *string {
	return s.Description
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) GetName() *string {
	return s.Name
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) SetClientToken(v string) *TransformEipSegmentToPublicIpAddressPoolRequest {
	s.ClientToken = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) SetDescription(v string) *TransformEipSegmentToPublicIpAddressPoolRequest {
	s.Description = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) SetInstanceId(v string) *TransformEipSegmentToPublicIpAddressPoolRequest {
	s.InstanceId = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) SetName(v string) *TransformEipSegmentToPublicIpAddressPoolRequest {
	s.Name = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) SetRegionId(v string) *TransformEipSegmentToPublicIpAddressPoolRequest {
	s.RegionId = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) SetResourceGroupId(v string) *TransformEipSegmentToPublicIpAddressPoolRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *TransformEipSegmentToPublicIpAddressPoolRequest) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProbeAccessPointNetworkQualityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointIds(v []*int32) *ProbeAccessPointNetworkQualityRequest
	GetAccessPointIds() []*int32
	SetOwnerAccount(v string) *ProbeAccessPointNetworkQualityRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ProbeAccessPointNetworkQualityRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ProbeAccessPointNetworkQualityRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ProbeAccessPointNetworkQualityRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ProbeAccessPointNetworkQualityRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *ProbeAccessPointNetworkQualityRequest
	GetSmartAGId() *string
}

type ProbeAccessPointNetworkQualityRequest struct {
	// The IDs of the access point.
	//
	// This parameter is required.
	AccessPointIds []*int32 `json:"AccessPointIds,omitempty" xml:"AccessPointIds,omitempty" type:"Repeated"`
	OwnerAccount   *string  `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region where the SAG instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shanghai
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag-6z21oj0vjjrx6s****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s ProbeAccessPointNetworkQualityRequest) String() string {
	return dara.Prettify(s)
}

func (s ProbeAccessPointNetworkQualityRequest) GoString() string {
	return s.String()
}

func (s *ProbeAccessPointNetworkQualityRequest) GetAccessPointIds() []*int32 {
	return s.AccessPointIds
}

func (s *ProbeAccessPointNetworkQualityRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ProbeAccessPointNetworkQualityRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ProbeAccessPointNetworkQualityRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ProbeAccessPointNetworkQualityRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ProbeAccessPointNetworkQualityRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ProbeAccessPointNetworkQualityRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *ProbeAccessPointNetworkQualityRequest) SetAccessPointIds(v []*int32) *ProbeAccessPointNetworkQualityRequest {
	s.AccessPointIds = v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) SetOwnerAccount(v string) *ProbeAccessPointNetworkQualityRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) SetOwnerId(v int64) *ProbeAccessPointNetworkQualityRequest {
	s.OwnerId = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) SetRegionId(v string) *ProbeAccessPointNetworkQualityRequest {
	s.RegionId = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) SetResourceOwnerAccount(v string) *ProbeAccessPointNetworkQualityRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) SetResourceOwnerId(v int64) *ProbeAccessPointNetworkQualityRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) SetSmartAGId(v string) *ProbeAccessPointNetworkQualityRequest {
	s.SmartAGId = &v
	return s
}

func (s *ProbeAccessPointNetworkQualityRequest) Validate() error {
	return dara.Validate(s)
}

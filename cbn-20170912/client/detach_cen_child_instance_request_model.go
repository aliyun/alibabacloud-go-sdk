// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDetachCenChildInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCenId(v string) *DetachCenChildInstanceRequest
	GetCenId() *string
	SetCenOwnerId(v int64) *DetachCenChildInstanceRequest
	GetCenOwnerId() *int64
	SetChildInstanceId(v string) *DetachCenChildInstanceRequest
	GetChildInstanceId() *string
	SetChildInstanceOwnerId(v int64) *DetachCenChildInstanceRequest
	GetChildInstanceOwnerId() *int64
	SetChildInstanceRegionId(v string) *DetachCenChildInstanceRequest
	GetChildInstanceRegionId() *string
	SetChildInstanceType(v string) *DetachCenChildInstanceRequest
	GetChildInstanceType() *string
	SetOwnerAccount(v string) *DetachCenChildInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DetachCenChildInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DetachCenChildInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DetachCenChildInstanceRequest
	GetResourceOwnerId() *int64
}

type DetachCenChildInstanceRequest struct {
	// The ID of the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// cen-7qthudw0ll6jmx****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The ID of the Alibaba Cloud account to which the CEN instance belongs.
	//
	// example:
	//
	// 1688000000000000
	CenOwnerId *int64 `json:"CenOwnerId,omitempty" xml:"CenOwnerId,omitempty"`
	// The ID of the network instance that you want to detach from the CEN instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-bp18sth14qii3pnvx****
	ChildInstanceId *string `json:"ChildInstanceId,omitempty" xml:"ChildInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the network instance belongs.
	//
	// example:
	//
	// 1699000000000000
	ChildInstanceOwnerId *int64 `json:"ChildInstanceOwnerId,omitempty" xml:"ChildInstanceOwnerId,omitempty"`
	// The ID of the region where the network instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/36063.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	ChildInstanceRegionId *string `json:"ChildInstanceRegionId,omitempty" xml:"ChildInstanceRegionId,omitempty"`
	// The type of the network instance. Valid values:
	//
	// 	- **VPC**: virtual private cloud (VPC)
	//
	// 	- **VBR**: virtual border router (VBR)
	//
	// 	- **CCN**: Cloud Connect Network (CCN) instance
	//
	// This parameter is required.
	//
	// example:
	//
	// VPC
	ChildInstanceType    *string `json:"ChildInstanceType,omitempty" xml:"ChildInstanceType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DetachCenChildInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DetachCenChildInstanceRequest) GoString() string {
	return s.String()
}

func (s *DetachCenChildInstanceRequest) GetCenId() *string {
	return s.CenId
}

func (s *DetachCenChildInstanceRequest) GetCenOwnerId() *int64 {
	return s.CenOwnerId
}

func (s *DetachCenChildInstanceRequest) GetChildInstanceId() *string {
	return s.ChildInstanceId
}

func (s *DetachCenChildInstanceRequest) GetChildInstanceOwnerId() *int64 {
	return s.ChildInstanceOwnerId
}

func (s *DetachCenChildInstanceRequest) GetChildInstanceRegionId() *string {
	return s.ChildInstanceRegionId
}

func (s *DetachCenChildInstanceRequest) GetChildInstanceType() *string {
	return s.ChildInstanceType
}

func (s *DetachCenChildInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DetachCenChildInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DetachCenChildInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DetachCenChildInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DetachCenChildInstanceRequest) SetCenId(v string) *DetachCenChildInstanceRequest {
	s.CenId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetCenOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.CenOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceId(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.ChildInstanceOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceRegionId(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceRegionId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetChildInstanceType(v string) *DetachCenChildInstanceRequest {
	s.ChildInstanceType = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetOwnerAccount(v string) *DetachCenChildInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetResourceOwnerAccount(v string) *DetachCenChildInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DetachCenChildInstanceRequest) SetResourceOwnerId(v int64) *DetachCenChildInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DetachCenChildInstanceRequest) Validate() error {
	return dara.Validate(s)
}

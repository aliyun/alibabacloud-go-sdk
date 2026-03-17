// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSAGDeviceInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSAGDeviceInfoRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSAGDeviceInfoRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSAGDeviceInfoRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSAGDeviceInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSAGDeviceInfoRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSAGDeviceInfoRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSAGDeviceInfoRequest
	GetSmartAGSn() *string
}

type DescribeSAGDeviceInfoRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// sag-7f3d9b6jwnuqn6****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// The serial number of the SAG device associated with the SAG instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// sag32a30****
	SmartAGSn *string `json:"SmartAGSn,omitempty" xml:"SmartAGSn,omitempty"`
}

func (s DescribeSAGDeviceInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSAGDeviceInfoRequest) GoString() string {
	return s.String()
}

func (s *DescribeSAGDeviceInfoRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSAGDeviceInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSAGDeviceInfoRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSAGDeviceInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSAGDeviceInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSAGDeviceInfoRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSAGDeviceInfoRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSAGDeviceInfoRequest) SetOwnerAccount(v string) *DescribeSAGDeviceInfoRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) SetOwnerId(v int64) *DescribeSAGDeviceInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) SetRegionId(v string) *DescribeSAGDeviceInfoRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) SetResourceOwnerAccount(v string) *DescribeSAGDeviceInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) SetResourceOwnerId(v int64) *DescribeSAGDeviceInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) SetSmartAGId(v string) *DescribeSAGDeviceInfoRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) SetSmartAGSn(v string) *DescribeSAGDeviceInfoRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSAGDeviceInfoRequest) Validate() error {
	return dara.Validate(s)
}

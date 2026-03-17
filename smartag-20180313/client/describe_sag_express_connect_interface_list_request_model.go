// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagExpressConnectInterfaceListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSagExpressConnectInterfaceListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSagExpressConnectInterfaceListRequest
	GetOwnerId() *int64
	SetPortName(v string) *DescribeSagExpressConnectInterfaceListRequest
	GetPortName() *string
	SetRegionId(v string) *DescribeSagExpressConnectInterfaceListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSagExpressConnectInterfaceListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSagExpressConnectInterfaceListRequest
	GetResourceOwnerId() *int64
	SetSmartAGId(v string) *DescribeSagExpressConnectInterfaceListRequest
	GetSmartAGId() *string
	SetSmartAGSn(v string) *DescribeSagExpressConnectInterfaceListRequest
	GetSmartAGSn() *string
}

type DescribeSagExpressConnectInterfaceListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the Express Connect circuit port.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	PortName *string `json:"PortName,omitempty" xml:"PortName,omitempty"`
	// The ID of the region where the Smart Access Gateway (SAG) instance is deployed.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/69813.html) operation to query the most recent region list.
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
	// sag-whfn****
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

func (s DescribeSagExpressConnectInterfaceListRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagExpressConnectInterfaceListRequest) GoString() string {
	return s.String()
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetPortName() *string {
	return s.PortName
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagExpressConnectInterfaceListRequest) GetSmartAGSn() *string {
	return s.SmartAGSn
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetOwnerAccount(v string) *DescribeSagExpressConnectInterfaceListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetOwnerId(v int64) *DescribeSagExpressConnectInterfaceListRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetPortName(v string) *DescribeSagExpressConnectInterfaceListRequest {
	s.PortName = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetRegionId(v string) *DescribeSagExpressConnectInterfaceListRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetResourceOwnerAccount(v string) *DescribeSagExpressConnectInterfaceListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetResourceOwnerId(v int64) *DescribeSagExpressConnectInterfaceListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetSmartAGId(v string) *DescribeSagExpressConnectInterfaceListRequest {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) SetSmartAGSn(v string) *DescribeSagExpressConnectInterfaceListRequest {
	s.SmartAGSn = &v
	return s
}

func (s *DescribeSagExpressConnectInterfaceListRequest) Validate() error {
	return dara.Validate(s)
}

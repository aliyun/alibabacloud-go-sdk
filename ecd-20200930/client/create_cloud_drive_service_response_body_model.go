// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCdsId(v string) *CreateCloudDriveServiceResponseBody
	GetCdsId() *string
	SetCdsName(v string) *CreateCloudDriveServiceResponseBody
	GetCdsName() *string
	SetCenId(v string) *CreateCloudDriveServiceResponseBody
	GetCenId() *string
	SetConflictCdsAndOrder(v *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) *CreateCloudDriveServiceResponseBody
	GetConflictCdsAndOrder() *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder
	SetDomainName(v string) *CreateCloudDriveServiceResponseBody
	GetDomainName() *string
	SetErrorCode(v string) *CreateCloudDriveServiceResponseBody
	GetErrorCode() *string
	SetMaxSize(v string) *CreateCloudDriveServiceResponseBody
	GetMaxSize() *string
	SetOfficeSiteType(v string) *CreateCloudDriveServiceResponseBody
	GetOfficeSiteType() *string
	SetOrderId(v string) *CreateCloudDriveServiceResponseBody
	GetOrderId() *string
	SetRequestId(v string) *CreateCloudDriveServiceResponseBody
	GetRequestId() *string
}

type CreateCloudDriveServiceResponseBody struct {
	// The ID of the enterprise network disk.
	//
	// example:
	//
	// cn-hangzhou+cds-778205****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The name of the cloud storage service.
	//
	// example:
	//
	// CDS_Windows_1126
	CdsName *string `json:"CdsName,omitempty" xml:"CdsName,omitempty"`
	// The ID of the Cloud Enterprise Network (CEN) instance.
	//
	// > If you want to connect to a cloud desktop through a VPC, you can add the office network to a CEN instance. This CEN instance connects to your on-premises network through a VPN or a dedicated line.
	//
	// example:
	//
	// cen-638u3wxds9snyc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The existing enterprise network disks or orders for enterprise network disks that conflict with the enterprise network disk that you want to create.
	ConflictCdsAndOrder *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder `json:"ConflictCdsAndOrder,omitempty" xml:"ConflictCdsAndOrder,omitempty" type:"Struct"`
	// The name of the Active Directory (AD) domain corresponding to the AD office network.
	//
	// example:
	//
	// test1.local
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// The error code.
	//
	// example:
	//
	// ExistConflictCds
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The maximum storage usage of the enterprise network disk. Unit: bytes.
	//
	// example:
	//
	// 5368709120
	MaxSize *string `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The office network type.
	//
	// example:
	//
	// AD_CONNECTOR
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The order ID. You can obtain the order ID on the Order Management page in the Alibaba Cloud User Center.
	//
	// example:
	//
	// 214552063030752
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 062B1439-709A-580E-85DF-CE97A156****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudDriveServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveServiceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveServiceResponseBody) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCloudDriveServiceResponseBody) GetCdsName() *string {
	return s.CdsName
}

func (s *CreateCloudDriveServiceResponseBody) GetCenId() *string {
	return s.CenId
}

func (s *CreateCloudDriveServiceResponseBody) GetConflictCdsAndOrder() *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder {
	return s.ConflictCdsAndOrder
}

func (s *CreateCloudDriveServiceResponseBody) GetDomainName() *string {
	return s.DomainName
}

func (s *CreateCloudDriveServiceResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateCloudDriveServiceResponseBody) GetMaxSize() *string {
	return s.MaxSize
}

func (s *CreateCloudDriveServiceResponseBody) GetOfficeSiteType() *string {
	return s.OfficeSiteType
}

func (s *CreateCloudDriveServiceResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateCloudDriveServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudDriveServiceResponseBody) SetCdsId(v string) *CreateCloudDriveServiceResponseBody {
	s.CdsId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetCdsName(v string) *CreateCloudDriveServiceResponseBody {
	s.CdsName = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetCenId(v string) *CreateCloudDriveServiceResponseBody {
	s.CenId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetConflictCdsAndOrder(v *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) *CreateCloudDriveServiceResponseBody {
	s.ConflictCdsAndOrder = v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetDomainName(v string) *CreateCloudDriveServiceResponseBody {
	s.DomainName = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetErrorCode(v string) *CreateCloudDriveServiceResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetMaxSize(v string) *CreateCloudDriveServiceResponseBody {
	s.MaxSize = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetOfficeSiteType(v string) *CreateCloudDriveServiceResponseBody {
	s.OfficeSiteType = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetOrderId(v string) *CreateCloudDriveServiceResponseBody {
	s.OrderId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) SetRequestId(v string) *CreateCloudDriveServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBody) Validate() error {
	if s.ConflictCdsAndOrder != nil {
		if err := s.ConflictCdsAndOrder.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCloudDriveServiceResponseBodyConflictCdsAndOrder struct {
	// The information of conflicting enterprise network disks.
	ConflictCds []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds `json:"ConflictCds,omitempty" xml:"ConflictCds,omitempty" type:"Repeated"`
	// The information of conflicting unpaid orders for enterprise network disks (orders for subscription enterprise network disks).
	ConflictOrder []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder `json:"ConflictOrder,omitempty" xml:"ConflictOrder,omitempty" type:"Repeated"`
}

func (s CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) GetConflictCds() []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds {
	return s.ConflictCds
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) GetConflictOrder() []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder {
	return s.ConflictOrder
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) SetConflictCds(v []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder {
	s.ConflictCds = v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) SetConflictOrder(v []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder {
	s.ConflictOrder = v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder) Validate() error {
	if s.ConflictCds != nil {
		for _, item := range s.ConflictCds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ConflictOrder != nil {
		for _, item := range s.ConflictOrder {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds struct {
	// The ID of the enterprise network disk.
	//
	// example:
	//
	// cn-hangzhou+cds-778205****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the regions that are supported by Elastic Desktop Service.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) SetCdsId(v string) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds {
	s.CdsId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) SetRegionId(v string) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds {
	s.RegionId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds) Validate() error {
	return dara.Validate(s)
}

type CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder struct {
	// The ID of the enterprise network disk (The order is not paid, and the enterprise network disk cannot be used).
	//
	// example:
	//
	// cn-hangzhou+cds-778205****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The order ID. You can obtain the order ID on the **Order Management*	- page in the Alibaba Cloud User Center.
	//
	// example:
	//
	// 22442411898****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) GetCdsId() *string {
	return s.CdsId
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) GetOrderId() *string {
	return s.OrderId
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) SetCdsId(v string) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder {
	s.CdsId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) SetOrderId(v string) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder {
	s.OrderId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) SetRegionId(v string) *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder {
	s.RegionId = &v
	return s
}

func (s *CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictOrder) Validate() error {
	return dara.Validate(s)
}

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
	// The ID of the enterprise drive.
	//
	// example:
	//
	// cn-hangzhou+cds-7782057786
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The name of the cloud disk that is created in Cloud Drive Service.
	//
	// example:
	//
	// CDS_Windows_1126
	CdsName *string `json:"CdsName,omitempty" xml:"CdsName,omitempty"`
	// The ID of the CEN instance.
	//
	// >  To allow end users to connect to cloud computers via virtual private clouds (VPCs), attach your office network to a CEN instance. The CEN instance connects to your on-premises network through VPN Gateway or Express Connect.
	//
	// example:
	//
	// cen-638u3wxds9snyc****
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The existing enterprise drive or its order that conflicts with the enterprise drive being created.
	ConflictCdsAndOrder *CreateCloudDriveServiceResponseBodyConflictCdsAndOrder `json:"ConflictCdsAndOrder,omitempty" xml:"ConflictCdsAndOrder,omitempty" type:"Struct"`
	// The domain name of the enterprise AD office network.
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
	// The maximum storage capacity of the enterprise drive. Unit: bytes.
	//
	// example:
	//
	// 213674622976
	MaxSize *string `json:"MaxSize,omitempty" xml:"MaxSize,omitempty"`
	// The type of the office network.
	//
	// Valid values:
	//
	// 	- SIMPLE: convenience office network.
	//
	// 	- AD_CONNECTOR: enterprise AD office network.
	//
	// example:
	//
	// AD_CONNECTOR
	OfficeSiteType *string `json:"OfficeSiteType,omitempty" xml:"OfficeSiteType,omitempty"`
	// The ID of the order. You can obtain an order ID on the Orders page in the Expenses and Costs console.
	//
	// example:
	//
	// 214552063030752
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 062B1439-709A-580E-85DF-CE97A1560565
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
	return dara.Validate(s)
}

type CreateCloudDriveServiceResponseBodyConflictCdsAndOrder struct {
	// The conflicting enterprise drive.
	ConflictCds []*CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds `json:"ConflictCds,omitempty" xml:"ConflictCds,omitempty" type:"Repeated"`
	// The subscription orders of the conflicting enterprise drives that are unpaid.
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
	return dara.Validate(s)
}

type CreateCloudDriveServiceResponseBodyConflictCdsAndOrderConflictCds struct {
	// The ID of the enterprise drive.
	//
	// example:
	//
	// cn-hangzhou+cds-778205****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the list of regions where Elastic Desktop Service (EDS) Enterprise is available.
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
	// The ID of the enterprise drive. The enterprise drive cannot be used if the order is unpaid.
	//
	// example:
	//
	// cn-hangzhou+cds-778205****
	CdsId *string `json:"CdsId,omitempty" xml:"CdsId,omitempty"`
	// The ID of the order. You can obtain an order ID on the **Orders*	- page in the Expenses and Costs console.
	//
	// example:
	//
	// 22442411898****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the region.
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

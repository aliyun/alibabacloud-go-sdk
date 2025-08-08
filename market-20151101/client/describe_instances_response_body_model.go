// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceItems(v *DescribeInstancesResponseBodyInstanceItems) *DescribeInstancesResponseBody
	GetInstanceItems() *DescribeInstancesResponseBodyInstanceItems
	SetPageNumber(v int32) *DescribeInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeInstancesResponseBody struct {
	InstanceItems *DescribeInstancesResponseBodyInstanceItems `json:"InstanceItems,omitempty" xml:"InstanceItems,omitempty" type:"Struct"`
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 54C22FB9-0CB1-5629-97A8-653FC7990F00
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 55
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBody) GetInstanceItems() *DescribeInstancesResponseBodyInstanceItems {
	return s.InstanceItems
}

func (s *DescribeInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstancesResponseBody) SetInstanceItems(v *DescribeInstancesResponseBodyInstanceItems) *DescribeInstancesResponseBody {
	s.InstanceItems = v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageNumber(v int32) *DescribeInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetPageSize(v int32) *DescribeInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetRequestId(v string) *DescribeInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstancesResponseBody) SetTotalCount(v int32) *DescribeInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstanceItems struct {
	InstanceItem []*DescribeInstancesResponseBodyInstanceItemsInstanceItem `json:"InstanceItem,omitempty" xml:"InstanceItem,omitempty" type:"Repeated"`
}

func (s DescribeInstancesResponseBodyInstanceItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstanceItems) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstanceItems) GetInstanceItem() []*DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	return s.InstanceItem
}

func (s *DescribeInstancesResponseBodyInstanceItems) SetInstanceItem(v []*DescribeInstancesResponseBodyInstanceItemsInstanceItem) *DescribeInstancesResponseBodyInstanceItems {
	s.InstanceItem = v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItems) Validate() error {
	return dara.Validate(s)
}

type DescribeInstancesResponseBodyInstanceItemsInstanceItem struct {
	// example:
	//
	// {}
	ApiJson *string `json:"ApiJson,omitempty" xml:"ApiJson,omitempty"`
	// example:
	//
	// {"frontEndUrl":"https://***.aliyun.com","password":"Sjtv***","adminUrl":"https://***.aiiyun.com","username":"aliyun***"}
	AppJson *string `json:"AppJson,omitempty" xml:"AppJson,omitempty"`
	// example:
	//
	// 1570634021000
	BeganOn *int64 `json:"BeganOn,omitempty" xml:"BeganOn,omitempty"`
	// example:
	//
	// 1570634021000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// example:
	//
	// 1570644021000
	EndOn      *int64  `json:"EndOn,omitempty" xml:"EndOn,omitempty"`
	ExtendJson *string `json:"ExtendJson,omitempty" xml:"ExtendJson,omitempty"`
	// example:
	//
	// {"password":"***","ip":"118.31.***.41","innerIp":"118.31.***.41","region":"","username":"***","beianInfo":""}
	HostJson *string `json:"HostJson,omitempty" xml:"HostJson,omitempty"`
	// example:
	//
	// {}
	IdaasJson *string `json:"IdaasJson,omitempty" xml:"IdaasJson,omitempty"`
	// example:
	//
	// {}
	ImageJson *string `json:"ImageJson,omitempty" xml:"ImageJson,omitempty"`
	// example:
	//
	// 1551111111
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 204211111111111
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// example:
	//
	// cmgj00**11
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	// example:
	//
	// cmgj00**11-prepay
	ProductSkuCode *string `json:"ProductSkuCode,omitempty" xml:"ProductSkuCode,omitempty"`
	// example:
	//
	// APP
	ProductType *string `json:"ProductType,omitempty" xml:"ProductType,omitempty"`
	// example:
	//
	// OPENED
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupplierName *string `json:"SupplierName,omitempty" xml:"SupplierName,omitempty"`
}

func (s DescribeInstancesResponseBodyInstanceItemsInstanceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstancesResponseBodyInstanceItemsInstanceItem) GoString() string {
	return s.String()
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetApiJson() *string {
	return s.ApiJson
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetAppJson() *string {
	return s.AppJson
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetBeganOn() *int64 {
	return s.BeganOn
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetEndOn() *int64 {
	return s.EndOn
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetExtendJson() *string {
	return s.ExtendJson
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetHostJson() *string {
	return s.HostJson
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetIdaasJson() *string {
	return s.IdaasJson
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetImageJson() *string {
	return s.ImageJson
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetOrderId() *int64 {
	return s.OrderId
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetProductSkuCode() *string {
	return s.ProductSkuCode
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetProductType() *string {
	return s.ProductType
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetStatus() *string {
	return s.Status
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) GetSupplierName() *string {
	return s.SupplierName
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetApiJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ApiJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetAppJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.AppJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetBeganOn(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.BeganOn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetCreatedOn(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.CreatedOn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetEndOn(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.EndOn = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetExtendJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ExtendJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetHostJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.HostJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetIdaasJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.IdaasJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetImageJson(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ImageJson = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetInstanceId(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetOrderId(v int64) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.OrderId = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductCode(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductName(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductSkuCode(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductSkuCode = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetProductType(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.ProductType = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetStatus(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.Status = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) SetSupplierName(v string) *DescribeInstancesResponseBodyInstanceItemsInstanceItem {
	s.SupplierName = &v
	return s
}

func (s *DescribeInstancesResponseBodyInstanceItemsInstanceItem) Validate() error {
	return dara.Validate(s)
}

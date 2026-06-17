// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseOrdersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeLicenseOrdersResponseBodyItems) *DescribeLicenseOrdersResponseBody
	GetItems() []*DescribeLicenseOrdersResponseBodyItems
	SetPageNumber(v int32) *DescribeLicenseOrdersResponseBody
	GetPageNumber() *int32
	SetPageRecordCount(v int32) *DescribeLicenseOrdersResponseBody
	GetPageRecordCount() *int32
	SetRequestId(v string) *DescribeLicenseOrdersResponseBody
	GetRequestId() *string
	SetTotalRecordCount(v int32) *DescribeLicenseOrdersResponseBody
	GetTotalRecordCount() *int32
}

type DescribeLicenseOrdersResponseBody struct {
	// The list of orders.
	Items []*DescribeLicenseOrdersResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of records on the current page.
	//
	// example:
	//
	// 12
	PageRecordCount *int32 `json:"PageRecordCount,omitempty" xml:"PageRecordCount,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34458CD3-33E0-4624-BFEF-840C15******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of records.
	//
	// example:
	//
	// 50
	TotalRecordCount *int32 `json:"TotalRecordCount,omitempty" xml:"TotalRecordCount,omitempty"`
}

func (s DescribeLicenseOrdersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrdersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrdersResponseBody) GetItems() []*DescribeLicenseOrdersResponseBodyItems {
	return s.Items
}

func (s *DescribeLicenseOrdersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLicenseOrdersResponseBody) GetPageRecordCount() *int32 {
	return s.PageRecordCount
}

func (s *DescribeLicenseOrdersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLicenseOrdersResponseBody) GetTotalRecordCount() *int32 {
	return s.TotalRecordCount
}

func (s *DescribeLicenseOrdersResponseBody) SetItems(v []*DescribeLicenseOrdersResponseBodyItems) *DescribeLicenseOrdersResponseBody {
	s.Items = v
	return s
}

func (s *DescribeLicenseOrdersResponseBody) SetPageNumber(v int32) *DescribeLicenseOrdersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBody) SetPageRecordCount(v int32) *DescribeLicenseOrdersResponseBody {
	s.PageRecordCount = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBody) SetRequestId(v string) *DescribeLicenseOrdersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBody) SetTotalRecordCount(v int32) *DescribeLicenseOrdersResponseBody {
	s.TotalRecordCount = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLicenseOrdersResponseBodyItems struct {
	// The number of activation codes that have been generated.
	//
	// example:
	//
	// 10
	ActivatedCodeCount *int32 `json:"ActivatedCodeCount,omitempty" xml:"ActivatedCodeCount,omitempty"`
	// The quota for requesting activation codes.
	//
	// example:
	//
	// 10
	ActivationCodeQuota *int32 `json:"ActivationCodeQuota,omitempty" xml:"ActivationCodeQuota,omitempty"`
	// The Alibaba Cloud order ID or virtual order ID.
	//
	// example:
	//
	// 227638319690519
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// Specifies whether the System Identifier can be left empty when an activation code is generated.
	//
	// example:
	//
	// false
	AllowEmptySystemIdentifier *bool `json:"AllowEmptySystemIdentifier,omitempty" xml:"AllowEmptySystemIdentifier,omitempty"`
	// The database type, such as PG, Oracle, or MySQL.
	//
	// example:
	//
	// PG
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-02-11 03:14:15
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-02-11 03:14:15
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the order is a virtual order. Virtual orders allow for pre-generating activation codes.
	//
	// example:
	//
	// false
	IsVirtualOrder *bool `json:"IsVirtualOrder,omitempty" xml:"IsVirtualOrder,omitempty"`
	// Indicates whether the virtual order is frozen. No more activation codes can be generated from a frozen order.
	//
	// example:
	//
	// false
	IsVirtualOrderFrozen *bool `json:"IsVirtualOrderFrozen,omitempty" xml:"IsVirtualOrderFrozen,omitempty"`
	// The package type. Valid values:
	//
	// - single_node_subscribe: single node (subscription)
	//
	// - single_node_long_term: single node (long-term)
	//
	// - primary_backup_subscribe: primary/standby (subscription)
	//
	// - primary_backup_long_term: primary/standby (long-term)
	//
	// - pre_generation_long_term: pre-generated (long-term)
	//
	// example:
	//
	// single_node_subscribe
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The validity period of the package. Common options are one year or long-term (30 years).
	//
	// example:
	//
	// 1 year
	PackageValidity *string `json:"PackageValidity,omitempty" xml:"PackageValidity,omitempty"`
	// The purchase channel. Valid values: \\`aliyun_market\\` (Alibaba Cloud Marketplace) and \\`aliyun_public\\` (standard purchase page).
	//
	// example:
	//
	// aliyun_public
	PurchaseChannel *string `json:"PurchaseChannel,omitempty" xml:"PurchaseChannel,omitempty"`
	// The virtual order ID.
	//
	// example:
	//
	// 227638319690519
	VirtualAliyunOrderId *string `json:"VirtualAliyunOrderId,omitempty" xml:"VirtualAliyunOrderId,omitempty"`
}

func (s DescribeLicenseOrdersResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrdersResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetActivatedCodeCount() *int32 {
	return s.ActivatedCodeCount
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetActivationCodeQuota() *int32 {
	return s.ActivationCodeQuota
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetAllowEmptySystemIdentifier() *bool {
	return s.AllowEmptySystemIdentifier
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetEngine() *string {
	return s.Engine
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetIsVirtualOrder() *bool {
	return s.IsVirtualOrder
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetIsVirtualOrderFrozen() *bool {
	return s.IsVirtualOrderFrozen
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetPackageValidity() *string {
	return s.PackageValidity
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetPurchaseChannel() *string {
	return s.PurchaseChannel
}

func (s *DescribeLicenseOrdersResponseBodyItems) GetVirtualAliyunOrderId() *string {
	return s.VirtualAliyunOrderId
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetActivatedCodeCount(v int32) *DescribeLicenseOrdersResponseBodyItems {
	s.ActivatedCodeCount = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetActivationCodeQuota(v int32) *DescribeLicenseOrdersResponseBodyItems {
	s.ActivationCodeQuota = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetAliyunOrderId(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.AliyunOrderId = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetAllowEmptySystemIdentifier(v bool) *DescribeLicenseOrdersResponseBodyItems {
	s.AllowEmptySystemIdentifier = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetEngine(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.Engine = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetGmtCreated(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.GmtCreated = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetGmtModified(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.GmtModified = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetIsVirtualOrder(v bool) *DescribeLicenseOrdersResponseBodyItems {
	s.IsVirtualOrder = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetIsVirtualOrderFrozen(v bool) *DescribeLicenseOrdersResponseBodyItems {
	s.IsVirtualOrderFrozen = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetPackageType(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.PackageType = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetPackageValidity(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.PackageValidity = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetPurchaseChannel(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.PurchaseChannel = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) SetVirtualAliyunOrderId(v string) *DescribeLicenseOrdersResponseBodyItems {
	s.VirtualAliyunOrderId = &v
	return s
}

func (s *DescribeLicenseOrdersResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

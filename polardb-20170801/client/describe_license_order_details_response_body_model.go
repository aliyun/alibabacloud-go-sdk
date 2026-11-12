// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLicenseOrderDetailsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivatedCodeCount(v int32) *DescribeLicenseOrderDetailsResponseBody
	GetActivatedCodeCount() *int32
	SetActivationCodeQuota(v int32) *DescribeLicenseOrderDetailsResponseBody
	GetActivationCodeQuota() *int32
	SetAliyunOrderId(v string) *DescribeLicenseOrderDetailsResponseBody
	GetAliyunOrderId() *string
	SetAllowEmptySystemIdentifier(v bool) *DescribeLicenseOrderDetailsResponseBody
	GetAllowEmptySystemIdentifier() *bool
	SetEngine(v string) *DescribeLicenseOrderDetailsResponseBody
	GetEngine() *string
	SetGmtCreated(v string) *DescribeLicenseOrderDetailsResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *DescribeLicenseOrderDetailsResponseBody
	GetGmtModified() *string
	SetIsVirtualOrder(v bool) *DescribeLicenseOrderDetailsResponseBody
	GetIsVirtualOrder() *bool
	SetIsVirtualOrderFrozen(v bool) *DescribeLicenseOrderDetailsResponseBody
	GetIsVirtualOrderFrozen() *bool
	SetPackageType(v string) *DescribeLicenseOrderDetailsResponseBody
	GetPackageType() *string
	SetPackageValidity(v string) *DescribeLicenseOrderDetailsResponseBody
	GetPackageValidity() *string
	SetPurchaseChannel(v string) *DescribeLicenseOrderDetailsResponseBody
	GetPurchaseChannel() *string
	SetRequestId(v string) *DescribeLicenseOrderDetailsResponseBody
	GetRequestId() *string
	SetVirtualOrderId(v string) *DescribeLicenseOrderDetailsResponseBody
	GetVirtualOrderId() *string
}

type DescribeLicenseOrderDetailsResponseBody struct {
	// The number of generated activation codes.
	//
	// example:
	//
	// 2
	ActivatedCodeCount *int32 `json:"ActivatedCodeCount,omitempty" xml:"ActivatedCodeCount,omitempty"`
	// The quota of activation codes that can be requested.
	//
	// example:
	//
	// 8
	ActivationCodeQuota *int32 `json:"ActivationCodeQuota,omitempty" xml:"ActivationCodeQuota,omitempty"`
	// The Alibaba Cloud order ID, including the virtual order ID.
	//
	// example:
	//
	// 239618016570503
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// Specifies whether the System Identifier can be left empty when generating an activation code.
	//
	// example:
	//
	// false
	AllowEmptySystemIdentifier *bool `json:"AllowEmptySystemIdentifier,omitempty" xml:"AllowEmptySystemIdentifier,omitempty"`
	// The database type, such as PostgreSQL, Oracle, or MySQL.
	//
	// example:
	//
	// PG
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2021-10-19 01:13:45
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Specifies whether the order is a virtual order. Virtual orders allow pre-generation of activation codes.
	//
	// example:
	//
	// false
	IsVirtualOrder *bool `json:"IsVirtualOrder,omitempty" xml:"IsVirtualOrder,omitempty"`
	// Specifies whether the virtual order is frozen. A frozen virtual order does not allow further generation of activation codes.
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
	// - primary_backup_subscribe: primary/secondary edition (subscription)
	//
	// - primary_backup_long_term: primary/secondary edition (long-term)
	//
	// - pre_generation_long_term: pre-generation (long-term)
	//
	// example:
	//
	// pre_generation_long_term
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The validity period of the package, which is typically one year or a long-term period of 30 years.
	//
	// example:
	//
	// 1 year
	PackageValidity *string `json:"PackageValidity,omitempty" xml:"PackageValidity,omitempty"`
	// The purchase channel. Valid values: aliyun_market (Alibaba Cloud Marketplace) and aliyun_public (regular purchase page).
	//
	// example:
	//
	// aliyun_market
	PurchaseChannel *string `json:"PurchaseChannel,omitempty" xml:"PurchaseChannel,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 22C0ACF0-DD29-4B67-9190-B7A48C******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The virtual order ID.
	//
	// example:
	//
	// 239618016570503
	VirtualOrderId *string `json:"VirtualOrderId,omitempty" xml:"VirtualOrderId,omitempty"`
}

func (s DescribeLicenseOrderDetailsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLicenseOrderDetailsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetActivatedCodeCount() *int32 {
	return s.ActivatedCodeCount
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetActivationCodeQuota() *int32 {
	return s.ActivationCodeQuota
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetAllowEmptySystemIdentifier() *bool {
	return s.AllowEmptySystemIdentifier
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetIsVirtualOrder() *bool {
	return s.IsVirtualOrder
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetIsVirtualOrderFrozen() *bool {
	return s.IsVirtualOrderFrozen
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetPackageType() *string {
	return s.PackageType
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetPackageValidity() *string {
	return s.PackageValidity
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetPurchaseChannel() *string {
	return s.PurchaseChannel
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLicenseOrderDetailsResponseBody) GetVirtualOrderId() *string {
	return s.VirtualOrderId
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetActivatedCodeCount(v int32) *DescribeLicenseOrderDetailsResponseBody {
	s.ActivatedCodeCount = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetActivationCodeQuota(v int32) *DescribeLicenseOrderDetailsResponseBody {
	s.ActivationCodeQuota = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetAliyunOrderId(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.AliyunOrderId = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetAllowEmptySystemIdentifier(v bool) *DescribeLicenseOrderDetailsResponseBody {
	s.AllowEmptySystemIdentifier = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetEngine(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetGmtCreated(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetGmtModified(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.GmtModified = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetIsVirtualOrder(v bool) *DescribeLicenseOrderDetailsResponseBody {
	s.IsVirtualOrder = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetIsVirtualOrderFrozen(v bool) *DescribeLicenseOrderDetailsResponseBody {
	s.IsVirtualOrderFrozen = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetPackageType(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.PackageType = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetPackageValidity(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.PackageValidity = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetPurchaseChannel(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.PurchaseChannel = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetRequestId(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) SetVirtualOrderId(v string) *DescribeLicenseOrderDetailsResponseBody {
	s.VirtualOrderId = &v
	return s
}

func (s *DescribeLicenseOrderDetailsResponseBody) Validate() error {
	return dara.Validate(s)
}

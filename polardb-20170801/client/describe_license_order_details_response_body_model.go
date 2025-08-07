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
	// The maximum number of activation codes that you can apply for.
	//
	// example:
	//
	// 8
	ActivationCodeQuota *int32 `json:"ActivationCodeQuota,omitempty" xml:"ActivationCodeQuota,omitempty"`
	// The Alibaba Cloud order ID (including the virtual order ID).
	//
	// example:
	//
	// 239618016570503
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// Indicates whether activation codes can be generated without the system identifier.
	//
	// example:
	//
	// false
	AllowEmptySystemIdentifier *bool `json:"AllowEmptySystemIdentifier,omitempty" xml:"AllowEmptySystemIdentifier,omitempty"`
	// The type of the engine. Valid values: PG, Oracle, and MySQL.
	//
	// example:
	//
	// PG
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The time when the order was created.
	//
	// example:
	//
	// 2021-10-19 01:13:45
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the order was last updated.
	//
	// example:
	//
	// 2024-10-16 16:46:20
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// Indicates whether the order is a virtual order (virtual orders allow pre-generation of activation codes).
	//
	// example:
	//
	// false
	IsVirtualOrder *bool `json:"IsVirtualOrder,omitempty" xml:"IsVirtualOrder,omitempty"`
	// Indicates whether the virtual order is frozen (activation codes cannot be generated for a frozen virtual order).
	//
	// example:
	//
	// false
	IsVirtualOrderFrozen *bool `json:"IsVirtualOrderFrozen,omitempty" xml:"IsVirtualOrderFrozen,omitempty"`
	// The plan type. Valid values:
	//
	// 	- single_node_subscribe
	//
	// 	- single_node_long_term
	//
	// 	- primary_backup_subscribe
	//
	// 	- primary_backup_long_term
	//
	// 	- pre_generation_long_term
	//
	// example:
	//
	// pre_generation_long_term
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The validity period of the plan, which is one year (common) or thirty years (long-term).
	//
	// example:
	//
	// 1 year
	PackageValidity *string `json:"PackageValidity,omitempty" xml:"PackageValidity,omitempty"`
	// The plan validity period, one year (common) or thirty years (long-term).
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

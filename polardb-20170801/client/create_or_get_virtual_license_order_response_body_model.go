// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrGetVirtualLicenseOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivatedCodeCount(v int32) *CreateOrGetVirtualLicenseOrderResponseBody
	GetActivatedCodeCount() *int32
	SetActivationCodeQuota(v int32) *CreateOrGetVirtualLicenseOrderResponseBody
	GetActivationCodeQuota() *int32
	SetAliyunOrderId(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetAliyunOrderId() *string
	SetAllowEmptySystemIdentifier(v bool) *CreateOrGetVirtualLicenseOrderResponseBody
	GetAllowEmptySystemIdentifier() *bool
	SetGmtCreated(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetGmtModified() *string
	SetIsVirtualOrder(v bool) *CreateOrGetVirtualLicenseOrderResponseBody
	GetIsVirtualOrder() *bool
	SetIsVirtualOrderFrozen(v bool) *CreateOrGetVirtualLicenseOrderResponseBody
	GetIsVirtualOrderFrozen() *bool
	SetPackageType(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetPackageType() *string
	SetPackageValidity(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetPackageValidity() *string
	SetPurchaseChannel(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetPurchaseChannel() *string
	SetRequestId(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetRequestId() *string
	SetVirtualOrderId(v string) *CreateOrGetVirtualLicenseOrderResponseBody
	GetVirtualOrderId() *string
}

type CreateOrGetVirtualLicenseOrderResponseBody struct {
	// The number of generated activation codes.
	//
	// example:
	//
	// 1
	ActivatedCodeCount *int32 `json:"ActivatedCodeCount,omitempty" xml:"ActivatedCodeCount,omitempty"`
	// The maximum number of activation codes that you can apply for.
	//
	// example:
	//
	// 10
	ActivationCodeQuota *int32 `json:"ActivationCodeQuota,omitempty" xml:"ActivationCodeQuota,omitempty"`
	// The Alibaba Cloud order ID (including the virtual order ID).
	//
	// example:
	//
	// 2233****445566
	AliyunOrderId *string `json:"AliyunOrderId,omitempty" xml:"AliyunOrderId,omitempty"`
	// Indicates whether activation codes can be generated without the system identifier.
	//
	// example:
	//
	// false
	AllowEmptySystemIdentifier *bool `json:"AllowEmptySystemIdentifier,omitempty" xml:"AllowEmptySystemIdentifier,omitempty"`
	// The time when the order was created.
	//
	// example:
	//
	// 2024-10-16 16:46:20
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
	// true
	IsVirtualOrder *bool `json:"IsVirtualOrder,omitempty" xml:"IsVirtualOrder,omitempty"`
	// Indicates whether the virtual order is frozen (activation codes cannot be generated for a frozen virtual order).
	//
	// example:
	//
	// false
	IsVirtualOrderFrozen *bool `json:"IsVirtualOrderFrozen,omitempty" xml:"IsVirtualOrderFrozen,omitempty"`
	// The plan type.
	//
	// example:
	//
	// pre_generation_long_term
	PackageType *string `json:"PackageType,omitempty" xml:"PackageType,omitempty"`
	// The validity period of the plan, which is one year (common) or thirty years (long-term).
	//
	// example:
	//
	// 30 years
	PackageValidity *string `json:"PackageValidity,omitempty" xml:"PackageValidity,omitempty"`
	// The purchase channel.
	//
	// example:
	//
	// aliyun_market
	PurchaseChannel *string `json:"PurchaseChannel,omitempty" xml:"PurchaseChannel,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 45D24263-7E3A-4140-9472-************
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the virtual order.
	//
	// example:
	//
	// 2024********483
	VirtualOrderId *string `json:"VirtualOrderId,omitempty" xml:"VirtualOrderId,omitempty"`
}

func (s CreateOrGetVirtualLicenseOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOrGetVirtualLicenseOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetActivatedCodeCount() *int32 {
	return s.ActivatedCodeCount
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetActivationCodeQuota() *int32 {
	return s.ActivationCodeQuota
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetAliyunOrderId() *string {
	return s.AliyunOrderId
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetAllowEmptySystemIdentifier() *bool {
	return s.AllowEmptySystemIdentifier
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetIsVirtualOrder() *bool {
	return s.IsVirtualOrder
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetIsVirtualOrderFrozen() *bool {
	return s.IsVirtualOrderFrozen
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetPackageType() *string {
	return s.PackageType
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetPackageValidity() *string {
	return s.PackageValidity
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetPurchaseChannel() *string {
	return s.PurchaseChannel
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) GetVirtualOrderId() *string {
	return s.VirtualOrderId
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetActivatedCodeCount(v int32) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.ActivatedCodeCount = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetActivationCodeQuota(v int32) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.ActivationCodeQuota = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetAliyunOrderId(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.AliyunOrderId = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetAllowEmptySystemIdentifier(v bool) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.AllowEmptySystemIdentifier = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetGmtCreated(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetGmtModified(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetIsVirtualOrder(v bool) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.IsVirtualOrder = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetIsVirtualOrderFrozen(v bool) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.IsVirtualOrderFrozen = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetPackageType(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.PackageType = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetPackageValidity(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.PackageValidity = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetPurchaseChannel(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.PurchaseChannel = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetRequestId(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) SetVirtualOrderId(v string) *CreateOrGetVirtualLicenseOrderResponseBody {
	s.VirtualOrderId = &v
	return s
}

func (s *CreateOrGetVirtualLicenseOrderResponseBody) Validate() error {
	return dara.Validate(s)
}

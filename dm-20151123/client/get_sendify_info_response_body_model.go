// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSendifyInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBuyResourcePackageUrl(v string) *GetSendifyInfoResponseBody
	GetBuyResourcePackageUrl() *string
	SetBuyUrl(v string) *GetSendifyInfoResponseBody
	GetBuyUrl() *string
	SetDiscount(v string) *GetSendifyInfoResponseBody
	GetDiscount() *string
	SetDowngradeUrl(v string) *GetSendifyInfoResponseBody
	GetDowngradeUrl() *string
	SetExpireTime(v string) *GetSendifyInfoResponseBody
	GetExpireTime() *string
	SetLearnMoreUrl(v string) *GetSendifyInfoResponseBody
	GetLearnMoreUrl() *string
	SetOpened(v bool) *GetSendifyInfoResponseBody
	GetOpened() *bool
	SetProductCode(v string) *GetSendifyInfoResponseBody
	GetProductCode() *string
	SetRenewUrl(v string) *GetSendifyInfoResponseBody
	GetRenewUrl() *string
	SetRequestId(v string) *GetSendifyInfoResponseBody
	GetRequestId() *string
	SetSpecUpgradeUrl(v string) *GetSendifyInfoResponseBody
	GetSpecUpgradeUrl() *string
	SetStatus(v string) *GetSendifyInfoResponseBody
	GetStatus() *string
	SetSupportTrial(v bool) *GetSendifyInfoResponseBody
	GetSupportTrial() *bool
	SetUpgradeUrl(v string) *GetSendifyInfoResponseBody
	GetUpgradeUrl() *string
}

type GetSendifyInfoResponseBody struct {
	// example:
	//
	// http
	BuyResourcePackageUrl *string `json:"BuyResourcePackageUrl,omitempty" xml:"BuyResourcePackageUrl,omitempty"`
	// example:
	//
	// http
	BuyUrl *string `json:"BuyUrl,omitempty" xml:"BuyUrl,omitempty"`
	// example:
	//
	// 0.3
	Discount *string `json:"Discount,omitempty" xml:"Discount,omitempty"`
	// example:
	//
	// http
	DowngradeUrl *string `json:"DowngradeUrl,omitempty" xml:"DowngradeUrl,omitempty"`
	// example:
	//
	// 到期
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// 了解更多
	LearnMoreUrl *string `json:"LearnMoreUrl,omitempty" xml:"LearnMoreUrl,omitempty"`
	Opened       *bool   `json:"Opened,omitempty" xml:"Opened,omitempty"`
	// example:
	//
	// TRIAL
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// example:
	//
	// https
	RenewUrl *string `json:"RenewUrl,omitempty" xml:"RenewUrl,omitempty"`
	// example:
	//
	// 1234
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// http
	SpecUpgradeUrl *string `json:"SpecUpgradeUrl,omitempty" xml:"SpecUpgradeUrl,omitempty"`
	// example:
	//
	// VALID
	Status       *string `json:"Status,omitempty" xml:"Status,omitempty"`
	SupportTrial *bool   `json:"SupportTrial,omitempty" xml:"SupportTrial,omitempty"`
	// example:
	//
	// 升配链接
	UpgradeUrl *string `json:"UpgradeUrl,omitempty" xml:"UpgradeUrl,omitempty"`
}

func (s GetSendifyInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSendifyInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetSendifyInfoResponseBody) GetBuyResourcePackageUrl() *string {
	return s.BuyResourcePackageUrl
}

func (s *GetSendifyInfoResponseBody) GetBuyUrl() *string {
	return s.BuyUrl
}

func (s *GetSendifyInfoResponseBody) GetDiscount() *string {
	return s.Discount
}

func (s *GetSendifyInfoResponseBody) GetDowngradeUrl() *string {
	return s.DowngradeUrl
}

func (s *GetSendifyInfoResponseBody) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *GetSendifyInfoResponseBody) GetLearnMoreUrl() *string {
	return s.LearnMoreUrl
}

func (s *GetSendifyInfoResponseBody) GetOpened() *bool {
	return s.Opened
}

func (s *GetSendifyInfoResponseBody) GetProductCode() *string {
	return s.ProductCode
}

func (s *GetSendifyInfoResponseBody) GetRenewUrl() *string {
	return s.RenewUrl
}

func (s *GetSendifyInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSendifyInfoResponseBody) GetSpecUpgradeUrl() *string {
	return s.SpecUpgradeUrl
}

func (s *GetSendifyInfoResponseBody) GetStatus() *string {
	return s.Status
}

func (s *GetSendifyInfoResponseBody) GetSupportTrial() *bool {
	return s.SupportTrial
}

func (s *GetSendifyInfoResponseBody) GetUpgradeUrl() *string {
	return s.UpgradeUrl
}

func (s *GetSendifyInfoResponseBody) SetBuyResourcePackageUrl(v string) *GetSendifyInfoResponseBody {
	s.BuyResourcePackageUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetBuyUrl(v string) *GetSendifyInfoResponseBody {
	s.BuyUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetDiscount(v string) *GetSendifyInfoResponseBody {
	s.Discount = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetDowngradeUrl(v string) *GetSendifyInfoResponseBody {
	s.DowngradeUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetExpireTime(v string) *GetSendifyInfoResponseBody {
	s.ExpireTime = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetLearnMoreUrl(v string) *GetSendifyInfoResponseBody {
	s.LearnMoreUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetOpened(v bool) *GetSendifyInfoResponseBody {
	s.Opened = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetProductCode(v string) *GetSendifyInfoResponseBody {
	s.ProductCode = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetRenewUrl(v string) *GetSendifyInfoResponseBody {
	s.RenewUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetRequestId(v string) *GetSendifyInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetSpecUpgradeUrl(v string) *GetSendifyInfoResponseBody {
	s.SpecUpgradeUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetStatus(v string) *GetSendifyInfoResponseBody {
	s.Status = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetSupportTrial(v bool) *GetSendifyInfoResponseBody {
	s.SupportTrial = &v
	return s
}

func (s *GetSendifyInfoResponseBody) SetUpgradeUrl(v string) *GetSendifyInfoResponseBody {
	s.UpgradeUrl = &v
	return s
}

func (s *GetSendifyInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

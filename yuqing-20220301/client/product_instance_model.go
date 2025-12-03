// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iProductInstance interface {
	dara.Model
	String() string
	GoString() string
	SetAppCode(v string) *ProductInstance
	GetAppCode() *string
	SetBuyerName(v string) *ProductInstance
	GetBuyerName() *string
	SetBuyerUid(v string) *ProductInstance
	GetBuyerUid() *string
	SetChannel(v string) *ProductInstance
	GetChannel() *string
	SetConfig(v string) *ProductInstance
	GetConfig() *string
	SetEnd(v int64) *ProductInstance
	GetEnd() *int64
	SetInstanceId(v string) *ProductInstance
	GetInstanceId() *string
	SetOrderNo(v string) *ProductInstance
	GetOrderNo() *string
	SetProductCode(v string) *ProductInstance
	GetProductCode() *string
	SetProductSpecCode(v string) *ProductInstance
	GetProductSpecCode() *string
	SetStart(v int64) *ProductInstance
	GetStart() *int64
	SetTenantName(v string) *ProductInstance
	GetTenantName() *string
	SetTenantUid(v string) *ProductInstance
	GetTenantUid() *string
}

type ProductInstance struct {
	// This parameter is required.
	//
	// example:
	//
	// EUWYEEQ
	AppCode *string `json:"appCode,omitempty" xml:"appCode,omitempty"`
	// example:
	//
	// xxx
	BuyerName *string `json:"buyerName,omitempty" xml:"buyerName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ID2343231321
	BuyerUid *string `json:"buyerUid,omitempty" xml:"buyerUid,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ALIYUN
	Channel *string `json:"channel,omitempty" xml:"channel,omitempty"`
	// example:
	//
	// {}
	Config *string `json:"config,omitempty" xml:"config,omitempty"`
	// example:
	//
	// 1640292843231
	End *int64 `json:"end,omitempty" xml:"end,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ID3928389103844
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// D23938474923u42
	OrderNo *string `json:"orderNo,omitempty" xml:"orderNo,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// C394884
	ProductCode *string `json:"productCode,omitempty" xml:"productCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// C847573
	ProductSpecCode *string `json:"productSpecCode,omitempty" xml:"productSpecCode,omitempty"`
	// example:
	//
	// 1640292843231
	Start *int64 `json:"start,omitempty" xml:"start,omitempty"`
	// example:
	//
	// XXX
	TenantName *string `json:"tenantName,omitempty" xml:"tenantName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// UID284747383
	TenantUid *string `json:"tenantUid,omitempty" xml:"tenantUid,omitempty"`
}

func (s ProductInstance) String() string {
	return dara.Prettify(s)
}

func (s ProductInstance) GoString() string {
	return s.String()
}

func (s *ProductInstance) GetAppCode() *string {
	return s.AppCode
}

func (s *ProductInstance) GetBuyerName() *string {
	return s.BuyerName
}

func (s *ProductInstance) GetBuyerUid() *string {
	return s.BuyerUid
}

func (s *ProductInstance) GetChannel() *string {
	return s.Channel
}

func (s *ProductInstance) GetConfig() *string {
	return s.Config
}

func (s *ProductInstance) GetEnd() *int64 {
	return s.End
}

func (s *ProductInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ProductInstance) GetOrderNo() *string {
	return s.OrderNo
}

func (s *ProductInstance) GetProductCode() *string {
	return s.ProductCode
}

func (s *ProductInstance) GetProductSpecCode() *string {
	return s.ProductSpecCode
}

func (s *ProductInstance) GetStart() *int64 {
	return s.Start
}

func (s *ProductInstance) GetTenantName() *string {
	return s.TenantName
}

func (s *ProductInstance) GetTenantUid() *string {
	return s.TenantUid
}

func (s *ProductInstance) SetAppCode(v string) *ProductInstance {
	s.AppCode = &v
	return s
}

func (s *ProductInstance) SetBuyerName(v string) *ProductInstance {
	s.BuyerName = &v
	return s
}

func (s *ProductInstance) SetBuyerUid(v string) *ProductInstance {
	s.BuyerUid = &v
	return s
}

func (s *ProductInstance) SetChannel(v string) *ProductInstance {
	s.Channel = &v
	return s
}

func (s *ProductInstance) SetConfig(v string) *ProductInstance {
	s.Config = &v
	return s
}

func (s *ProductInstance) SetEnd(v int64) *ProductInstance {
	s.End = &v
	return s
}

func (s *ProductInstance) SetInstanceId(v string) *ProductInstance {
	s.InstanceId = &v
	return s
}

func (s *ProductInstance) SetOrderNo(v string) *ProductInstance {
	s.OrderNo = &v
	return s
}

func (s *ProductInstance) SetProductCode(v string) *ProductInstance {
	s.ProductCode = &v
	return s
}

func (s *ProductInstance) SetProductSpecCode(v string) *ProductInstance {
	s.ProductSpecCode = &v
	return s
}

func (s *ProductInstance) SetStart(v int64) *ProductInstance {
	s.Start = &v
	return s
}

func (s *ProductInstance) SetTenantName(v string) *ProductInstance {
	s.TenantName = &v
	return s
}

func (s *ProductInstance) SetTenantUid(v string) *ProductInstance {
	s.TenantUid = &v
	return s
}

func (s *ProductInstance) Validate() error {
	return dara.Validate(s)
}

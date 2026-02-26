// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGoodsShippingNoticeCreateCmd interface {
	dara.Model
	String() string
	GoString() string
	SetCpCode(v string) *GoodsShippingNoticeCreateCmd
	GetCpCode() *string
	SetDisputeId(v string) *GoodsShippingNoticeCreateCmd
	GetDisputeId() *string
	SetLogisticsNo(v string) *GoodsShippingNoticeCreateCmd
	GetLogisticsNo() *string
}

type GoodsShippingNoticeCreateCmd struct {
	// This parameter is required.
	//
	// example:
	//
	// SF
	CpCode *string `json:"cpCode,omitempty" xml:"cpCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 6693****4352
	DisputeId *string `json:"disputeId,omitempty" xml:"disputeId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// SF145****4351
	LogisticsNo *string `json:"logisticsNo,omitempty" xml:"logisticsNo,omitempty"`
}

func (s GoodsShippingNoticeCreateCmd) String() string {
	return dara.Prettify(s)
}

func (s GoodsShippingNoticeCreateCmd) GoString() string {
	return s.String()
}

func (s *GoodsShippingNoticeCreateCmd) GetCpCode() *string {
	return s.CpCode
}

func (s *GoodsShippingNoticeCreateCmd) GetDisputeId() *string {
	return s.DisputeId
}

func (s *GoodsShippingNoticeCreateCmd) GetLogisticsNo() *string {
	return s.LogisticsNo
}

func (s *GoodsShippingNoticeCreateCmd) SetCpCode(v string) *GoodsShippingNoticeCreateCmd {
	s.CpCode = &v
	return s
}

func (s *GoodsShippingNoticeCreateCmd) SetDisputeId(v string) *GoodsShippingNoticeCreateCmd {
	s.DisputeId = &v
	return s
}

func (s *GoodsShippingNoticeCreateCmd) SetLogisticsNo(v string) *GoodsShippingNoticeCreateCmd {
	s.LogisticsNo = &v
	return s
}

func (s *GoodsShippingNoticeCreateCmd) Validate() error {
	return dara.Validate(s)
}

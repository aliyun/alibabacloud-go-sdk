// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckCommodityStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCommodities(v []*CheckCommodityStatusResponseBodyCommodities) *CheckCommodityStatusResponseBody
	GetCommodities() []*CheckCommodityStatusResponseBodyCommodities
	SetRequestId(v string) *CheckCommodityStatusResponseBody
	GetRequestId() *string
}

type CheckCommodityStatusResponseBody struct {
	// The commodity status list. Each item contains the commodity code and indicates whether the commodity capability is enabled for the current account.
	Commodities []*CheckCommodityStatusResponseBodyCommodities `json:"commodities,omitempty" xml:"commodities,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CheckCommodityStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckCommodityStatusResponseBody) GoString() string {
	return s.String()
}

func (s *CheckCommodityStatusResponseBody) GetCommodities() []*CheckCommodityStatusResponseBodyCommodities {
	return s.Commodities
}

func (s *CheckCommodityStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckCommodityStatusResponseBody) SetCommodities(v []*CheckCommodityStatusResponseBodyCommodities) *CheckCommodityStatusResponseBody {
	s.Commodities = v
	return s
}

func (s *CheckCommodityStatusResponseBody) SetRequestId(v string) *CheckCommodityStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckCommodityStatusResponseBody) Validate() error {
	if s.Commodities != nil {
		for _, item := range s.Commodities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckCommodityStatusResponseBodyCommodities struct {
	// The commodity code.
	//
	// example:
	//
	// apigateway_aicredits_public_cn
	CommodityCode *string `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	// Indicates whether the commodity capability is enabled for the current account.
	//
	// example:
	//
	// true
	Enable *bool `json:"enable,omitempty" xml:"enable,omitempty"`
}

func (s CheckCommodityStatusResponseBodyCommodities) String() string {
	return dara.Prettify(s)
}

func (s CheckCommodityStatusResponseBodyCommodities) GoString() string {
	return s.String()
}

func (s *CheckCommodityStatusResponseBodyCommodities) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *CheckCommodityStatusResponseBodyCommodities) GetEnable() *bool {
	return s.Enable
}

func (s *CheckCommodityStatusResponseBodyCommodities) SetCommodityCode(v string) *CheckCommodityStatusResponseBodyCommodities {
	s.CommodityCode = &v
	return s
}

func (s *CheckCommodityStatusResponseBodyCommodities) SetEnable(v bool) *CheckCommodityStatusResponseBodyCommodities {
	s.Enable = &v
	return s
}

func (s *CheckCommodityStatusResponseBodyCommodities) Validate() error {
	return dara.Validate(s)
}

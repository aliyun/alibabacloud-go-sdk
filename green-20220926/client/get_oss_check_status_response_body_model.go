// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOssCheckStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBid(v string) *GetOssCheckStatusResponseBody
	GetBid() *string
	SetBuy(v bool) *GetOssCheckStatusResponseBody
	GetBuy() *bool
	SetCommodityCode(v string) *GetOssCheckStatusResponseBody
	GetCommodityCode() *string
	SetIndebt(v bool) *GetOssCheckStatusResponseBody
	GetIndebt() *bool
	SetRamStatus(v string) *GetOssCheckStatusResponseBody
	GetRamStatus() *string
	SetRequestId(v string) *GetOssCheckStatusResponseBody
	GetRequestId() *string
	SetSlsStatus(v string) *GetOssCheckStatusResponseBody
	GetSlsStatus() *string
}

type GetOssCheckStatusResponseBody struct {
	// Bid.
	//
	// example:
	//
	// 26842
	Bid *string `json:"Bid,omitempty" xml:"Bid,omitempty"`
	// Whether a product has been activated on Alibaba Cloud.
	//
	// example:
	//
	// True
	Buy *bool `json:"Buy,omitempty" xml:"Buy,omitempty"`
	// Commodity code.
	//
	// example:
	//
	// xxx
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Whether there is an outstanding payment.
	//
	// example:
	//
	// False
	Indebt *bool `json:"Indebt,omitempty" xml:"Indebt,omitempty"`
	// Whether internal security is authorized.
	//
	// example:
	//
	// True
	RamStatus *string `json:"RamStatus,omitempty" xml:"RamStatus,omitempty"`
	// ID assigned by the backend, used to uniquely identify a request. Can be used for troubleshooting.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Whether log analysis function is authorized.
	//
	// example:
	//
	// True
	SlsStatus *string `json:"SlsStatus,omitempty" xml:"SlsStatus,omitempty"`
}

func (s GetOssCheckStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOssCheckStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetOssCheckStatusResponseBody) GetBid() *string {
	return s.Bid
}

func (s *GetOssCheckStatusResponseBody) GetBuy() *bool {
	return s.Buy
}

func (s *GetOssCheckStatusResponseBody) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *GetOssCheckStatusResponseBody) GetIndebt() *bool {
	return s.Indebt
}

func (s *GetOssCheckStatusResponseBody) GetRamStatus() *string {
	return s.RamStatus
}

func (s *GetOssCheckStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOssCheckStatusResponseBody) GetSlsStatus() *string {
	return s.SlsStatus
}

func (s *GetOssCheckStatusResponseBody) SetBid(v string) *GetOssCheckStatusResponseBody {
	s.Bid = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetBuy(v bool) *GetOssCheckStatusResponseBody {
	s.Buy = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetCommodityCode(v string) *GetOssCheckStatusResponseBody {
	s.CommodityCode = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetIndebt(v bool) *GetOssCheckStatusResponseBody {
	s.Indebt = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetRamStatus(v string) *GetOssCheckStatusResponseBody {
	s.RamStatus = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetRequestId(v string) *GetOssCheckStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) SetSlsStatus(v string) *GetOssCheckStatusResponseBody {
	s.SlsStatus = &v
	return s
}

func (s *GetOssCheckStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHotelOrderDetailResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetHotelOrderDetailResponseBody
	GetCode() *int32
	SetMessage(v string) *GetHotelOrderDetailResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHotelOrderDetailResponseBody
	GetRequestId() *string
	SetResult(v []*GetHotelOrderDetailResponseBodyResult) *GetHotelOrderDetailResponseBody
	GetResult() []*GetHotelOrderDetailResponseBodyResult
}

type GetHotelOrderDetailResponseBody struct {
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6F579407-13C4-1708-AFA2-B657BE5FE8F5
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*GetHotelOrderDetailResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetHotelOrderDetailResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailResponseBody) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetHotelOrderDetailResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHotelOrderDetailResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHotelOrderDetailResponseBody) GetResult() []*GetHotelOrderDetailResponseBodyResult {
	return s.Result
}

func (s *GetHotelOrderDetailResponseBody) SetCode(v int32) *GetHotelOrderDetailResponseBody {
	s.Code = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetMessage(v string) *GetHotelOrderDetailResponseBody {
	s.Message = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetRequestId(v string) *GetHotelOrderDetailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHotelOrderDetailResponseBody) SetResult(v []*GetHotelOrderDetailResponseBodyResult) *GetHotelOrderDetailResponseBody {
	s.Result = v
	return s
}

func (s *GetHotelOrderDetailResponseBody) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHotelOrderDetailResponseBodyResult struct {
	// example:
	//
	// 200
	ApplyAmt *int64 `json:"ApplyAmt,omitempty" xml:"ApplyAmt,omitempty"`
	// example:
	//
	// 1659952892000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// example:
	//
	// https://ailabsaicloudservice.alicdn.com/hotel/icon/jiudianmianban_fuwushangpintu/weixiu/dianqilei/chuanglian.png
	ItemUrl *string `json:"ItemUrl,omitempty" xml:"ItemUrl,omitempty"`
	// example:
	//
	// 窗帘
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 1
	Quantity *int64 `json:"Quantity,omitempty" xml:"Quantity,omitempty"`
}

func (s GetHotelOrderDetailResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetHotelOrderDetailResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetHotelOrderDetailResponseBodyResult) GetApplyAmt() *int64 {
	return s.ApplyAmt
}

func (s *GetHotelOrderDetailResponseBodyResult) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetHotelOrderDetailResponseBodyResult) GetItemUrl() *string {
	return s.ItemUrl
}

func (s *GetHotelOrderDetailResponseBodyResult) GetName() *string {
	return s.Name
}

func (s *GetHotelOrderDetailResponseBodyResult) GetQuantity() *int64 {
	return s.Quantity
}

func (s *GetHotelOrderDetailResponseBodyResult) SetApplyAmt(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.ApplyAmt = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetGmtCreate(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.GmtCreate = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetItemUrl(v string) *GetHotelOrderDetailResponseBodyResult {
	s.ItemUrl = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetName(v string) *GetHotelOrderDetailResponseBodyResult {
	s.Name = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) SetQuantity(v int64) *GetHotelOrderDetailResponseBodyResult {
	s.Quantity = &v
	return s
}

func (s *GetHotelOrderDetailResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

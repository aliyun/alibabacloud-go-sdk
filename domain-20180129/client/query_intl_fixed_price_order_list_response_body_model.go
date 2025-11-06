// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryIntlFixedPriceOrderListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v *QueryIntlFixedPriceOrderListResponseBodyModule) *QueryIntlFixedPriceOrderListResponseBody
	GetModule() *QueryIntlFixedPriceOrderListResponseBodyModule
	SetRequestId(v string) *QueryIntlFixedPriceOrderListResponseBody
	GetRequestId() *string
}

type QueryIntlFixedPriceOrderListResponseBody struct {
	Module *QueryIntlFixedPriceOrderListResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Struct"`
	// example:
	//
	// D6CB3623-4726-4947-AC2B-2C6E673B447C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryIntlFixedPriceOrderListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryIntlFixedPriceOrderListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryIntlFixedPriceOrderListResponseBody) GetModule() *QueryIntlFixedPriceOrderListResponseBodyModule {
	return s.Module
}

func (s *QueryIntlFixedPriceOrderListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryIntlFixedPriceOrderListResponseBody) SetModule(v *QueryIntlFixedPriceOrderListResponseBodyModule) *QueryIntlFixedPriceOrderListResponseBody {
	s.Module = v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBody) SetRequestId(v string) *QueryIntlFixedPriceOrderListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBody) Validate() error {
	if s.Module != nil {
		if err := s.Module.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryIntlFixedPriceOrderListResponseBodyModule struct {
	CurrentPageNum *int32                                                `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryIntlFixedPriceOrderListResponseBodyModuleData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	PageSize       *int32                                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	TotalItemNum   *int32                                                `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	TotalPageNum   *int32                                                `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryIntlFixedPriceOrderListResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryIntlFixedPriceOrderListResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) GetData() []*QueryIntlFixedPriceOrderListResponseBodyModuleData {
	return s.Data
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) SetCurrentPageNum(v int32) *QueryIntlFixedPriceOrderListResponseBodyModule {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) SetData(v []*QueryIntlFixedPriceOrderListResponseBodyModuleData) *QueryIntlFixedPriceOrderListResponseBodyModule {
	s.Data = v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) SetPageSize(v int32) *QueryIntlFixedPriceOrderListResponseBodyModule {
	s.PageSize = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) SetTotalItemNum(v int32) *QueryIntlFixedPriceOrderListResponseBodyModule {
	s.TotalItemNum = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) SetTotalPageNum(v int32) *QueryIntlFixedPriceOrderListResponseBodyModule {
	s.TotalPageNum = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModule) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryIntlFixedPriceOrderListResponseBodyModuleData struct {
	BizId      *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	CreateTime *int64  `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Domain     *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	OrderType  *int32  `json:"OrderType,omitempty" xml:"OrderType,omitempty"`
	Price      *int64  `json:"Price,omitempty" xml:"Price,omitempty"`
	Status     *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	UpdateTime *int64  `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserId     *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s QueryIntlFixedPriceOrderListResponseBodyModuleData) String() string {
	return dara.Prettify(s)
}

func (s QueryIntlFixedPriceOrderListResponseBodyModuleData) GoString() string {
	return s.String()
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetBizId() *string {
	return s.BizId
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetCreateTime() *int64 {
	return s.CreateTime
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetDomain() *string {
	return s.Domain
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetOrderType() *int32 {
	return s.OrderType
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetPrice() *int64 {
	return s.Price
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetStatus() *int32 {
	return s.Status
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) GetUserId() *string {
	return s.UserId
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetBizId(v string) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.BizId = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetCreateTime(v int64) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.CreateTime = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetDomain(v string) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.Domain = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetOrderType(v int32) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.OrderType = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetPrice(v int64) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.Price = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetStatus(v int32) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.Status = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetUpdateTime(v int64) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.UpdateTime = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) SetUserId(v string) *QueryIntlFixedPriceOrderListResponseBodyModuleData {
	s.UserId = &v
	return s
}

func (s *QueryIntlFixedPriceOrderListResponseBodyModuleData) Validate() error {
	return dara.Validate(s)
}

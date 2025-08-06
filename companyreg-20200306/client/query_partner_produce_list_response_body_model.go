// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPartnerProduceListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int64) *QueryPartnerProduceListResponseBody
	GetCurrentPageNum() *int64
	SetData(v []*QueryPartnerProduceListResponseBodyData) *QueryPartnerProduceListResponseBody
	GetData() []*QueryPartnerProduceListResponseBodyData
	SetPageSize(v int64) *QueryPartnerProduceListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *QueryPartnerProduceListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int64) *QueryPartnerProduceListResponseBody
	GetTotalItemNum() *int64
	SetTotalPageNum(v int64) *QueryPartnerProduceListResponseBody
	GetTotalPageNum() *int64
}

type QueryPartnerProduceListResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int64                                     `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryPartnerProduceListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 1A13ABB5-7649-5031-B55C-D2E38F9F189D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 34
	TotalItemNum *int64 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 1
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryPartnerProduceListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerProduceListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListResponseBody) GetCurrentPageNum() *int64 {
	return s.CurrentPageNum
}

func (s *QueryPartnerProduceListResponseBody) GetData() []*QueryPartnerProduceListResponseBodyData {
	return s.Data
}

func (s *QueryPartnerProduceListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPartnerProduceListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPartnerProduceListResponseBody) GetTotalItemNum() *int64 {
	return s.TotalItemNum
}

func (s *QueryPartnerProduceListResponseBody) GetTotalPageNum() *int64 {
	return s.TotalPageNum
}

func (s *QueryPartnerProduceListResponseBody) SetCurrentPageNum(v int64) *QueryPartnerProduceListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetData(v []*QueryPartnerProduceListResponseBodyData) *QueryPartnerProduceListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetPageSize(v int64) *QueryPartnerProduceListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetRequestId(v string) *QueryPartnerProduceListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetTotalItemNum(v int64) *QueryPartnerProduceListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) SetTotalPageNum(v int64) *QueryPartnerProduceListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryPartnerProduceListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPartnerProduceListResponseBodyData struct {
	// example:
	//
	// P20211216204717000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.cdn
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 18600000001
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
}

func (s QueryPartnerProduceListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerProduceListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *QueryPartnerProduceListResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *QueryPartnerProduceListResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *QueryPartnerProduceListResponseBodyData) SetBizId(v string) *QueryPartnerProduceListResponseBodyData {
	s.BizId = &v
	return s
}

func (s *QueryPartnerProduceListResponseBodyData) SetBizType(v string) *QueryPartnerProduceListResponseBodyData {
	s.BizType = &v
	return s
}

func (s *QueryPartnerProduceListResponseBodyData) SetMobile(v string) *QueryPartnerProduceListResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryPartnerProduceListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPartnerIntentionListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int64) *QueryPartnerIntentionListResponseBody
	GetCurrentPageNum() *int64
	SetData(v []*QueryPartnerIntentionListResponseBodyData) *QueryPartnerIntentionListResponseBody
	GetData() []*QueryPartnerIntentionListResponseBodyData
	SetPageSize(v int64) *QueryPartnerIntentionListResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *QueryPartnerIntentionListResponseBody
	GetRequestId() *string
	SetTotalItemNum(v int64) *QueryPartnerIntentionListResponseBody
	GetTotalItemNum() *int64
	SetTotalPageNum(v int64) *QueryPartnerIntentionListResponseBody
	GetTotalPageNum() *int64
}

type QueryPartnerIntentionListResponseBody struct {
	// example:
	//
	// 2
	CurrentPageNum *int64                                       `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*QueryPartnerIntentionListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 78E9DC76-7DFD-5975-99B0-4A95E8A92F5D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 319
	TotalItemNum *int64 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 13
	TotalPageNum *int64 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s QueryPartnerIntentionListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerIntentionListResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListResponseBody) GetCurrentPageNum() *int64 {
	return s.CurrentPageNum
}

func (s *QueryPartnerIntentionListResponseBody) GetData() []*QueryPartnerIntentionListResponseBodyData {
	return s.Data
}

func (s *QueryPartnerIntentionListResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *QueryPartnerIntentionListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPartnerIntentionListResponseBody) GetTotalItemNum() *int64 {
	return s.TotalItemNum
}

func (s *QueryPartnerIntentionListResponseBody) GetTotalPageNum() *int64 {
	return s.TotalPageNum
}

func (s *QueryPartnerIntentionListResponseBody) SetCurrentPageNum(v int64) *QueryPartnerIntentionListResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetData(v []*QueryPartnerIntentionListResponseBodyData) *QueryPartnerIntentionListResponseBody {
	s.Data = v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetPageSize(v int64) *QueryPartnerIntentionListResponseBody {
	s.PageSize = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetRequestId(v string) *QueryPartnerIntentionListResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetTotalItemNum(v int64) *QueryPartnerIntentionListResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) SetTotalPageNum(v int64) *QueryPartnerIntentionListResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBody) Validate() error {
	return dara.Validate(s)
}

type QueryPartnerIntentionListResponseBodyData struct {
	// example:
	//
	// I20211117092704000001
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// example:
	//
	// esp.wangwen
	BizType *string `json:"BizType,omitempty" xml:"BizType,omitempty"`
	// example:
	//
	// 18700000003
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s QueryPartnerIntentionListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerIntentionListResponseBodyData) GoString() string {
	return s.String()
}

func (s *QueryPartnerIntentionListResponseBodyData) GetBizId() *string {
	return s.BizId
}

func (s *QueryPartnerIntentionListResponseBodyData) GetBizType() *string {
	return s.BizType
}

func (s *QueryPartnerIntentionListResponseBodyData) GetMobile() *string {
	return s.Mobile
}

func (s *QueryPartnerIntentionListResponseBodyData) SetBizId(v string) *QueryPartnerIntentionListResponseBodyData {
	s.BizId = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBodyData) SetBizType(v string) *QueryPartnerIntentionListResponseBodyData {
	s.BizType = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBodyData) SetMobile(v string) *QueryPartnerIntentionListResponseBodyData {
	s.Mobile = &v
	return s
}

func (s *QueryPartnerIntentionListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

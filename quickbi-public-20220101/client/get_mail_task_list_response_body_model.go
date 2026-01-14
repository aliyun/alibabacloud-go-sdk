// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMailTaskListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetMailTaskListResponseBody
	GetRequestId() *string
	SetResult(v *GetMailTaskListResponseBodyResult) *GetMailTaskListResponseBody
	GetResult() *GetMailTaskListResponseBodyResult
	SetSuccess(v bool) *GetMailTaskListResponseBody
	GetSuccess() *bool
}

type GetMailTaskListResponseBody struct {
	// example:
	//
	// 0c52************8e1952a3
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *GetMailTaskListResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetMailTaskListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskListResponseBody) GoString() string {
	return s.String()
}

func (s *GetMailTaskListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMailTaskListResponseBody) GetResult() *GetMailTaskListResponseBodyResult {
	return s.Result
}

func (s *GetMailTaskListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetMailTaskListResponseBody) SetRequestId(v string) *GetMailTaskListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMailTaskListResponseBody) SetResult(v *GetMailTaskListResponseBodyResult) *GetMailTaskListResponseBody {
	s.Result = v
	return s
}

func (s *GetMailTaskListResponseBody) SetSuccess(v bool) *GetMailTaskListResponseBody {
	s.Success = &v
	return s
}

func (s *GetMailTaskListResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetMailTaskListResponseBodyResult struct {
	Data []*GetMailTaskListResponseBodyResultData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// null
	Next *int32 `json:"Next,omitempty" xml:"Next,omitempty"`
	// example:
	//
	// 1
	PageNum *int32 `json:"PageNum,omitempty" xml:"PageNum,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// null
	Pre *int32 `json:"Pre,omitempty" xml:"Pre,omitempty"`
	// example:
	//
	// 1
	TotalNum *int32 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
	// example:
	//
	// 1
	TotalPages *int32 `json:"TotalPages,omitempty" xml:"TotalPages,omitempty"`
}

func (s GetMailTaskListResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskListResponseBodyResult) GoString() string {
	return s.String()
}

func (s *GetMailTaskListResponseBodyResult) GetData() []*GetMailTaskListResponseBodyResultData {
	return s.Data
}

func (s *GetMailTaskListResponseBodyResult) GetNext() *int32 {
	return s.Next
}

func (s *GetMailTaskListResponseBodyResult) GetPageNum() *int32 {
	return s.PageNum
}

func (s *GetMailTaskListResponseBodyResult) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetMailTaskListResponseBodyResult) GetPre() *int32 {
	return s.Pre
}

func (s *GetMailTaskListResponseBodyResult) GetTotalNum() *int32 {
	return s.TotalNum
}

func (s *GetMailTaskListResponseBodyResult) GetTotalPages() *int32 {
	return s.TotalPages
}

func (s *GetMailTaskListResponseBodyResult) SetData(v []*GetMailTaskListResponseBodyResultData) *GetMailTaskListResponseBodyResult {
	s.Data = v
	return s
}

func (s *GetMailTaskListResponseBodyResult) SetNext(v int32) *GetMailTaskListResponseBodyResult {
	s.Next = &v
	return s
}

func (s *GetMailTaskListResponseBodyResult) SetPageNum(v int32) *GetMailTaskListResponseBodyResult {
	s.PageNum = &v
	return s
}

func (s *GetMailTaskListResponseBodyResult) SetPageSize(v int32) *GetMailTaskListResponseBodyResult {
	s.PageSize = &v
	return s
}

func (s *GetMailTaskListResponseBodyResult) SetPre(v int32) *GetMailTaskListResponseBodyResult {
	s.Pre = &v
	return s
}

func (s *GetMailTaskListResponseBodyResult) SetTotalNum(v int32) *GetMailTaskListResponseBodyResult {
	s.TotalNum = &v
	return s
}

func (s *GetMailTaskListResponseBodyResult) SetTotalPages(v int32) *GetMailTaskListResponseBodyResult {
	s.TotalPages = &v
	return s
}

func (s *GetMailTaskListResponseBodyResult) Validate() error {
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

type GetMailTaskListResponseBodyResultData struct {
	// example:
	//
	// test
	BizOwnerName *string `json:"BizOwnerName,omitempty" xml:"BizOwnerName,omitempty"`
	// example:
	//
	// asag****2423
	BizOwnerUserId *string `json:"BizOwnerUserId,omitempty" xml:"BizOwnerUserId,omitempty"`
	// example:
	//
	// 2342526
	MailId *string `json:"MailId,omitempty" xml:"MailId,omitempty"`
	// example:
	//
	// true
	Paused *bool `json:"Paused,omitempty" xml:"Paused,omitempty"`
	// example:
	//
	// test
	SubscribeName *string `json:"SubscribeName,omitempty" xml:"SubscribeName,omitempty"`
}

func (s GetMailTaskListResponseBodyResultData) String() string {
	return dara.Prettify(s)
}

func (s GetMailTaskListResponseBodyResultData) GoString() string {
	return s.String()
}

func (s *GetMailTaskListResponseBodyResultData) GetBizOwnerName() *string {
	return s.BizOwnerName
}

func (s *GetMailTaskListResponseBodyResultData) GetBizOwnerUserId() *string {
	return s.BizOwnerUserId
}

func (s *GetMailTaskListResponseBodyResultData) GetMailId() *string {
	return s.MailId
}

func (s *GetMailTaskListResponseBodyResultData) GetPaused() *bool {
	return s.Paused
}

func (s *GetMailTaskListResponseBodyResultData) GetSubscribeName() *string {
	return s.SubscribeName
}

func (s *GetMailTaskListResponseBodyResultData) SetBizOwnerName(v string) *GetMailTaskListResponseBodyResultData {
	s.BizOwnerName = &v
	return s
}

func (s *GetMailTaskListResponseBodyResultData) SetBizOwnerUserId(v string) *GetMailTaskListResponseBodyResultData {
	s.BizOwnerUserId = &v
	return s
}

func (s *GetMailTaskListResponseBodyResultData) SetMailId(v string) *GetMailTaskListResponseBodyResultData {
	s.MailId = &v
	return s
}

func (s *GetMailTaskListResponseBodyResultData) SetPaused(v bool) *GetMailTaskListResponseBodyResultData {
	s.Paused = &v
	return s
}

func (s *GetMailTaskListResponseBodyResultData) SetSubscribeName(v string) *GetMailTaskListResponseBodyResultData {
	s.SubscribeName = &v
	return s
}

func (s *GetMailTaskListResponseBodyResultData) Validate() error {
	return dara.Validate(s)
}

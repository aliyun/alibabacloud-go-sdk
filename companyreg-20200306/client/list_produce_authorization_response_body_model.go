// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProduceAuthorizationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentPageNum(v int32) *ListProduceAuthorizationResponseBody
	GetCurrentPageNum() *int32
	SetData(v []*ListProduceAuthorizationResponseBodyData) *ListProduceAuthorizationResponseBody
	GetData() []*ListProduceAuthorizationResponseBodyData
	SetPageSize(v int32) *ListProduceAuthorizationResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListProduceAuthorizationResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProduceAuthorizationResponseBody
	GetSuccess() *bool
	SetTotalItemNum(v int32) *ListProduceAuthorizationResponseBody
	GetTotalItemNum() *int32
	SetTotalPageNum(v int32) *ListProduceAuthorizationResponseBody
	GetTotalPageNum() *int32
}

type ListProduceAuthorizationResponseBody struct {
	// example:
	//
	// 1
	CurrentPageNum *int32                                      `json:"CurrentPageNum,omitempty" xml:"CurrentPageNum,omitempty"`
	Data           []*ListProduceAuthorizationResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 20
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 10AAC56B-C512-5860-9A9E-B949431E7174
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 292
	TotalItemNum *int32 `json:"TotalItemNum,omitempty" xml:"TotalItemNum,omitempty"`
	// example:
	//
	// 27
	TotalPageNum *int32 `json:"TotalPageNum,omitempty" xml:"TotalPageNum,omitempty"`
}

func (s ListProduceAuthorizationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProduceAuthorizationResponseBody) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationResponseBody) GetCurrentPageNum() *int32 {
	return s.CurrentPageNum
}

func (s *ListProduceAuthorizationResponseBody) GetData() []*ListProduceAuthorizationResponseBodyData {
	return s.Data
}

func (s *ListProduceAuthorizationResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListProduceAuthorizationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProduceAuthorizationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProduceAuthorizationResponseBody) GetTotalItemNum() *int32 {
	return s.TotalItemNum
}

func (s *ListProduceAuthorizationResponseBody) GetTotalPageNum() *int32 {
	return s.TotalPageNum
}

func (s *ListProduceAuthorizationResponseBody) SetCurrentPageNum(v int32) *ListProduceAuthorizationResponseBody {
	s.CurrentPageNum = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetData(v []*ListProduceAuthorizationResponseBodyData) *ListProduceAuthorizationResponseBody {
	s.Data = v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetPageSize(v int32) *ListProduceAuthorizationResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetRequestId(v string) *ListProduceAuthorizationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetSuccess(v bool) *ListProduceAuthorizationResponseBody {
	s.Success = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetTotalItemNum(v int32) *ListProduceAuthorizationResponseBody {
	s.TotalItemNum = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) SetTotalPageNum(v int32) *ListProduceAuthorizationResponseBody {
	s.TotalPageNum = &v
	return s
}

func (s *ListProduceAuthorizationResponseBody) Validate() error {
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

type ListProduceAuthorizationResponseBodyData struct {
	// example:
	//
	// 12195411612139999
	AuthorizedUserId *string `json:"AuthorizedUserId,omitempty" xml:"AuthorizedUserId,omitempty"`
	// example:
	//
	// test@alibaba-inc.com
	AuthorizedUserName *string `json:"AuthorizedUserName,omitempty" xml:"AuthorizedUserName,omitempty"`
}

func (s ListProduceAuthorizationResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListProduceAuthorizationResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListProduceAuthorizationResponseBodyData) GetAuthorizedUserId() *string {
	return s.AuthorizedUserId
}

func (s *ListProduceAuthorizationResponseBodyData) GetAuthorizedUserName() *string {
	return s.AuthorizedUserName
}

func (s *ListProduceAuthorizationResponseBodyData) SetAuthorizedUserId(v string) *ListProduceAuthorizationResponseBodyData {
	s.AuthorizedUserId = &v
	return s
}

func (s *ListProduceAuthorizationResponseBodyData) SetAuthorizedUserName(v string) *ListProduceAuthorizationResponseBodyData {
	s.AuthorizedUserName = &v
	return s
}

func (s *ListProduceAuthorizationResponseBodyData) Validate() error {
	return dara.Validate(s)
}

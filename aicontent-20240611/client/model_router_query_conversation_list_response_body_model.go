// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryConversationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModelRouterQueryConversationListResponseBodyData) *ModelRouterQueryConversationListResponseBody
	GetData() *ModelRouterQueryConversationListResponseBodyData
	SetErrCode(v string) *ModelRouterQueryConversationListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryConversationListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryConversationListResponseBody
	GetHttpStatusCode() *int32
	SetRequestId(v string) *ModelRouterQueryConversationListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryConversationListResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryConversationListResponseBody struct {
	// example:
	//
	// []
	Data *ModelRouterQueryConversationListResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryConversationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationListResponseBody) GetData() *ModelRouterQueryConversationListResponseBodyData {
	return s.Data
}

func (s *ModelRouterQueryConversationListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryConversationListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryConversationListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryConversationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryConversationListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryConversationListResponseBody) SetData(v *ModelRouterQueryConversationListResponseBodyData) *ModelRouterQueryConversationListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetErrCode(v string) *ModelRouterQueryConversationListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetErrMessage(v string) *ModelRouterQueryConversationListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryConversationListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetRequestId(v string) *ModelRouterQueryConversationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetSuccess(v bool) *ModelRouterQueryConversationListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelRouterQueryConversationListResponseBodyData struct {
	List      []*ConversationDTO `json:"list,omitempty" xml:"list,omitempty" type:"Repeated"`
	MaxResult *string            `json:"maxResult,omitempty" xml:"maxResult,omitempty"`
	// example:
	//
	// None
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	Page *int32 `json:"page,omitempty" xml:"page,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// 0
	Total *int32 `json:"total,omitempty" xml:"total,omitempty"`
}

func (s ModelRouterQueryConversationListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationListResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationListResponseBodyData) GetList() []*ConversationDTO {
	return s.List
}

func (s *ModelRouterQueryConversationListResponseBodyData) GetMaxResult() *string {
	return s.MaxResult
}

func (s *ModelRouterQueryConversationListResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryConversationListResponseBodyData) GetPage() *int32 {
	return s.Page
}

func (s *ModelRouterQueryConversationListResponseBodyData) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryConversationListResponseBodyData) GetTotal() *int32 {
	return s.Total
}

func (s *ModelRouterQueryConversationListResponseBodyData) SetList(v []*ConversationDTO) *ModelRouterQueryConversationListResponseBodyData {
	s.List = v
	return s
}

func (s *ModelRouterQueryConversationListResponseBodyData) SetMaxResult(v string) *ModelRouterQueryConversationListResponseBodyData {
	s.MaxResult = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBodyData) SetNextToken(v string) *ModelRouterQueryConversationListResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBodyData) SetPage(v int32) *ModelRouterQueryConversationListResponseBodyData {
	s.Page = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBodyData) SetPageSize(v int32) *ModelRouterQueryConversationListResponseBodyData {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBodyData) SetTotal(v int32) *ModelRouterQueryConversationListResponseBodyData {
	s.Total = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

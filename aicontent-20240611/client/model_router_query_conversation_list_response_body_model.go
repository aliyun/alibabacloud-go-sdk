// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryConversationListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ConversationDTO) *ModelRouterQueryConversationListResponseBody
	GetData() []*ConversationDTO
	SetErrCode(v string) *ModelRouterQueryConversationListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryConversationListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryConversationListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryConversationListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryConversationListResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryConversationListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryConversationListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryConversationListResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryConversationListResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryConversationListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryConversationListResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryConversationListResponseBody struct {
	// example:
	//
	// []
	Data []*ConversationDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// maxResults
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// nextToken
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// example:
	//
	// 1
	PageIndex *int32 `json:"pageIndex,omitempty" xml:"pageIndex,omitempty"`
	// example:
	//
	// 10
	PageSize *int32 `json:"pageSize,omitempty" xml:"pageSize,omitempty"`
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// skip
	//
	// example:
	//
	// 10
	Skip *int32 `json:"skip,omitempty" xml:"skip,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
	// example:
	//
	// 100
	TotalCount *int32 `json:"totalCount,omitempty" xml:"totalCount,omitempty"`
}

func (s ModelRouterQueryConversationListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryConversationListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryConversationListResponseBody) GetData() []*ConversationDTO {
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

func (s *ModelRouterQueryConversationListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryConversationListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryConversationListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryConversationListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryConversationListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryConversationListResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryConversationListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryConversationListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryConversationListResponseBody) SetData(v []*ConversationDTO) *ModelRouterQueryConversationListResponseBody {
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

func (s *ModelRouterQueryConversationListResponseBody) SetMaxResults(v int32) *ModelRouterQueryConversationListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetNextToken(v string) *ModelRouterQueryConversationListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetPageIndex(v int32) *ModelRouterQueryConversationListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetPageSize(v int32) *ModelRouterQueryConversationListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetRequestId(v string) *ModelRouterQueryConversationListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetSkip(v int32) *ModelRouterQueryConversationListResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetSuccess(v bool) *ModelRouterQueryConversationListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) SetTotalCount(v int32) *ModelRouterQueryConversationListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryConversationListResponseBody) Validate() error {
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

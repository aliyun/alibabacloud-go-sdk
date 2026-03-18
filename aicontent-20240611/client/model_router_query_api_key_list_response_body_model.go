// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryApiKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ApiKeyDTO) *ModelRouterQueryApiKeyListResponseBody
	GetData() []*ApiKeyDTO
	SetErrCode(v string) *ModelRouterQueryApiKeyListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryApiKeyListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryApiKeyListResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryApiKeyListResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryApiKeyListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryApiKeyListResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryApiKeyListResponseBody struct {
	// example:
	//
	// []
	Data []*ApiKeyDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ModelRouterQueryApiKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryApiKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetData() []*ApiKeyDTO {
	return s.Data
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryApiKeyListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetData(v []*ApiKeyDTO) *ModelRouterQueryApiKeyListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetErrCode(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetErrMessage(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetMaxResults(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetNextToken(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetPageIndex(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetPageSize(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetRequestId(v string) *ModelRouterQueryApiKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetSkip(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetSuccess(v bool) *ModelRouterQueryApiKeyListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) SetTotalCount(v int32) *ModelRouterQueryApiKeyListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryApiKeyListResponseBody) Validate() error {
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

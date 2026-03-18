// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosProvidersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *ModelRouterQueryNacosProvidersResponseBody
	GetData() []*string
	SetErrCode(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryNacosProvidersResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryNacosProvidersResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryNacosProvidersResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryNacosProvidersResponseBody struct {
	// example:
	//
	// []
	Data []*string `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ModelRouterQueryNacosProvidersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosProvidersResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetData() []*string {
	return s.Data
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryNacosProvidersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetData(v []*string) *ModelRouterQueryNacosProvidersResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetErrCode(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetErrMessage(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetMaxResults(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetNextToken(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetPageIndex(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetPageSize(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetRequestId(v string) *ModelRouterQueryNacosProvidersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetSkip(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetSuccess(v bool) *ModelRouterQueryNacosProvidersResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) SetTotalCount(v int32) *ModelRouterQueryNacosProvidersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryNacosProvidersResponseBody) Validate() error {
	return dara.Validate(s)
}

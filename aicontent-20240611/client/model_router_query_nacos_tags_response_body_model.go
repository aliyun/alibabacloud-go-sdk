// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryNacosTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*string) *ModelRouterQueryNacosTagsResponseBody
	GetData() []*string
	SetErrCode(v string) *ModelRouterQueryNacosTagsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryNacosTagsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryNacosTagsResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryNacosTagsResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryNacosTagsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryNacosTagsResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryNacosTagsResponseBody struct {
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

func (s ModelRouterQueryNacosTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryNacosTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetData() []*string {
	return s.Data
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryNacosTagsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetData(v []*string) *ModelRouterQueryNacosTagsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetErrCode(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetErrMessage(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetMaxResults(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetNextToken(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetPageIndex(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetPageSize(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetRequestId(v string) *ModelRouterQueryNacosTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetSkip(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetSuccess(v bool) *ModelRouterQueryNacosTagsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) SetTotalCount(v int32) *ModelRouterQueryNacosTagsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryNacosTagsResponseBody) Validate() error {
	return dara.Validate(s)
}

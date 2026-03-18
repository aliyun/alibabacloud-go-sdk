// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryModelListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ModelDTO) *ModelRouterQueryModelListResponseBody
	GetData() []*ModelDTO
	SetErrCode(v string) *ModelRouterQueryModelListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryModelListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryModelListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryModelListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryModelListResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryModelListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryModelListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryModelListResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryModelListResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryModelListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryModelListResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryModelListResponseBody struct {
	// example:
	//
	// []
	Data []*ModelDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ModelRouterQueryModelListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryModelListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryModelListResponseBody) GetData() []*ModelDTO {
	return s.Data
}

func (s *ModelRouterQueryModelListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryModelListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryModelListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryModelListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryModelListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryModelListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryModelListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryModelListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryModelListResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryModelListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryModelListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryModelListResponseBody) SetData(v []*ModelDTO) *ModelRouterQueryModelListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetErrCode(v string) *ModelRouterQueryModelListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetErrMessage(v string) *ModelRouterQueryModelListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryModelListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetMaxResults(v int32) *ModelRouterQueryModelListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetNextToken(v string) *ModelRouterQueryModelListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetPageIndex(v int32) *ModelRouterQueryModelListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetPageSize(v int32) *ModelRouterQueryModelListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetRequestId(v string) *ModelRouterQueryModelListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetSkip(v int32) *ModelRouterQueryModelListResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetSuccess(v bool) *ModelRouterQueryModelListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) SetTotalCount(v int32) *ModelRouterQueryModelListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryModelListResponseBody) Validate() error {
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

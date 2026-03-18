// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryClientListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ClientDTO) *ModelRouterQueryClientListResponseBody
	GetData() []*ClientDTO
	SetErrCode(v string) *ModelRouterQueryClientListResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryClientListResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryClientListResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryClientListResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryClientListResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryClientListResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryClientListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryClientListResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryClientListResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryClientListResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryClientListResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryClientListResponseBody struct {
	// example:
	//
	// []
	Data []*ClientDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ModelRouterQueryClientListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryClientListResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryClientListResponseBody) GetData() []*ClientDTO {
	return s.Data
}

func (s *ModelRouterQueryClientListResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryClientListResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryClientListResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryClientListResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryClientListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryClientListResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryClientListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryClientListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryClientListResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryClientListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryClientListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryClientListResponseBody) SetData(v []*ClientDTO) *ModelRouterQueryClientListResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetErrCode(v string) *ModelRouterQueryClientListResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetErrMessage(v string) *ModelRouterQueryClientListResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryClientListResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetMaxResults(v int32) *ModelRouterQueryClientListResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetNextToken(v string) *ModelRouterQueryClientListResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetPageIndex(v int32) *ModelRouterQueryClientListResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetPageSize(v int32) *ModelRouterQueryClientListResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetRequestId(v string) *ModelRouterQueryClientListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetSkip(v int32) *ModelRouterQueryClientListResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetSuccess(v bool) *ModelRouterQueryClientListResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) SetTotalCount(v int32) *ModelRouterQueryClientListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryClientListResponseBody) Validate() error {
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

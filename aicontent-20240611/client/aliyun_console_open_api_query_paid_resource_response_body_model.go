// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAliyunConsoleOpenApiQueryPaidResourceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*PaidResourceDTO) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetData() []*PaidResourceDTO
	SetErrCode(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetRequestId() *string
	SetSkip(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody
	GetTotalCount() *int32
}

type AliyunConsoleOpenApiQueryPaidResourceResponseBody struct {
	// example:
	//
	// []
	Data []*PaidResourceDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s AliyunConsoleOpenApiQueryPaidResourceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AliyunConsoleOpenApiQueryPaidResourceResponseBody) GoString() string {
	return s.String()
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetData() []*PaidResourceDTO {
	return s.Data
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetData(v []*PaidResourceDTO) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.Data = v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetErrCode(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetErrMessage(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetHttpStatusCode(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetMaxResults(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.MaxResults = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetNextToken(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.NextToken = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetPageIndex(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.PageIndex = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetPageSize(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.PageSize = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetRequestId(v string) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.RequestId = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetSkip(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.Skip = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetSuccess(v bool) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.Success = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) SetTotalCount(v int32) *AliyunConsoleOpenApiQueryPaidResourceResponseBody {
	s.TotalCount = &v
	return s
}

func (s *AliyunConsoleOpenApiQueryPaidResourceResponseBody) Validate() error {
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

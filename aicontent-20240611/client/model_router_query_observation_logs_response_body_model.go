// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationLogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*RequestLogDTO) *ModelRouterQueryObservationLogsResponseBody
	GetData() []*RequestLogDTO
	SetErrCode(v string) *ModelRouterQueryObservationLogsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryObservationLogsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryObservationLogsResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryObservationLogsResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryObservationLogsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryObservationLogsResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryObservationLogsResponseBody struct {
	// example:
	//
	// []
	Data []*RequestLogDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ModelRouterQueryObservationLogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationLogsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetData() []*RequestLogDTO {
	return s.Data
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryObservationLogsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetData(v []*RequestLogDTO) *ModelRouterQueryObservationLogsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetErrCode(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetErrMessage(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetMaxResults(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetNextToken(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetPageIndex(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetPageSize(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetRequestId(v string) *ModelRouterQueryObservationLogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetSkip(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetSuccess(v bool) *ModelRouterQueryObservationLogsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) SetTotalCount(v int32) *ModelRouterQueryObservationLogsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryObservationLogsResponseBody) Validate() error {
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

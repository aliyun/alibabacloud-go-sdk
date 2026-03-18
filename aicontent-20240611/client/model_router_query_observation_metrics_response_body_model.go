// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryObservationMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*ModelMetricsDTO) *ModelRouterQueryObservationMetricsResponseBody
	GetData() []*ModelMetricsDTO
	SetErrCode(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetNextToken() *string
	SetPageIndex(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetPageIndex() *int32
	SetPageSize(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ModelRouterQueryObservationMetricsResponseBody
	GetRequestId() *string
	SetSkip(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetSkip() *int32
	SetSuccess(v bool) *ModelRouterQueryObservationMetricsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ModelRouterQueryObservationMetricsResponseBody
	GetTotalCount() *int32
}

type ModelRouterQueryObservationMetricsResponseBody struct {
	// example:
	//
	// []
	Data []*ModelMetricsDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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

func (s ModelRouterQueryObservationMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryObservationMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetData() []*ModelMetricsDTO {
	return s.Data
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetPageIndex() *int32 {
	return s.PageIndex
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetSkip() *int32 {
	return s.Skip
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryObservationMetricsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetData(v []*ModelMetricsDTO) *ModelRouterQueryObservationMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetErrCode(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetErrMessage(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetMaxResults(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetNextToken(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetPageIndex(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.PageIndex = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetPageSize(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetRequestId(v string) *ModelRouterQueryObservationMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetSkip(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.Skip = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetSuccess(v bool) *ModelRouterQueryObservationMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) SetTotalCount(v int32) *ModelRouterQueryObservationMetricsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ModelRouterQueryObservationMetricsResponseBody) Validate() error {
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

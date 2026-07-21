// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostTrendMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CostQueryTrendDTO) *ModelRouterQueryCostTrendMetricsResponseBody
	GetData() *CostQueryTrendDTO
	SetErrCode(v string) *ModelRouterQueryCostTrendMetricsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryCostTrendMetricsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryCostTrendMetricsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryCostTrendMetricsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryCostTrendMetricsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryCostTrendMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryCostTrendMetricsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryCostTrendMetricsResponseBody struct {
	// The data object containing the cost trend metrics.
	//
	// example:
	//
	// []
	Data *CostQueryTrendDTO `json:"data,omitempty" xml:"data,omitempty"`
	// The error code returned when the request fails.
	//
	// example:
	//
	// UNKNOWN_ERROR
	ErrCode *string `json:"errCode,omitempty" xml:"errCode,omitempty"`
	// The error message returned when the request fails.
	//
	// example:
	//
	// 未知错误
	ErrMessage *string `json:"errMessage,omitempty" xml:"errMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// The maximum number of results returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"maxResults,omitempty" xml:"maxResults,omitempty"`
	// A token to retrieve the next page of results. If no token is returned, all results have been retrieved.
	//
	// example:
	//
	// xxxx-xxx-xxxxx
	NextToken *string `json:"nextToken,omitempty" xml:"nextToken,omitempty"`
	// The unique request ID. If you encounter a problem, provide this ID to technical support for troubleshooting.
	//
	// example:
	//
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// Indicates whether the request was successful. A value of `true` indicates success, and `false` indicates failure.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryCostTrendMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostTrendMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetData() *CostQueryTrendDTO {
	return s.Data
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetData(v *CostQueryTrendDTO) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetErrCode(v string) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetErrMessage(v string) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetMaxResults(v int32) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetNextToken(v string) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetRequestId(v string) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) SetSuccess(v bool) *ModelRouterQueryCostTrendMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryCostTrendMetricsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

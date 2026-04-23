// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelRouterQueryCostOverviewMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*MetricValueDTO) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetData() []*MetricValueDTO
	SetErrCode(v string) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v int32) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ModelRouterQueryCostOverviewMetricsResponseBody
	GetSuccess() *bool
}

type ModelRouterQueryCostOverviewMetricsResponseBody struct {
	// example:
	//
	// []
	Data []*MetricValueDTO `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
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
	// xxxx-xxxx-xxxx-xxxxxxxx
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelRouterQueryCostOverviewMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelRouterQueryCostOverviewMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetData() []*MetricValueDTO {
	return s.Data
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetData(v []*MetricValueDTO) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.Data = v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetErrCode(v string) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetErrMessage(v string) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetHttpStatusCode(v int32) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetMaxResults(v int32) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetNextToken(v string) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetRequestId(v string) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) SetSuccess(v bool) *ModelRouterQueryCostOverviewMetricsResponseBody {
	s.Success = &v
	return s
}

func (s *ModelRouterQueryCostOverviewMetricsResponseBody) Validate() error {
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

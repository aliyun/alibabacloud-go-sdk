// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRunMetricsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetrics(v []*RunMetric) *ListRunMetricsResponseBody
	GetMetrics() []*RunMetric
	SetNextPageToken(v int64) *ListRunMetricsResponseBody
	GetNextPageToken() *int64
	SetRequestId(v string) *ListRunMetricsResponseBody
	GetRequestId() *string
}

type ListRunMetricsResponseBody struct {
	// The metrics.
	Metrics []*RunMetric `json:"Metrics,omitempty" xml:"Metrics,omitempty" type:"Repeated"`
	// The pagination token that is used to retrieve the next page. You do not need to specify this parameter for the first request. You must specify the pagination token in the result of the previous query. If the pagination token is 0, no next page exists. You can obtain the pagination token that is used to retrieve the next page in the value of the **NextPageToken*	- field.
	//
	// example:
	//
	// 0
	NextPageToken *int64 `json:"NextPageToken,omitempty" xml:"NextPageToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// ADF6D849-*****-7E7030F0CE53
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListRunMetricsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListRunMetricsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRunMetricsResponseBody) GetMetrics() []*RunMetric {
	return s.Metrics
}

func (s *ListRunMetricsResponseBody) GetNextPageToken() *int64 {
	return s.NextPageToken
}

func (s *ListRunMetricsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListRunMetricsResponseBody) SetMetrics(v []*RunMetric) *ListRunMetricsResponseBody {
	s.Metrics = v
	return s
}

func (s *ListRunMetricsResponseBody) SetNextPageToken(v int64) *ListRunMetricsResponseBody {
	s.NextPageToken = &v
	return s
}

func (s *ListRunMetricsResponseBody) SetRequestId(v string) *ListRunMetricsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRunMetricsResponseBody) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryHistoryAvgMetricListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryHistoryAvgMetricListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryHistoryAvgMetricListResponse
	GetStatusCode() *int32
	SetBody(v *QueryHistoryAvgMetricListResponseBody) *QueryHistoryAvgMetricListResponse
	GetBody() *QueryHistoryAvgMetricListResponseBody
}

type QueryHistoryAvgMetricListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryHistoryAvgMetricListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryHistoryAvgMetricListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryHistoryAvgMetricListResponse) GoString() string {
	return s.String()
}

func (s *QueryHistoryAvgMetricListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryHistoryAvgMetricListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryHistoryAvgMetricListResponse) GetBody() *QueryHistoryAvgMetricListResponseBody {
	return s.Body
}

func (s *QueryHistoryAvgMetricListResponse) SetHeaders(v map[string]*string) *QueryHistoryAvgMetricListResponse {
	s.Headers = v
	return s
}

func (s *QueryHistoryAvgMetricListResponse) SetStatusCode(v int32) *QueryHistoryAvgMetricListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryHistoryAvgMetricListResponse) SetBody(v *QueryHistoryAvgMetricListResponseBody) *QueryHistoryAvgMetricListResponse {
	s.Body = v
	return s
}

func (s *QueryHistoryAvgMetricListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

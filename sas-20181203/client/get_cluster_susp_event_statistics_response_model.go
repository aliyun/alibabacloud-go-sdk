// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetClusterSuspEventStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetClusterSuspEventStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetClusterSuspEventStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetClusterSuspEventStatisticsResponseBody) *GetClusterSuspEventStatisticsResponse
	GetBody() *GetClusterSuspEventStatisticsResponseBody
}

type GetClusterSuspEventStatisticsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetClusterSuspEventStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetClusterSuspEventStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetClusterSuspEventStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetClusterSuspEventStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetClusterSuspEventStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetClusterSuspEventStatisticsResponse) GetBody() *GetClusterSuspEventStatisticsResponseBody {
	return s.Body
}

func (s *GetClusterSuspEventStatisticsResponse) SetHeaders(v map[string]*string) *GetClusterSuspEventStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetClusterSuspEventStatisticsResponse) SetStatusCode(v int32) *GetClusterSuspEventStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetClusterSuspEventStatisticsResponse) SetBody(v *GetClusterSuspEventStatisticsResponseBody) *GetClusterSuspEventStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetClusterSuspEventStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

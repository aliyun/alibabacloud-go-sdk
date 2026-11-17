// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSourceUsageStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSourceUsageStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSourceUsageStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListSourceUsageStatisticsResponseBody) *ListSourceUsageStatisticsResponse
	GetBody() *ListSourceUsageStatisticsResponseBody
}

type ListSourceUsageStatisticsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSourceUsageStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSourceUsageStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSourceUsageStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListSourceUsageStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSourceUsageStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSourceUsageStatisticsResponse) GetBody() *ListSourceUsageStatisticsResponseBody {
	return s.Body
}

func (s *ListSourceUsageStatisticsResponse) SetHeaders(v map[string]*string) *ListSourceUsageStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListSourceUsageStatisticsResponse) SetStatusCode(v int32) *ListSourceUsageStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSourceUsageStatisticsResponse) SetBody(v *ListSourceUsageStatisticsResponseBody) *ListSourceUsageStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListSourceUsageStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrafficStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrafficStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrafficStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrafficStatisticsResponseBody) *ListTrafficStatisticsResponse
	GetBody() *ListTrafficStatisticsResponseBody
}

type ListTrafficStatisticsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrafficStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrafficStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrafficStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListTrafficStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrafficStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrafficStatisticsResponse) GetBody() *ListTrafficStatisticsResponseBody {
	return s.Body
}

func (s *ListTrafficStatisticsResponse) SetHeaders(v map[string]*string) *ListTrafficStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListTrafficStatisticsResponse) SetStatusCode(v int32) *ListTrafficStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrafficStatisticsResponse) SetBody(v *ListTrafficStatisticsResponseBody) *ListTrafficStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListTrafficStatisticsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

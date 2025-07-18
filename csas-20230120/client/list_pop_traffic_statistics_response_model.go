// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPopTrafficStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPopTrafficStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPopTrafficStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *ListPopTrafficStatisticsResponseBody) *ListPopTrafficStatisticsResponse
	GetBody() *ListPopTrafficStatisticsResponseBody
}

type ListPopTrafficStatisticsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPopTrafficStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPopTrafficStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPopTrafficStatisticsResponse) GoString() string {
	return s.String()
}

func (s *ListPopTrafficStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPopTrafficStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPopTrafficStatisticsResponse) GetBody() *ListPopTrafficStatisticsResponseBody {
	return s.Body
}

func (s *ListPopTrafficStatisticsResponse) SetHeaders(v map[string]*string) *ListPopTrafficStatisticsResponse {
	s.Headers = v
	return s
}

func (s *ListPopTrafficStatisticsResponse) SetStatusCode(v int32) *ListPopTrafficStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPopTrafficStatisticsResponse) SetBody(v *ListPopTrafficStatisticsResponseBody) *ListPopTrafficStatisticsResponse {
	s.Body = v
	return s
}

func (s *ListPopTrafficStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSuspiciousStatisticsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSuspiciousStatisticsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSuspiciousStatisticsResponse
	GetStatusCode() *int32
	SetBody(v *GetSuspiciousStatisticsResponseBody) *GetSuspiciousStatisticsResponse
	GetBody() *GetSuspiciousStatisticsResponseBody
}

type GetSuspiciousStatisticsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSuspiciousStatisticsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSuspiciousStatisticsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSuspiciousStatisticsResponse) GoString() string {
	return s.String()
}

func (s *GetSuspiciousStatisticsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSuspiciousStatisticsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSuspiciousStatisticsResponse) GetBody() *GetSuspiciousStatisticsResponseBody {
	return s.Body
}

func (s *GetSuspiciousStatisticsResponse) SetHeaders(v map[string]*string) *GetSuspiciousStatisticsResponse {
	s.Headers = v
	return s
}

func (s *GetSuspiciousStatisticsResponse) SetStatusCode(v int32) *GetSuspiciousStatisticsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSuspiciousStatisticsResponse) SetBody(v *GetSuspiciousStatisticsResponseBody) *GetSuspiciousStatisticsResponse {
	s.Body = v
	return s
}

func (s *GetSuspiciousStatisticsResponse) Validate() error {
	return dara.Validate(s)
}

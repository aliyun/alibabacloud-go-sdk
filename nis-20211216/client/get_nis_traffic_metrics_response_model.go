// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisTrafficMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNisTrafficMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNisTrafficMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetNisTrafficMetricsResponseBody) *GetNisTrafficMetricsResponse
	GetBody() *GetNisTrafficMetricsResponseBody
}

type GetNisTrafficMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNisTrafficMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNisTrafficMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNisTrafficMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetNisTrafficMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNisTrafficMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNisTrafficMetricsResponse) GetBody() *GetNisTrafficMetricsResponseBody {
	return s.Body
}

func (s *GetNisTrafficMetricsResponse) SetHeaders(v map[string]*string) *GetNisTrafficMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetNisTrafficMetricsResponse) SetStatusCode(v int32) *GetNisTrafficMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNisTrafficMetricsResponse) SetBody(v *GetNisTrafficMetricsResponseBody) *GetNisTrafficMetricsResponse {
	s.Body = v
	return s
}

func (s *GetNisTrafficMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

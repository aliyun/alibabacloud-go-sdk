// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNisNetworkMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNisNetworkMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNisNetworkMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetNisNetworkMetricsResponseBody) *GetNisNetworkMetricsResponse
	GetBody() *GetNisNetworkMetricsResponseBody
}

type GetNisNetworkMetricsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNisNetworkMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNisNetworkMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNisNetworkMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetNisNetworkMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNisNetworkMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNisNetworkMetricsResponse) GetBody() *GetNisNetworkMetricsResponseBody {
	return s.Body
}

func (s *GetNisNetworkMetricsResponse) SetHeaders(v map[string]*string) *GetNisNetworkMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetNisNetworkMetricsResponse) SetStatusCode(v int32) *GetNisNetworkMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNisNetworkMetricsResponse) SetBody(v *GetNisNetworkMetricsResponseBody) *GetNisNetworkMetricsResponse {
	s.Body = v
	return s
}

func (s *GetNisNetworkMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

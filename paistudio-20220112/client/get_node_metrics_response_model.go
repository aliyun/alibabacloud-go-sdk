// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetNodeMetricsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetNodeMetricsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetNodeMetricsResponse
	GetStatusCode() *int32
	SetBody(v *GetNodeMetricsResponseBody) *GetNodeMetricsResponse
	GetBody() *GetNodeMetricsResponseBody
}

type GetNodeMetricsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetNodeMetricsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetNodeMetricsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetNodeMetricsResponse) GoString() string {
	return s.String()
}

func (s *GetNodeMetricsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetNodeMetricsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetNodeMetricsResponse) GetBody() *GetNodeMetricsResponseBody {
	return s.Body
}

func (s *GetNodeMetricsResponse) SetHeaders(v map[string]*string) *GetNodeMetricsResponse {
	s.Headers = v
	return s
}

func (s *GetNodeMetricsResponse) SetStatusCode(v int32) *GetNodeMetricsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetNodeMetricsResponse) SetBody(v *GetNodeMetricsResponseBody) *GetNodeMetricsResponse {
	s.Body = v
	return s
}

func (s *GetNodeMetricsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

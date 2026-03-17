// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHealthCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHealthCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHealthCheckResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHealthCheckResponseBody) *DeleteHealthCheckResponse
	GetBody() *DeleteHealthCheckResponseBody
}

type DeleteHealthCheckResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHealthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *DeleteHealthCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHealthCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHealthCheckResponse) GetBody() *DeleteHealthCheckResponseBody {
	return s.Body
}

func (s *DeleteHealthCheckResponse) SetHeaders(v map[string]*string) *DeleteHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *DeleteHealthCheckResponse) SetStatusCode(v int32) *DeleteHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHealthCheckResponse) SetBody(v *DeleteHealthCheckResponseBody) *DeleteHealthCheckResponse {
	s.Body = v
	return s
}

func (s *DeleteHealthCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

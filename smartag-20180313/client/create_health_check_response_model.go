// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHealthCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHealthCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHealthCheckResponse
	GetStatusCode() *int32
	SetBody(v *CreateHealthCheckResponseBody) *CreateHealthCheckResponse
	GetBody() *CreateHealthCheckResponseBody
}

type CreateHealthCheckResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHealthCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHealthCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHealthCheckResponse) GoString() string {
	return s.String()
}

func (s *CreateHealthCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHealthCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHealthCheckResponse) GetBody() *CreateHealthCheckResponseBody {
	return s.Body
}

func (s *CreateHealthCheckResponse) SetHeaders(v map[string]*string) *CreateHealthCheckResponse {
	s.Headers = v
	return s
}

func (s *CreateHealthCheckResponse) SetStatusCode(v int32) *CreateHealthCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHealthCheckResponse) SetBody(v *CreateHealthCheckResponseBody) *CreateHealthCheckResponse {
	s.Body = v
	return s
}

func (s *CreateHealthCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

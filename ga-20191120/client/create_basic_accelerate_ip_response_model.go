// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBasicAccelerateIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBasicAccelerateIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBasicAccelerateIpResponse
	GetStatusCode() *int32
	SetBody(v *CreateBasicAccelerateIpResponseBody) *CreateBasicAccelerateIpResponse
	GetBody() *CreateBasicAccelerateIpResponseBody
}

type CreateBasicAccelerateIpResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBasicAccelerateIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBasicAccelerateIpResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBasicAccelerateIpResponse) GoString() string {
	return s.String()
}

func (s *CreateBasicAccelerateIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBasicAccelerateIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBasicAccelerateIpResponse) GetBody() *CreateBasicAccelerateIpResponseBody {
	return s.Body
}

func (s *CreateBasicAccelerateIpResponse) SetHeaders(v map[string]*string) *CreateBasicAccelerateIpResponse {
	s.Headers = v
	return s
}

func (s *CreateBasicAccelerateIpResponse) SetStatusCode(v int32) *CreateBasicAccelerateIpResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBasicAccelerateIpResponse) SetBody(v *CreateBasicAccelerateIpResponseBody) *CreateBasicAccelerateIpResponse {
	s.Body = v
	return s
}

func (s *CreateBasicAccelerateIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeregisterResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeregisterResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeregisterResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *DeregisterResourceTypeResponseBody) *DeregisterResourceTypeResponse
	GetBody() *DeregisterResourceTypeResponseBody
}

type DeregisterResourceTypeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeregisterResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeregisterResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeregisterResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *DeregisterResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeregisterResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeregisterResourceTypeResponse) GetBody() *DeregisterResourceTypeResponseBody {
	return s.Body
}

func (s *DeregisterResourceTypeResponse) SetHeaders(v map[string]*string) *DeregisterResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *DeregisterResourceTypeResponse) SetStatusCode(v int32) *DeregisterResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeregisterResourceTypeResponse) SetBody(v *DeregisterResourceTypeResponseBody) *DeregisterResourceTypeResponse {
	s.Body = v
	return s
}

func (s *DeregisterResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

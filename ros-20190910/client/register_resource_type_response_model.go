// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterResourceTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterResourceTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterResourceTypeResponse
	GetStatusCode() *int32
	SetBody(v *RegisterResourceTypeResponseBody) *RegisterResourceTypeResponse
	GetBody() *RegisterResourceTypeResponseBody
}

type RegisterResourceTypeResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterResourceTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterResourceTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterResourceTypeResponse) GoString() string {
	return s.String()
}

func (s *RegisterResourceTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterResourceTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterResourceTypeResponse) GetBody() *RegisterResourceTypeResponseBody {
	return s.Body
}

func (s *RegisterResourceTypeResponse) SetHeaders(v map[string]*string) *RegisterResourceTypeResponse {
	s.Headers = v
	return s
}

func (s *RegisterResourceTypeResponse) SetStatusCode(v int32) *RegisterResourceTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterResourceTypeResponse) SetBody(v *RegisterResourceTypeResponseBody) *RegisterResourceTypeResponse {
	s.Body = v
	return s
}

func (s *RegisterResourceTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

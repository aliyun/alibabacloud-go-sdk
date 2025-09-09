// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRegisterServiceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RegisterServiceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RegisterServiceResponse
	GetStatusCode() *int32
	SetBody(v *RegisterServiceResponseBody) *RegisterServiceResponse
	GetBody() *RegisterServiceResponseBody
}

type RegisterServiceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RegisterServiceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RegisterServiceResponse) String() string {
	return dara.Prettify(s)
}

func (s RegisterServiceResponse) GoString() string {
	return s.String()
}

func (s *RegisterServiceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RegisterServiceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RegisterServiceResponse) GetBody() *RegisterServiceResponseBody {
	return s.Body
}

func (s *RegisterServiceResponse) SetHeaders(v map[string]*string) *RegisterServiceResponse {
	s.Headers = v
	return s
}

func (s *RegisterServiceResponse) SetStatusCode(v int32) *RegisterServiceResponse {
	s.StatusCode = &v
	return s
}

func (s *RegisterServiceResponse) SetBody(v *RegisterServiceResponseBody) *RegisterServiceResponse {
	s.Body = v
	return s
}

func (s *RegisterServiceResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomHostnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomHostnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomHostnameResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomHostnameResponseBody) *UpdateCustomHostnameResponse
	GetBody() *UpdateCustomHostnameResponseBody
}

type UpdateCustomHostnameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomHostnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomHostnameResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomHostnameResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomHostnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomHostnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomHostnameResponse) GetBody() *UpdateCustomHostnameResponseBody {
	return s.Body
}

func (s *UpdateCustomHostnameResponse) SetHeaders(v map[string]*string) *UpdateCustomHostnameResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomHostnameResponse) SetStatusCode(v int32) *UpdateCustomHostnameResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomHostnameResponse) SetBody(v *UpdateCustomHostnameResponseBody) *UpdateCustomHostnameResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomHostnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

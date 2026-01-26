// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCustomHostnameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCustomHostnameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCustomHostnameResponse
	GetStatusCode() *int32
	SetBody(v *CreateCustomHostnameResponseBody) *CreateCustomHostnameResponse
	GetBody() *CreateCustomHostnameResponseBody
}

type CreateCustomHostnameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomHostnameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomHostnameResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCustomHostnameResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomHostnameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCustomHostnameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCustomHostnameResponse) GetBody() *CreateCustomHostnameResponseBody {
	return s.Body
}

func (s *CreateCustomHostnameResponse) SetHeaders(v map[string]*string) *CreateCustomHostnameResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomHostnameResponse) SetStatusCode(v int32) *CreateCustomHostnameResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomHostnameResponse) SetBody(v *CreateCustomHostnameResponseBody) *CreateCustomHostnameResponse {
	s.Body = v
	return s
}

func (s *CreateCustomHostnameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

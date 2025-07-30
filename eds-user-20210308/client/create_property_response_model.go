// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePropertyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePropertyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePropertyResponse
	GetStatusCode() *int32
	SetBody(v *CreatePropertyResponseBody) *CreatePropertyResponse
	GetBody() *CreatePropertyResponseBody
}

type CreatePropertyResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePropertyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePropertyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePropertyResponse) GoString() string {
	return s.String()
}

func (s *CreatePropertyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePropertyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePropertyResponse) GetBody() *CreatePropertyResponseBody {
	return s.Body
}

func (s *CreatePropertyResponse) SetHeaders(v map[string]*string) *CreatePropertyResponse {
	s.Headers = v
	return s
}

func (s *CreatePropertyResponse) SetStatusCode(v int32) *CreatePropertyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePropertyResponse) SetBody(v *CreatePropertyResponseBody) *CreatePropertyResponse {
	s.Body = v
	return s
}

func (s *CreatePropertyResponse) Validate() error {
	return dara.Validate(s)
}

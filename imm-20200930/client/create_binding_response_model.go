// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBindingResponse
	GetStatusCode() *int32
	SetBody(v *CreateBindingResponseBody) *CreateBindingResponse
	GetBody() *CreateBindingResponseBody
}

type CreateBindingResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBindingResponse) GoString() string {
	return s.String()
}

func (s *CreateBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBindingResponse) GetBody() *CreateBindingResponseBody {
	return s.Body
}

func (s *CreateBindingResponse) SetHeaders(v map[string]*string) *CreateBindingResponse {
	s.Headers = v
	return s
}

func (s *CreateBindingResponse) SetStatusCode(v int32) *CreateBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBindingResponse) SetBody(v *CreateBindingResponseBody) *CreateBindingResponse {
	s.Body = v
	return s
}

func (s *CreateBindingResponse) Validate() error {
	return dara.Validate(s)
}

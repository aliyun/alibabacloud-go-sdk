// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBindingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBindingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBindingResponse
	GetStatusCode() *int32
	SetBody(v *GetBindingResponseBody) *GetBindingResponse
	GetBody() *GetBindingResponseBody
}

type GetBindingResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBindingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBindingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBindingResponse) GoString() string {
	return s.String()
}

func (s *GetBindingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBindingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBindingResponse) GetBody() *GetBindingResponseBody {
	return s.Body
}

func (s *GetBindingResponse) SetHeaders(v map[string]*string) *GetBindingResponse {
	s.Headers = v
	return s
}

func (s *GetBindingResponse) SetStatusCode(v int32) *GetBindingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBindingResponse) SetBody(v *GetBindingResponseBody) *GetBindingResponse {
	s.Body = v
	return s
}

func (s *GetBindingResponse) Validate() error {
	return dara.Validate(s)
}

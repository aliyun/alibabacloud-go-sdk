// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomTextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomTextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomTextResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomTextResponseBody) *UpdateCustomTextResponse
	GetBody() *UpdateCustomTextResponseBody
}

type UpdateCustomTextResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomTextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomTextResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomTextResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomTextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomTextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomTextResponse) GetBody() *UpdateCustomTextResponseBody {
	return s.Body
}

func (s *UpdateCustomTextResponse) SetHeaders(v map[string]*string) *UpdateCustomTextResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomTextResponse) SetStatusCode(v int32) *UpdateCustomTextResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomTextResponse) SetBody(v *UpdateCustomTextResponseBody) *UpdateCustomTextResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomTextResponse) Validate() error {
	return dara.Validate(s)
}

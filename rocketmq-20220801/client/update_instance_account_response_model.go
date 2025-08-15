// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceAccountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceAccountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceAccountResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceAccountResponseBody) *UpdateInstanceAccountResponse
	GetBody() *UpdateInstanceAccountResponseBody
}

type UpdateInstanceAccountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceAccountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceAccountResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceAccountResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceAccountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceAccountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceAccountResponse) GetBody() *UpdateInstanceAccountResponseBody {
	return s.Body
}

func (s *UpdateInstanceAccountResponse) SetHeaders(v map[string]*string) *UpdateInstanceAccountResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceAccountResponse) SetStatusCode(v int32) *UpdateInstanceAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceAccountResponse) SetBody(v *UpdateInstanceAccountResponseBody) *UpdateInstanceAccountResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceAccountResponse) Validate() error {
	return dara.Validate(s)
}

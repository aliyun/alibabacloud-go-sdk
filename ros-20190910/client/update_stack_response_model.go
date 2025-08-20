// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateStackResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateStackResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateStackResponse
	GetStatusCode() *int32
	SetBody(v *UpdateStackResponseBody) *UpdateStackResponse
	GetBody() *UpdateStackResponseBody
}

type UpdateStackResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStackResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStackResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateStackResponse) GoString() string {
	return s.String()
}

func (s *UpdateStackResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateStackResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateStackResponse) GetBody() *UpdateStackResponseBody {
	return s.Body
}

func (s *UpdateStackResponse) SetHeaders(v map[string]*string) *UpdateStackResponse {
	s.Headers = v
	return s
}

func (s *UpdateStackResponse) SetStatusCode(v int32) *UpdateStackResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStackResponse) SetBody(v *UpdateStackResponseBody) *UpdateStackResponse {
	s.Body = v
	return s
}

func (s *UpdateStackResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAliasResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAliasResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAliasResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse
	GetBody() *UpdateAliasResponseBody
}

type UpdateAliasResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAliasResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAliasResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAliasResponse) GoString() string {
	return s.String()
}

func (s *UpdateAliasResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAliasResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAliasResponse) GetBody() *UpdateAliasResponseBody {
	return s.Body
}

func (s *UpdateAliasResponse) SetHeaders(v map[string]*string) *UpdateAliasResponse {
	s.Headers = v
	return s
}

func (s *UpdateAliasResponse) SetStatusCode(v int32) *UpdateAliasResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAliasResponse) SetBody(v *UpdateAliasResponseBody) *UpdateAliasResponse {
	s.Body = v
	return s
}

func (s *UpdateAliasResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAclResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAclResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAclResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAclResponseBody) *UpdateAclResponse
	GetBody() *UpdateAclResponseBody
}

type UpdateAclResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAclResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAclResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAclResponse) GoString() string {
	return s.String()
}

func (s *UpdateAclResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAclResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAclResponse) GetBody() *UpdateAclResponseBody {
	return s.Body
}

func (s *UpdateAclResponse) SetHeaders(v map[string]*string) *UpdateAclResponse {
	s.Headers = v
	return s
}

func (s *UpdateAclResponse) SetStatusCode(v int32) *UpdateAclResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAclResponse) SetBody(v *UpdateAclResponseBody) *UpdateAclResponse {
	s.Body = v
	return s
}

func (s *UpdateAclResponse) Validate() error {
	return dara.Validate(s)
}

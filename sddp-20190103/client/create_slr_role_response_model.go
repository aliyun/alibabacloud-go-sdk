// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSlrRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSlrRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSlrRoleResponse
	GetStatusCode() *int32
	SetBody(v *CreateSlrRoleResponseBody) *CreateSlrRoleResponse
	GetBody() *CreateSlrRoleResponseBody
}

type CreateSlrRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSlrRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSlrRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSlrRoleResponse) GoString() string {
	return s.String()
}

func (s *CreateSlrRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSlrRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSlrRoleResponse) GetBody() *CreateSlrRoleResponseBody {
	return s.Body
}

func (s *CreateSlrRoleResponse) SetHeaders(v map[string]*string) *CreateSlrRoleResponse {
	s.Headers = v
	return s
}

func (s *CreateSlrRoleResponse) SetStatusCode(v int32) *CreateSlrRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSlrRoleResponse) SetBody(v *CreateSlrRoleResponseBody) *CreateSlrRoleResponse {
	s.Body = v
	return s
}

func (s *CreateSlrRoleResponse) Validate() error {
	return dara.Validate(s)
}

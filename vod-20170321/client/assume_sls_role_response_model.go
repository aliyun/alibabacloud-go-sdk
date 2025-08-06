// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeSlsRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeSlsRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeSlsRoleResponse
	GetStatusCode() *int32
	SetBody(v *AssumeSlsRoleResponseBody) *AssumeSlsRoleResponse
	GetBody() *AssumeSlsRoleResponseBody
}

type AssumeSlsRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeSlsRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeSlsRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AssumeSlsRoleResponse) GoString() string {
	return s.String()
}

func (s *AssumeSlsRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeSlsRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeSlsRoleResponse) GetBody() *AssumeSlsRoleResponseBody {
	return s.Body
}

func (s *AssumeSlsRoleResponse) SetHeaders(v map[string]*string) *AssumeSlsRoleResponse {
	s.Headers = v
	return s
}

func (s *AssumeSlsRoleResponse) SetStatusCode(v int32) *AssumeSlsRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeSlsRoleResponse) SetBody(v *AssumeSlsRoleResponseBody) *AssumeSlsRoleResponse {
	s.Body = v
	return s
}

func (s *AssumeSlsRoleResponse) Validate() error {
	return dara.Validate(s)
}

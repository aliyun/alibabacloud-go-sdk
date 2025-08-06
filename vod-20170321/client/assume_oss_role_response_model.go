// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeOssRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeOssRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeOssRoleResponse
	GetStatusCode() *int32
	SetBody(v *AssumeOssRoleResponseBody) *AssumeOssRoleResponse
	GetBody() *AssumeOssRoleResponseBody
}

type AssumeOssRoleResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeOssRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeOssRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AssumeOssRoleResponse) GoString() string {
	return s.String()
}

func (s *AssumeOssRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeOssRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeOssRoleResponse) GetBody() *AssumeOssRoleResponseBody {
	return s.Body
}

func (s *AssumeOssRoleResponse) SetHeaders(v map[string]*string) *AssumeOssRoleResponse {
	s.Headers = v
	return s
}

func (s *AssumeOssRoleResponse) SetStatusCode(v int32) *AssumeOssRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeOssRoleResponse) SetBody(v *AssumeOssRoleResponseBody) *AssumeOssRoleResponse {
	s.Body = v
	return s
}

func (s *AssumeOssRoleResponse) Validate() error {
	return dara.Validate(s)
}

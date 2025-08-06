// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeExperienceRoleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AssumeExperienceRoleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AssumeExperienceRoleResponse
	GetStatusCode() *int32
	SetBody(v *AssumeExperienceRoleResponseBody) *AssumeExperienceRoleResponse
	GetBody() *AssumeExperienceRoleResponseBody
}

type AssumeExperienceRoleResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AssumeExperienceRoleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AssumeExperienceRoleResponse) String() string {
	return dara.Prettify(s)
}

func (s AssumeExperienceRoleResponse) GoString() string {
	return s.String()
}

func (s *AssumeExperienceRoleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AssumeExperienceRoleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AssumeExperienceRoleResponse) GetBody() *AssumeExperienceRoleResponseBody {
	return s.Body
}

func (s *AssumeExperienceRoleResponse) SetHeaders(v map[string]*string) *AssumeExperienceRoleResponse {
	s.Headers = v
	return s
}

func (s *AssumeExperienceRoleResponse) SetStatusCode(v int32) *AssumeExperienceRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *AssumeExperienceRoleResponse) SetBody(v *AssumeExperienceRoleResponseBody) *AssumeExperienceRoleResponse {
	s.Body = v
	return s
}

func (s *AssumeExperienceRoleResponse) Validate() error {
	return dara.Validate(s)
}

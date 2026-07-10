// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLangfuseOrgMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLangfuseOrgMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLangfuseOrgMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateLangfuseOrgMemberResponseBody) *CreateLangfuseOrgMemberResponse
	GetBody() *CreateLangfuseOrgMemberResponseBody
}

type CreateLangfuseOrgMemberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLangfuseOrgMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLangfuseOrgMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLangfuseOrgMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateLangfuseOrgMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLangfuseOrgMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLangfuseOrgMemberResponse) GetBody() *CreateLangfuseOrgMemberResponseBody {
	return s.Body
}

func (s *CreateLangfuseOrgMemberResponse) SetHeaders(v map[string]*string) *CreateLangfuseOrgMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateLangfuseOrgMemberResponse) SetStatusCode(v int32) *CreateLangfuseOrgMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLangfuseOrgMemberResponse) SetBody(v *CreateLangfuseOrgMemberResponseBody) *CreateLangfuseOrgMemberResponse {
	s.Body = v
	return s
}

func (s *CreateLangfuseOrgMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

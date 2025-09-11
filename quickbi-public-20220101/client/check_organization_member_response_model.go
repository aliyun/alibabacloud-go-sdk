// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckOrganizationMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckOrganizationMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckOrganizationMemberResponse
	GetStatusCode() *int32
	SetBody(v *CheckOrganizationMemberResponseBody) *CheckOrganizationMemberResponse
	GetBody() *CheckOrganizationMemberResponseBody
}

type CheckOrganizationMemberResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckOrganizationMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckOrganizationMemberResponse) GoString() string {
	return s.String()
}

func (s *CheckOrganizationMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckOrganizationMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckOrganizationMemberResponse) GetBody() *CheckOrganizationMemberResponseBody {
	return s.Body
}

func (s *CheckOrganizationMemberResponse) SetHeaders(v map[string]*string) *CheckOrganizationMemberResponse {
	s.Headers = v
	return s
}

func (s *CheckOrganizationMemberResponse) SetStatusCode(v int32) *CheckOrganizationMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckOrganizationMemberResponse) SetBody(v *CheckOrganizationMemberResponseBody) *CheckOrganizationMemberResponse {
	s.Body = v
	return s
}

func (s *CheckOrganizationMemberResponse) Validate() error {
	return dara.Validate(s)
}

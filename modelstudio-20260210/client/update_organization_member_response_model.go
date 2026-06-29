// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateOrganizationMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateOrganizationMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateOrganizationMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateOrganizationMemberResponseBody) *UpdateOrganizationMemberResponse
	GetBody() *UpdateOrganizationMemberResponseBody
}

type UpdateOrganizationMemberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateOrganizationMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateOrganizationMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateOrganizationMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateOrganizationMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateOrganizationMemberResponse) GetBody() *UpdateOrganizationMemberResponseBody {
	return s.Body
}

func (s *UpdateOrganizationMemberResponse) SetHeaders(v map[string]*string) *UpdateOrganizationMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateOrganizationMemberResponse) SetStatusCode(v int32) *UpdateOrganizationMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateOrganizationMemberResponse) SetBody(v *UpdateOrganizationMemberResponseBody) *UpdateOrganizationMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateOrganizationMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

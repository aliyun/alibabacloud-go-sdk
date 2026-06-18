// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddOrganizationMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddOrganizationMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddOrganizationMemberResponse
	GetStatusCode() *int32
	SetBody(v *AddOrganizationMemberResponseBody) *AddOrganizationMemberResponse
	GetBody() *AddOrganizationMemberResponseBody
}

type AddOrganizationMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddOrganizationMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s AddOrganizationMemberResponse) GoString() string {
	return s.String()
}

func (s *AddOrganizationMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddOrganizationMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddOrganizationMemberResponse) GetBody() *AddOrganizationMemberResponseBody {
	return s.Body
}

func (s *AddOrganizationMemberResponse) SetHeaders(v map[string]*string) *AddOrganizationMemberResponse {
	s.Headers = v
	return s
}

func (s *AddOrganizationMemberResponse) SetStatusCode(v int32) *AddOrganizationMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *AddOrganizationMemberResponse) SetBody(v *AddOrganizationMemberResponseBody) *AddOrganizationMemberResponse {
	s.Body = v
	return s
}

func (s *AddOrganizationMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

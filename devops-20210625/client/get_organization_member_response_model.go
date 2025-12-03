// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOrganizationMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOrganizationMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOrganizationMemberResponse
	GetStatusCode() *int32
	SetBody(v *GetOrganizationMemberResponseBody) *GetOrganizationMemberResponse
	GetBody() *GetOrganizationMemberResponseBody
}

type GetOrganizationMemberResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOrganizationMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOrganizationMemberResponse) GoString() string {
	return s.String()
}

func (s *GetOrganizationMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOrganizationMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOrganizationMemberResponse) GetBody() *GetOrganizationMemberResponseBody {
	return s.Body
}

func (s *GetOrganizationMemberResponse) SetHeaders(v map[string]*string) *GetOrganizationMemberResponse {
	s.Headers = v
	return s
}

func (s *GetOrganizationMemberResponse) SetStatusCode(v int32) *GetOrganizationMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOrganizationMemberResponse) SetBody(v *GetOrganizationMemberResponseBody) *GetOrganizationMemberResponse {
	s.Body = v
	return s
}

func (s *GetOrganizationMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

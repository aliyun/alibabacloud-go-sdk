// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveOrganizationMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveOrganizationMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveOrganizationMemberResponse
	GetStatusCode() *int32
	SetBody(v *RemoveOrganizationMemberResponseBody) *RemoveOrganizationMemberResponse
	GetBody() *RemoveOrganizationMemberResponseBody
}

type RemoveOrganizationMemberResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveOrganizationMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveOrganizationMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveOrganizationMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveOrganizationMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveOrganizationMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveOrganizationMemberResponse) GetBody() *RemoveOrganizationMemberResponseBody {
	return s.Body
}

func (s *RemoveOrganizationMemberResponse) SetHeaders(v map[string]*string) *RemoveOrganizationMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveOrganizationMemberResponse) SetStatusCode(v int32) *RemoveOrganizationMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveOrganizationMemberResponse) SetBody(v *RemoveOrganizationMemberResponseBody) *RemoveOrganizationMemberResponse {
	s.Body = v
	return s
}

func (s *RemoveOrganizationMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

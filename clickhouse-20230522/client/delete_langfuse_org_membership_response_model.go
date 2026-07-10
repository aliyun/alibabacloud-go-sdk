// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLangfuseOrgMembershipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLangfuseOrgMembershipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLangfuseOrgMembershipResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLangfuseOrgMembershipResponseBody) *DeleteLangfuseOrgMembershipResponse
	GetBody() *DeleteLangfuseOrgMembershipResponseBody
}

type DeleteLangfuseOrgMembershipResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLangfuseOrgMembershipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLangfuseOrgMembershipResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLangfuseOrgMembershipResponse) GoString() string {
	return s.String()
}

func (s *DeleteLangfuseOrgMembershipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLangfuseOrgMembershipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLangfuseOrgMembershipResponse) GetBody() *DeleteLangfuseOrgMembershipResponseBody {
	return s.Body
}

func (s *DeleteLangfuseOrgMembershipResponse) SetHeaders(v map[string]*string) *DeleteLangfuseOrgMembershipResponse {
	s.Headers = v
	return s
}

func (s *DeleteLangfuseOrgMembershipResponse) SetStatusCode(v int32) *DeleteLangfuseOrgMembershipResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLangfuseOrgMembershipResponse) SetBody(v *DeleteLangfuseOrgMembershipResponseBody) *DeleteLangfuseOrgMembershipResponse {
	s.Body = v
	return s
}

func (s *DeleteLangfuseOrgMembershipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

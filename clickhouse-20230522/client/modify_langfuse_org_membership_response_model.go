// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseOrgMembershipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLangfuseOrgMembershipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLangfuseOrgMembershipResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLangfuseOrgMembershipResponseBody) *ModifyLangfuseOrgMembershipResponse
	GetBody() *ModifyLangfuseOrgMembershipResponseBody
}

type ModifyLangfuseOrgMembershipResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLangfuseOrgMembershipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLangfuseOrgMembershipResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseOrgMembershipResponse) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseOrgMembershipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLangfuseOrgMembershipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLangfuseOrgMembershipResponse) GetBody() *ModifyLangfuseOrgMembershipResponseBody {
	return s.Body
}

func (s *ModifyLangfuseOrgMembershipResponse) SetHeaders(v map[string]*string) *ModifyLangfuseOrgMembershipResponse {
	s.Headers = v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponse) SetStatusCode(v int32) *ModifyLangfuseOrgMembershipResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponse) SetBody(v *ModifyLangfuseOrgMembershipResponseBody) *ModifyLangfuseOrgMembershipResponse {
	s.Body = v
	return s
}

func (s *ModifyLangfuseOrgMembershipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLangfuseProjectMembershipResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyLangfuseProjectMembershipResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyLangfuseProjectMembershipResponse
	GetStatusCode() *int32
	SetBody(v *ModifyLangfuseProjectMembershipResponseBody) *ModifyLangfuseProjectMembershipResponse
	GetBody() *ModifyLangfuseProjectMembershipResponseBody
}

type ModifyLangfuseProjectMembershipResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyLangfuseProjectMembershipResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyLangfuseProjectMembershipResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyLangfuseProjectMembershipResponse) GoString() string {
	return s.String()
}

func (s *ModifyLangfuseProjectMembershipResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyLangfuseProjectMembershipResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyLangfuseProjectMembershipResponse) GetBody() *ModifyLangfuseProjectMembershipResponseBody {
	return s.Body
}

func (s *ModifyLangfuseProjectMembershipResponse) SetHeaders(v map[string]*string) *ModifyLangfuseProjectMembershipResponse {
	s.Headers = v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponse) SetStatusCode(v int32) *ModifyLangfuseProjectMembershipResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponse) SetBody(v *ModifyLangfuseProjectMembershipResponseBody) *ModifyLangfuseProjectMembershipResponse {
	s.Body = v
	return s
}

func (s *ModifyLangfuseProjectMembershipResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

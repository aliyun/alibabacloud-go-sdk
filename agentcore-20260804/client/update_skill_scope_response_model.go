// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillScopeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillScopeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillScopeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillScopeResponseBody) *UpdateSkillScopeResponse
	GetBody() *UpdateSkillScopeResponseBody
}

type UpdateSkillScopeResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillScopeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillScopeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillScopeResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillScopeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillScopeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillScopeResponse) GetBody() *UpdateSkillScopeResponseBody {
	return s.Body
}

func (s *UpdateSkillScopeResponse) SetHeaders(v map[string]*string) *UpdateSkillScopeResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillScopeResponse) SetStatusCode(v int32) *UpdateSkillScopeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillScopeResponse) SetBody(v *UpdateSkillScopeResponseBody) *UpdateSkillScopeResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillScopeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

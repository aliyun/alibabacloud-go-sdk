// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillGroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillGroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillGroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillGroupConfigResponseBody) *UpdateSkillGroupConfigResponse
	GetBody() *UpdateSkillGroupConfigResponseBody
}

type UpdateSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillGroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillGroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillGroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillGroupConfigResponse) GetBody() *UpdateSkillGroupConfigResponseBody {
	return s.Body
}

func (s *UpdateSkillGroupConfigResponse) SetHeaders(v map[string]*string) *UpdateSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillGroupConfigResponse) SetStatusCode(v int32) *UpdateSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillGroupConfigResponse) SetBody(v *UpdateSkillGroupConfigResponseBody) *UpdateSkillGroupConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillGroupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

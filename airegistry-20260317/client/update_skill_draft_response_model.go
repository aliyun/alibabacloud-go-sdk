// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillDraftResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillDraftResponseBody) *UpdateSkillDraftResponse
	GetBody() *UpdateSkillDraftResponseBody
}

type UpdateSkillDraftResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillDraftResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillDraftResponse) GetBody() *UpdateSkillDraftResponseBody {
	return s.Body
}

func (s *UpdateSkillDraftResponse) SetHeaders(v map[string]*string) *UpdateSkillDraftResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillDraftResponse) SetStatusCode(v int32) *UpdateSkillDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillDraftResponse) SetBody(v *UpdateSkillDraftResponseBody) *UpdateSkillDraftResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

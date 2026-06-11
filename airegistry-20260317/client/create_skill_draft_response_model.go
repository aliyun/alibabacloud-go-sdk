// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillDraftResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillDraftResponseBody) *CreateSkillDraftResponse
	GetBody() *CreateSkillDraftResponseBody
}

type CreateSkillDraftResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillDraftResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillDraftResponse) GetBody() *CreateSkillDraftResponseBody {
	return s.Body
}

func (s *CreateSkillDraftResponse) SetHeaders(v map[string]*string) *CreateSkillDraftResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillDraftResponse) SetStatusCode(v int32) *CreateSkillDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillDraftResponse) SetBody(v *CreateSkillDraftResponseBody) *CreateSkillDraftResponse {
	s.Body = v
	return s
}

func (s *CreateSkillDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

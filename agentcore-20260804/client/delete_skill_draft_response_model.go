// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillDraftResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSkillDraftResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSkillDraftResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSkillDraftResponseBody) *DeleteSkillDraftResponse
	GetBody() *DeleteSkillDraftResponseBody
}

type DeleteSkillDraftResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSkillDraftResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSkillDraftResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillDraftResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillDraftResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSkillDraftResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSkillDraftResponse) GetBody() *DeleteSkillDraftResponseBody {
	return s.Body
}

func (s *DeleteSkillDraftResponse) SetHeaders(v map[string]*string) *DeleteSkillDraftResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillDraftResponse) SetStatusCode(v int32) *DeleteSkillDraftResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillDraftResponse) SetBody(v *DeleteSkillDraftResponseBody) *DeleteSkillDraftResponse {
	s.Body = v
	return s
}

func (s *DeleteSkillDraftResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

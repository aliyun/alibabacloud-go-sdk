// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSkillResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSkillResponseBody) *DeleteSkillResponse
	GetBody() *DeleteSkillResponseBody
}

type DeleteSkillResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSkillResponse) GetBody() *DeleteSkillResponseBody {
	return s.Body
}

func (s *DeleteSkillResponse) SetHeaders(v map[string]*string) *DeleteSkillResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillResponse) SetStatusCode(v int32) *DeleteSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillResponse) SetBody(v *DeleteSkillResponseBody) *DeleteSkillResponse {
	s.Body = v
	return s
}

func (s *DeleteSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

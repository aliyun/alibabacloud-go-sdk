// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillGroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSkillGroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSkillGroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSkillGroupConfigResponseBody) *DeleteSkillGroupConfigResponse
	GetBody() *DeleteSkillGroupConfigResponseBody
}

type DeleteSkillGroupConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSkillGroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillGroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSkillGroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSkillGroupConfigResponse) GetBody() *DeleteSkillGroupConfigResponseBody {
	return s.Body
}

func (s *DeleteSkillGroupConfigResponse) SetHeaders(v map[string]*string) *DeleteSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillGroupConfigResponse) SetStatusCode(v int32) *DeleteSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillGroupConfigResponse) SetBody(v *DeleteSkillGroupConfigResponseBody) *DeleteSkillGroupConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteSkillGroupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSkillSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSkillSpaceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSkillSpaceResponseBody) *DeleteSkillSpaceResponse
	GetBody() *DeleteSkillSpaceResponseBody
}

type DeleteSkillSpaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSkillSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSkillSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillSpaceResponse) GoString() string {
	return s.String()
}

func (s *DeleteSkillSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSkillSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSkillSpaceResponse) GetBody() *DeleteSkillSpaceResponseBody {
	return s.Body
}

func (s *DeleteSkillSpaceResponse) SetHeaders(v map[string]*string) *DeleteSkillSpaceResponse {
	s.Headers = v
	return s
}

func (s *DeleteSkillSpaceResponse) SetStatusCode(v int32) *DeleteSkillSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSkillSpaceResponse) SetBody(v *DeleteSkillSpaceResponseBody) *DeleteSkillSpaceResponse {
	s.Body = v
	return s
}

func (s *DeleteSkillSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

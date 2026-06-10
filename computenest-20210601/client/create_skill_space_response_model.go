// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillSpaceResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillSpaceResponseBody) *CreateSkillSpaceResponse
	GetBody() *CreateSkillSpaceResponseBody
}

type CreateSkillSpaceResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillSpaceResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillSpaceResponse) GetBody() *CreateSkillSpaceResponseBody {
	return s.Body
}

func (s *CreateSkillSpaceResponse) SetHeaders(v map[string]*string) *CreateSkillSpaceResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillSpaceResponse) SetStatusCode(v int32) *CreateSkillSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillSpaceResponse) SetBody(v *CreateSkillSpaceResponseBody) *CreateSkillSpaceResponse {
	s.Body = v
	return s
}

func (s *CreateSkillSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

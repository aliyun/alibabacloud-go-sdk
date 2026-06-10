// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillSpaceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillSpaceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillSpaceResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillSpaceResponseBody) *GetSkillSpaceResponse
	GetBody() *GetSkillSpaceResponseBody
}

type GetSkillSpaceResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillSpaceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillSpaceResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillSpaceResponse) GoString() string {
	return s.String()
}

func (s *GetSkillSpaceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillSpaceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillSpaceResponse) GetBody() *GetSkillSpaceResponseBody {
	return s.Body
}

func (s *GetSkillSpaceResponse) SetHeaders(v map[string]*string) *GetSkillSpaceResponse {
	s.Headers = v
	return s
}

func (s *GetSkillSpaceResponse) SetStatusCode(v int32) *GetSkillSpaceResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillSpaceResponse) SetBody(v *GetSkillSpaceResponseBody) *GetSkillSpaceResponse {
	s.Body = v
	return s
}

func (s *GetSkillSpaceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

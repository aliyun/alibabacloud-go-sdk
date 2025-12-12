// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillGroupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillGroupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillGroupConfigResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillGroupConfigResponseBody) *ListSkillGroupConfigResponse
	GetBody() *ListSkillGroupConfigResponseBody
}

type ListSkillGroupConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillGroupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillGroupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillGroupConfigResponse) GoString() string {
	return s.String()
}

func (s *ListSkillGroupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillGroupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillGroupConfigResponse) GetBody() *ListSkillGroupConfigResponseBody {
	return s.Body
}

func (s *ListSkillGroupConfigResponse) SetHeaders(v map[string]*string) *ListSkillGroupConfigResponse {
	s.Headers = v
	return s
}

func (s *ListSkillGroupConfigResponse) SetStatusCode(v int32) *ListSkillGroupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillGroupConfigResponse) SetBody(v *ListSkillGroupConfigResponseBody) *ListSkillGroupConfigResponse {
	s.Body = v
	return s
}

func (s *ListSkillGroupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

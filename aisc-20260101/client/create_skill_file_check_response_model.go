// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillFileCheckResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillFileCheckResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillFileCheckResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillFileCheckResponseBody) *CreateSkillFileCheckResponse
	GetBody() *CreateSkillFileCheckResponseBody
}

type CreateSkillFileCheckResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillFileCheckResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillFileCheckResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileCheckResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillFileCheckResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillFileCheckResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillFileCheckResponse) GetBody() *CreateSkillFileCheckResponseBody {
	return s.Body
}

func (s *CreateSkillFileCheckResponse) SetHeaders(v map[string]*string) *CreateSkillFileCheckResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillFileCheckResponse) SetStatusCode(v int32) *CreateSkillFileCheckResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillFileCheckResponse) SetBody(v *CreateSkillFileCheckResponseBody) *CreateSkillFileCheckResponse {
	s.Body = v
	return s
}

func (s *CreateSkillFileCheckResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

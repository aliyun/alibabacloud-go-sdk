// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillResponseBody) *GetSkillResponse
	GetBody() *GetSkillResponseBody
}

type GetSkillResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillResponse) GoString() string {
	return s.String()
}

func (s *GetSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillResponse) GetBody() *GetSkillResponseBody {
	return s.Body
}

func (s *GetSkillResponse) SetHeaders(v map[string]*string) *GetSkillResponse {
	s.Headers = v
	return s
}

func (s *GetSkillResponse) SetStatusCode(v int32) *GetSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillResponse) SetBody(v *GetSkillResponseBody) *GetSkillResponse {
	s.Body = v
	return s
}

func (s *GetSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

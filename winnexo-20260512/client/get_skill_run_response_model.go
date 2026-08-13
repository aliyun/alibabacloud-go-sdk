// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillRunResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillRunResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillRunResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillRunResponseBody) *GetSkillRunResponse
	GetBody() *GetSkillRunResponseBody
}

type GetSkillRunResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillRunResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillRunResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillRunResponse) GoString() string {
	return s.String()
}

func (s *GetSkillRunResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillRunResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillRunResponse) GetBody() *GetSkillRunResponseBody {
	return s.Body
}

func (s *GetSkillRunResponse) SetHeaders(v map[string]*string) *GetSkillRunResponse {
	s.Headers = v
	return s
}

func (s *GetSkillRunResponse) SetStatusCode(v int32) *GetSkillRunResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillRunResponse) SetBody(v *GetSkillRunResponseBody) *GetSkillRunResponse {
	s.Body = v
	return s
}

func (s *GetSkillRunResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

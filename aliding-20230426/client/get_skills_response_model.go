// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSkillsResponse
	GetStatusCode() *int32
	SetBody(v *GetSkillsResponseBody) *GetSkillsResponse
	GetBody() *GetSkillsResponseBody
}

type GetSkillsResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSkillsResponse) GoString() string {
	return s.String()
}

func (s *GetSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSkillsResponse) GetBody() *GetSkillsResponseBody {
	return s.Body
}

func (s *GetSkillsResponse) SetHeaders(v map[string]*string) *GetSkillsResponse {
	s.Headers = v
	return s
}

func (s *GetSkillsResponse) SetStatusCode(v int32) *GetSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSkillsResponse) SetBody(v *GetSkillsResponseBody) *GetSkillsResponse {
	s.Body = v
	return s
}

func (s *GetSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

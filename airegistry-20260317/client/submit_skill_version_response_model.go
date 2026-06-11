// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSkillVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitSkillVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitSkillVersionResponse
	GetStatusCode() *int32
	SetBody(v *SubmitSkillVersionResponseBody) *SubmitSkillVersionResponse
	GetBody() *SubmitSkillVersionResponseBody
}

type SubmitSkillVersionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitSkillVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitSkillVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionResponse) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitSkillVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitSkillVersionResponse) GetBody() *SubmitSkillVersionResponseBody {
	return s.Body
}

func (s *SubmitSkillVersionResponse) SetHeaders(v map[string]*string) *SubmitSkillVersionResponse {
	s.Headers = v
	return s
}

func (s *SubmitSkillVersionResponse) SetStatusCode(v int32) *SubmitSkillVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitSkillVersionResponse) SetBody(v *SubmitSkillVersionResponseBody) *SubmitSkillVersionResponse {
	s.Body = v
	return s
}

func (s *SubmitSkillVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

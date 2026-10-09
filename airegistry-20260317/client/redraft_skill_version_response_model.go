// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedraftSkillVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RedraftSkillVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RedraftSkillVersionResponse
	GetStatusCode() *int32
	SetBody(v *RedraftSkillVersionResponseBody) *RedraftSkillVersionResponse
	GetBody() *RedraftSkillVersionResponseBody
}

type RedraftSkillVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RedraftSkillVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RedraftSkillVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s RedraftSkillVersionResponse) GoString() string {
	return s.String()
}

func (s *RedraftSkillVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RedraftSkillVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RedraftSkillVersionResponse) GetBody() *RedraftSkillVersionResponseBody {
	return s.Body
}

func (s *RedraftSkillVersionResponse) SetHeaders(v map[string]*string) *RedraftSkillVersionResponse {
	s.Headers = v
	return s
}

func (s *RedraftSkillVersionResponse) SetStatusCode(v int32) *RedraftSkillVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *RedraftSkillVersionResponse) SetBody(v *RedraftSkillVersionResponseBody) *RedraftSkillVersionResponse {
	s.Body = v
	return s
}

func (s *RedraftSkillVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

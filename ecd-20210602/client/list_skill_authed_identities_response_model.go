// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillAuthedIdentitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillAuthedIdentitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillAuthedIdentitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillAuthedIdentitiesResponseBody) *ListSkillAuthedIdentitiesResponse
	GetBody() *ListSkillAuthedIdentitiesResponseBody
}

type ListSkillAuthedIdentitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillAuthedIdentitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillAuthedIdentitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillAuthedIdentitiesResponse) GoString() string {
	return s.String()
}

func (s *ListSkillAuthedIdentitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillAuthedIdentitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillAuthedIdentitiesResponse) GetBody() *ListSkillAuthedIdentitiesResponseBody {
	return s.Body
}

func (s *ListSkillAuthedIdentitiesResponse) SetHeaders(v map[string]*string) *ListSkillAuthedIdentitiesResponse {
	s.Headers = v
	return s
}

func (s *ListSkillAuthedIdentitiesResponse) SetStatusCode(v int32) *ListSkillAuthedIdentitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillAuthedIdentitiesResponse) SetBody(v *ListSkillAuthedIdentitiesResponseBody) *ListSkillAuthedIdentitiesResponse {
	s.Body = v
	return s
}

func (s *ListSkillAuthedIdentitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

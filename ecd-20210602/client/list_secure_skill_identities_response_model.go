// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSecureSkillIdentitiesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSecureSkillIdentitiesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSecureSkillIdentitiesResponse
	GetStatusCode() *int32
	SetBody(v *ListSecureSkillIdentitiesResponseBody) *ListSecureSkillIdentitiesResponse
	GetBody() *ListSecureSkillIdentitiesResponseBody
}

type ListSecureSkillIdentitiesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSecureSkillIdentitiesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSecureSkillIdentitiesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSecureSkillIdentitiesResponse) GoString() string {
	return s.String()
}

func (s *ListSecureSkillIdentitiesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSecureSkillIdentitiesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSecureSkillIdentitiesResponse) GetBody() *ListSecureSkillIdentitiesResponseBody {
	return s.Body
}

func (s *ListSecureSkillIdentitiesResponse) SetHeaders(v map[string]*string) *ListSecureSkillIdentitiesResponse {
	s.Headers = v
	return s
}

func (s *ListSecureSkillIdentitiesResponse) SetStatusCode(v int32) *ListSecureSkillIdentitiesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSecureSkillIdentitiesResponse) SetBody(v *ListSecureSkillIdentitiesResponseBody) *ListSecureSkillIdentitiesResponse {
	s.Body = v
	return s
}

func (s *ListSecureSkillIdentitiesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

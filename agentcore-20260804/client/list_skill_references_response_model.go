// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillReferencesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillReferencesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillReferencesResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillReferencesResponseBody) *ListSkillReferencesResponse
	GetBody() *ListSkillReferencesResponseBody
}

type ListSkillReferencesResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillReferencesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillReferencesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillReferencesResponse) GoString() string {
	return s.String()
}

func (s *ListSkillReferencesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillReferencesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillReferencesResponse) GetBody() *ListSkillReferencesResponseBody {
	return s.Body
}

func (s *ListSkillReferencesResponse) SetHeaders(v map[string]*string) *ListSkillReferencesResponse {
	s.Headers = v
	return s
}

func (s *ListSkillReferencesResponse) SetStatusCode(v int32) *ListSkillReferencesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillReferencesResponse) SetBody(v *ListSkillReferencesResponseBody) *ListSkillReferencesResponse {
	s.Body = v
	return s
}

func (s *ListSkillReferencesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

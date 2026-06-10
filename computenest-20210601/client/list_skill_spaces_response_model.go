// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillSpacesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillSpacesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillSpacesResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillSpacesResponseBody) *ListSkillSpacesResponse
	GetBody() *ListSkillSpacesResponseBody
}

type ListSkillSpacesResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillSpacesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillSpacesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillSpacesResponse) GoString() string {
	return s.String()
}

func (s *ListSkillSpacesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillSpacesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillSpacesResponse) GetBody() *ListSkillSpacesResponseBody {
	return s.Body
}

func (s *ListSkillSpacesResponse) SetHeaders(v map[string]*string) *ListSkillSpacesResponse {
	s.Headers = v
	return s
}

func (s *ListSkillSpacesResponse) SetStatusCode(v int32) *ListSkillSpacesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillSpacesResponse) SetBody(v *ListSkillSpacesResponseBody) *ListSkillSpacesResponse {
	s.Body = v
	return s
}

func (s *ListSkillSpacesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

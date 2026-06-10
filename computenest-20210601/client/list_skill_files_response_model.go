// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSkillFilesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSkillFilesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSkillFilesResponse
	GetStatusCode() *int32
	SetBody(v *ListSkillFilesResponseBody) *ListSkillFilesResponse
	GetBody() *ListSkillFilesResponseBody
}

type ListSkillFilesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSkillFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSkillFilesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSkillFilesResponse) GoString() string {
	return s.String()
}

func (s *ListSkillFilesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSkillFilesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSkillFilesResponse) GetBody() *ListSkillFilesResponseBody {
	return s.Body
}

func (s *ListSkillFilesResponse) SetHeaders(v map[string]*string) *ListSkillFilesResponse {
	s.Headers = v
	return s
}

func (s *ListSkillFilesResponse) SetStatusCode(v int32) *ListSkillFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSkillFilesResponse) SetBody(v *ListSkillFilesResponseBody) *ListSkillFilesResponse {
	s.Body = v
	return s
}

func (s *ListSkillFilesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

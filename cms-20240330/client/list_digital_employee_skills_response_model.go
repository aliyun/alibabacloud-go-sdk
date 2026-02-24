// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeeSkillsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDigitalEmployeeSkillsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDigitalEmployeeSkillsResponse
	GetStatusCode() *int32
	SetBody(v *ListDigitalEmployeeSkillsResponseBody) *ListDigitalEmployeeSkillsResponse
	GetBody() *ListDigitalEmployeeSkillsResponseBody
}

type ListDigitalEmployeeSkillsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDigitalEmployeeSkillsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDigitalEmployeeSkillsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillsResponse) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDigitalEmployeeSkillsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDigitalEmployeeSkillsResponse) GetBody() *ListDigitalEmployeeSkillsResponseBody {
	return s.Body
}

func (s *ListDigitalEmployeeSkillsResponse) SetHeaders(v map[string]*string) *ListDigitalEmployeeSkillsResponse {
	s.Headers = v
	return s
}

func (s *ListDigitalEmployeeSkillsResponse) SetStatusCode(v int32) *ListDigitalEmployeeSkillsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDigitalEmployeeSkillsResponse) SetBody(v *ListDigitalEmployeeSkillsResponseBody) *ListDigitalEmployeeSkillsResponse {
	s.Body = v
	return s
}

func (s *ListDigitalEmployeeSkillsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

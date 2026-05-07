// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDigitalEmployeeSkillVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDigitalEmployeeSkillVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDigitalEmployeeSkillVersionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDigitalEmployeeSkillVersionsResponseBody) *ListDigitalEmployeeSkillVersionsResponse
	GetBody() *ListDigitalEmployeeSkillVersionsResponseBody
}

type ListDigitalEmployeeSkillVersionsResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDigitalEmployeeSkillVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDigitalEmployeeSkillVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDigitalEmployeeSkillVersionsResponse) GoString() string {
	return s.String()
}

func (s *ListDigitalEmployeeSkillVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDigitalEmployeeSkillVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDigitalEmployeeSkillVersionsResponse) GetBody() *ListDigitalEmployeeSkillVersionsResponseBody {
	return s.Body
}

func (s *ListDigitalEmployeeSkillVersionsResponse) SetHeaders(v map[string]*string) *ListDigitalEmployeeSkillVersionsResponse {
	s.Headers = v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponse) SetStatusCode(v int32) *ListDigitalEmployeeSkillVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponse) SetBody(v *ListDigitalEmployeeSkillVersionsResponseBody) *ListDigitalEmployeeSkillVersionsResponse {
	s.Body = v
	return s
}

func (s *ListDigitalEmployeeSkillVersionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

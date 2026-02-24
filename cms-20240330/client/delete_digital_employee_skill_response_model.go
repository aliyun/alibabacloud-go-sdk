// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDigitalEmployeeSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDigitalEmployeeSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDigitalEmployeeSkillResponse
	GetStatusCode() *int32
	SetBody(v *DeleteDigitalEmployeeSkillResponseBody) *DeleteDigitalEmployeeSkillResponse
	GetBody() *DeleteDigitalEmployeeSkillResponseBody
}

type DeleteDigitalEmployeeSkillResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteDigitalEmployeeSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteDigitalEmployeeSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDigitalEmployeeSkillResponse) GoString() string {
	return s.String()
}

func (s *DeleteDigitalEmployeeSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDigitalEmployeeSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDigitalEmployeeSkillResponse) GetBody() *DeleteDigitalEmployeeSkillResponseBody {
	return s.Body
}

func (s *DeleteDigitalEmployeeSkillResponse) SetHeaders(v map[string]*string) *DeleteDigitalEmployeeSkillResponse {
	s.Headers = v
	return s
}

func (s *DeleteDigitalEmployeeSkillResponse) SetStatusCode(v int32) *DeleteDigitalEmployeeSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDigitalEmployeeSkillResponse) SetBody(v *DeleteDigitalEmployeeSkillResponseBody) *DeleteDigitalEmployeeSkillResponse {
	s.Body = v
	return s
}

func (s *DeleteDigitalEmployeeSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

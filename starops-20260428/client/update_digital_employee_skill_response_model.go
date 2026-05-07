// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDigitalEmployeeSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDigitalEmployeeSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDigitalEmployeeSkillResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDigitalEmployeeSkillResponseBody) *UpdateDigitalEmployeeSkillResponse
	GetBody() *UpdateDigitalEmployeeSkillResponseBody
}

type UpdateDigitalEmployeeSkillResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDigitalEmployeeSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDigitalEmployeeSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDigitalEmployeeSkillResponse) GoString() string {
	return s.String()
}

func (s *UpdateDigitalEmployeeSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDigitalEmployeeSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDigitalEmployeeSkillResponse) GetBody() *UpdateDigitalEmployeeSkillResponseBody {
	return s.Body
}

func (s *UpdateDigitalEmployeeSkillResponse) SetHeaders(v map[string]*string) *UpdateDigitalEmployeeSkillResponse {
	s.Headers = v
	return s
}

func (s *UpdateDigitalEmployeeSkillResponse) SetStatusCode(v int32) *UpdateDigitalEmployeeSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDigitalEmployeeSkillResponse) SetBody(v *UpdateDigitalEmployeeSkillResponseBody) *UpdateDigitalEmployeeSkillResponse {
	s.Body = v
	return s
}

func (s *UpdateDigitalEmployeeSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

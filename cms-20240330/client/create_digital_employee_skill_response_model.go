// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDigitalEmployeeSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDigitalEmployeeSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDigitalEmployeeSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreateDigitalEmployeeSkillResponseBody) *CreateDigitalEmployeeSkillResponse
	GetBody() *CreateDigitalEmployeeSkillResponseBody
}

type CreateDigitalEmployeeSkillResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDigitalEmployeeSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDigitalEmployeeSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDigitalEmployeeSkillResponse) GoString() string {
	return s.String()
}

func (s *CreateDigitalEmployeeSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDigitalEmployeeSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDigitalEmployeeSkillResponse) GetBody() *CreateDigitalEmployeeSkillResponseBody {
	return s.Body
}

func (s *CreateDigitalEmployeeSkillResponse) SetHeaders(v map[string]*string) *CreateDigitalEmployeeSkillResponse {
	s.Headers = v
	return s
}

func (s *CreateDigitalEmployeeSkillResponse) SetStatusCode(v int32) *CreateDigitalEmployeeSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDigitalEmployeeSkillResponse) SetBody(v *CreateDigitalEmployeeSkillResponseBody) *CreateDigitalEmployeeSkillResponse {
	s.Body = v
	return s
}

func (s *CreateDigitalEmployeeSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

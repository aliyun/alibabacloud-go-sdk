// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDigitalEmployeeSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetDigitalEmployeeSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetDigitalEmployeeSkillResponse
	GetStatusCode() *int32
	SetBody(v *GetDigitalEmployeeSkillResponseBody) *GetDigitalEmployeeSkillResponse
	GetBody() *GetDigitalEmployeeSkillResponseBody
}

type GetDigitalEmployeeSkillResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDigitalEmployeeSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDigitalEmployeeSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s GetDigitalEmployeeSkillResponse) GoString() string {
	return s.String()
}

func (s *GetDigitalEmployeeSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetDigitalEmployeeSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetDigitalEmployeeSkillResponse) GetBody() *GetDigitalEmployeeSkillResponseBody {
	return s.Body
}

func (s *GetDigitalEmployeeSkillResponse) SetHeaders(v map[string]*string) *GetDigitalEmployeeSkillResponse {
	s.Headers = v
	return s
}

func (s *GetDigitalEmployeeSkillResponse) SetStatusCode(v int32) *GetDigitalEmployeeSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDigitalEmployeeSkillResponse) SetBody(v *GetDigitalEmployeeSkillResponseBody) *GetDigitalEmployeeSkillResponse {
	s.Body = v
	return s
}

func (s *GetDigitalEmployeeSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

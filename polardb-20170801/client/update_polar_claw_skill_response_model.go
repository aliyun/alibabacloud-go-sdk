// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePolarClawSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePolarClawSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePolarClawSkillResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePolarClawSkillResponseBody) *UpdatePolarClawSkillResponse
	GetBody() *UpdatePolarClawSkillResponseBody
}

type UpdatePolarClawSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePolarClawSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePolarClawSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePolarClawSkillResponse) GoString() string {
	return s.String()
}

func (s *UpdatePolarClawSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePolarClawSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePolarClawSkillResponse) GetBody() *UpdatePolarClawSkillResponseBody {
	return s.Body
}

func (s *UpdatePolarClawSkillResponse) SetHeaders(v map[string]*string) *UpdatePolarClawSkillResponse {
	s.Headers = v
	return s
}

func (s *UpdatePolarClawSkillResponse) SetStatusCode(v int32) *UpdatePolarClawSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePolarClawSkillResponse) SetBody(v *UpdatePolarClawSkillResponseBody) *UpdatePolarClawSkillResponse {
	s.Body = v
	return s
}

func (s *UpdatePolarClawSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

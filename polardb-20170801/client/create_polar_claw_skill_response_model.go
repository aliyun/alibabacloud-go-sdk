// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePolarClawSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreatePolarClawSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreatePolarClawSkillResponse
	GetStatusCode() *int32
	SetBody(v *CreatePolarClawSkillResponseBody) *CreatePolarClawSkillResponse
	GetBody() *CreatePolarClawSkillResponseBody
}

type CreatePolarClawSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreatePolarClawSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreatePolarClawSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CreatePolarClawSkillResponse) GoString() string {
	return s.String()
}

func (s *CreatePolarClawSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreatePolarClawSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreatePolarClawSkillResponse) GetBody() *CreatePolarClawSkillResponseBody {
	return s.Body
}

func (s *CreatePolarClawSkillResponse) SetHeaders(v map[string]*string) *CreatePolarClawSkillResponse {
	s.Headers = v
	return s
}

func (s *CreatePolarClawSkillResponse) SetStatusCode(v int32) *CreatePolarClawSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CreatePolarClawSkillResponse) SetBody(v *CreatePolarClawSkillResponseBody) *CreatePolarClawSkillResponse {
	s.Body = v
	return s
}

func (s *CreatePolarClawSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

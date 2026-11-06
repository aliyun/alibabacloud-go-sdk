// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPolarClawSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyPolarClawSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyPolarClawSkillResponse
	GetStatusCode() *int32
	SetBody(v *ModifyPolarClawSkillResponseBody) *ModifyPolarClawSkillResponse
	GetBody() *ModifyPolarClawSkillResponseBody
}

type ModifyPolarClawSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyPolarClawSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyPolarClawSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyPolarClawSkillResponse) GoString() string {
	return s.String()
}

func (s *ModifyPolarClawSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyPolarClawSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyPolarClawSkillResponse) GetBody() *ModifyPolarClawSkillResponseBody {
	return s.Body
}

func (s *ModifyPolarClawSkillResponse) SetHeaders(v map[string]*string) *ModifyPolarClawSkillResponse {
	s.Headers = v
	return s
}

func (s *ModifyPolarClawSkillResponse) SetStatusCode(v int32) *ModifyPolarClawSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyPolarClawSkillResponse) SetBody(v *ModifyPolarClawSkillResponseBody) *ModifyPolarClawSkillResponse {
	s.Body = v
	return s
}

func (s *ModifyPolarClawSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

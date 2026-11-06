// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePolarClawSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeletePolarClawSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeletePolarClawSkillResponse
	GetStatusCode() *int32
	SetBody(v *DeletePolarClawSkillResponseBody) *DeletePolarClawSkillResponse
	GetBody() *DeletePolarClawSkillResponseBody
}

type DeletePolarClawSkillResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeletePolarClawSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeletePolarClawSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s DeletePolarClawSkillResponse) GoString() string {
	return s.String()
}

func (s *DeletePolarClawSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeletePolarClawSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeletePolarClawSkillResponse) GetBody() *DeletePolarClawSkillResponseBody {
	return s.Body
}

func (s *DeletePolarClawSkillResponse) SetHeaders(v map[string]*string) *DeletePolarClawSkillResponse {
	s.Headers = v
	return s
}

func (s *DeletePolarClawSkillResponse) SetStatusCode(v int32) *DeletePolarClawSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *DeletePolarClawSkillResponse) SetBody(v *DeletePolarClawSkillResponseBody) *DeletePolarClawSkillResponse {
	s.Body = v
	return s
}

func (s *DeletePolarClawSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

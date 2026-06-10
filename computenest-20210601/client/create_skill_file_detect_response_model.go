// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillFileDetectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSkillFileDetectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSkillFileDetectResponse
	GetStatusCode() *int32
	SetBody(v *CreateSkillFileDetectResponseBody) *CreateSkillFileDetectResponse
	GetBody() *CreateSkillFileDetectResponseBody
}

type CreateSkillFileDetectResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSkillFileDetectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSkillFileDetectResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileDetectResponse) GoString() string {
	return s.String()
}

func (s *CreateSkillFileDetectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSkillFileDetectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSkillFileDetectResponse) GetBody() *CreateSkillFileDetectResponseBody {
	return s.Body
}

func (s *CreateSkillFileDetectResponse) SetHeaders(v map[string]*string) *CreateSkillFileDetectResponse {
	s.Headers = v
	return s
}

func (s *CreateSkillFileDetectResponse) SetStatusCode(v int32) *CreateSkillFileDetectResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSkillFileDetectResponse) SetBody(v *CreateSkillFileDetectResponseBody) *CreateSkillFileDetectResponse {
	s.Body = v
	return s
}

func (s *CreateSkillFileDetectResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

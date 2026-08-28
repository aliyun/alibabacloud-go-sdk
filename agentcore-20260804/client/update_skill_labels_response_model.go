// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSkillLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSkillLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSkillLabelsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSkillLabelsResponseBody) *UpdateSkillLabelsResponse
	GetBody() *UpdateSkillLabelsResponseBody
}

type UpdateSkillLabelsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSkillLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSkillLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSkillLabelsResponse) GoString() string {
	return s.String()
}

func (s *UpdateSkillLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSkillLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSkillLabelsResponse) GetBody() *UpdateSkillLabelsResponseBody {
	return s.Body
}

func (s *UpdateSkillLabelsResponse) SetHeaders(v map[string]*string) *UpdateSkillLabelsResponse {
	s.Headers = v
	return s
}

func (s *UpdateSkillLabelsResponse) SetStatusCode(v int32) *UpdateSkillLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSkillLabelsResponse) SetBody(v *UpdateSkillLabelsResponseBody) *UpdateSkillLabelsResponse {
	s.Body = v
	return s
}

func (s *UpdateSkillLabelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

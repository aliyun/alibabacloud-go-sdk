// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iForcePublishSkillVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ForcePublishSkillVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ForcePublishSkillVersionResponse
	GetStatusCode() *int32
	SetBody(v *ForcePublishSkillVersionResponseBody) *ForcePublishSkillVersionResponse
	GetBody() *ForcePublishSkillVersionResponseBody
}

type ForcePublishSkillVersionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ForcePublishSkillVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ForcePublishSkillVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s ForcePublishSkillVersionResponse) GoString() string {
	return s.String()
}

func (s *ForcePublishSkillVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ForcePublishSkillVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ForcePublishSkillVersionResponse) GetBody() *ForcePublishSkillVersionResponseBody {
	return s.Body
}

func (s *ForcePublishSkillVersionResponse) SetHeaders(v map[string]*string) *ForcePublishSkillVersionResponse {
	s.Headers = v
	return s
}

func (s *ForcePublishSkillVersionResponse) SetStatusCode(v int32) *ForcePublishSkillVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *ForcePublishSkillVersionResponse) SetBody(v *ForcePublishSkillVersionResponseBody) *ForcePublishSkillVersionResponse {
	s.Body = v
	return s
}

func (s *ForcePublishSkillVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

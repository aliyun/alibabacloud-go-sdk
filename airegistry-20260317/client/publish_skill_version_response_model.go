// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishSkillVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishSkillVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishSkillVersionResponse
	GetStatusCode() *int32
	SetBody(v *PublishSkillVersionResponseBody) *PublishSkillVersionResponse
	GetBody() *PublishSkillVersionResponseBody
}

type PublishSkillVersionResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishSkillVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishSkillVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishSkillVersionResponse) GoString() string {
	return s.String()
}

func (s *PublishSkillVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishSkillVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishSkillVersionResponse) GetBody() *PublishSkillVersionResponseBody {
	return s.Body
}

func (s *PublishSkillVersionResponse) SetHeaders(v map[string]*string) *PublishSkillVersionResponse {
	s.Headers = v
	return s
}

func (s *PublishSkillVersionResponse) SetStatusCode(v int32) *PublishSkillVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishSkillVersionResponse) SetBody(v *PublishSkillVersionResponseBody) *PublishSkillVersionResponse {
	s.Body = v
	return s
}

func (s *PublishSkillVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

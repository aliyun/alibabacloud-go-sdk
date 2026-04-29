// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudListSkillResponseBody) *CloudListSkillResponse
	GetBody() *CloudListSkillResponseBody
}

type CloudListSkillResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudListSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListSkillResponse) GetBody() *CloudListSkillResponseBody {
	return s.Body
}

func (s *CloudListSkillResponse) SetHeaders(v map[string]*string) *CloudListSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudListSkillResponse) SetStatusCode(v int32) *CloudListSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListSkillResponse) SetBody(v *CloudListSkillResponseBody) *CloudListSkillResponse {
	s.Body = v
	return s
}

func (s *CloudListSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteSkillResponseBody) *CloudDeleteSkillResponse
	GetBody() *CloudDeleteSkillResponseBody
}

type CloudDeleteSkillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteSkillResponse) GetBody() *CloudDeleteSkillResponseBody {
	return s.Body
}

func (s *CloudDeleteSkillResponse) SetHeaders(v map[string]*string) *CloudDeleteSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteSkillResponse) SetStatusCode(v int32) *CloudDeleteSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteSkillResponse) SetBody(v *CloudDeleteSkillResponseBody) *CloudDeleteSkillResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

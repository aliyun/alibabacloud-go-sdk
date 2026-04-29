// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateSkillResponseBody) *CloudCreateSkillResponse
	GetBody() *CloudCreateSkillResponseBody
}

type CloudCreateSkillResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateSkillResponse) GetBody() *CloudCreateSkillResponseBody {
	return s.Body
}

func (s *CloudCreateSkillResponse) SetHeaders(v map[string]*string) *CloudCreateSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateSkillResponse) SetStatusCode(v int32) *CloudCreateSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateSkillResponse) SetBody(v *CloudCreateSkillResponseBody) *CloudCreateSkillResponse {
	s.Body = v
	return s
}

func (s *CloudCreateSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

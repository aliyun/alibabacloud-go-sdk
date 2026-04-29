// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudCreateQueueSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudCreateQueueSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudCreateQueueSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudCreateQueueSkillResponseBody) *CloudCreateQueueSkillResponse
	GetBody() *CloudCreateQueueSkillResponseBody
}

type CloudCreateQueueSkillResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudCreateQueueSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudCreateQueueSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudCreateQueueSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudCreateQueueSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudCreateQueueSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudCreateQueueSkillResponse) GetBody() *CloudCreateQueueSkillResponseBody {
	return s.Body
}

func (s *CloudCreateQueueSkillResponse) SetHeaders(v map[string]*string) *CloudCreateQueueSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudCreateQueueSkillResponse) SetStatusCode(v int32) *CloudCreateQueueSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudCreateQueueSkillResponse) SetBody(v *CloudCreateQueueSkillResponseBody) *CloudCreateQueueSkillResponse {
	s.Body = v
	return s
}

func (s *CloudCreateQueueSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

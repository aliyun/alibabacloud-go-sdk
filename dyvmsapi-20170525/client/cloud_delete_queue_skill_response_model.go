// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudDeleteQueueSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudDeleteQueueSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudDeleteQueueSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudDeleteQueueSkillResponseBody) *CloudDeleteQueueSkillResponse
	GetBody() *CloudDeleteQueueSkillResponseBody
}

type CloudDeleteQueueSkillResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudDeleteQueueSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudDeleteQueueSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudDeleteQueueSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudDeleteQueueSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudDeleteQueueSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudDeleteQueueSkillResponse) GetBody() *CloudDeleteQueueSkillResponseBody {
	return s.Body
}

func (s *CloudDeleteQueueSkillResponse) SetHeaders(v map[string]*string) *CloudDeleteQueueSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudDeleteQueueSkillResponse) SetStatusCode(v int32) *CloudDeleteQueueSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudDeleteQueueSkillResponse) SetBody(v *CloudDeleteQueueSkillResponseBody) *CloudDeleteQueueSkillResponse {
	s.Body = v
	return s
}

func (s *CloudDeleteQueueSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

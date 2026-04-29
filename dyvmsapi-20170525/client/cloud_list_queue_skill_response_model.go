// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudListQueueSkillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloudListQueueSkillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloudListQueueSkillResponse
	GetStatusCode() *int32
	SetBody(v *CloudListQueueSkillResponseBody) *CloudListQueueSkillResponse
	GetBody() *CloudListQueueSkillResponseBody
}

type CloudListQueueSkillResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloudListQueueSkillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloudListQueueSkillResponse) String() string {
	return dara.Prettify(s)
}

func (s CloudListQueueSkillResponse) GoString() string {
	return s.String()
}

func (s *CloudListQueueSkillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloudListQueueSkillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloudListQueueSkillResponse) GetBody() *CloudListQueueSkillResponseBody {
	return s.Body
}

func (s *CloudListQueueSkillResponse) SetHeaders(v map[string]*string) *CloudListQueueSkillResponse {
	s.Headers = v
	return s
}

func (s *CloudListQueueSkillResponse) SetStatusCode(v int32) *CloudListQueueSkillResponse {
	s.StatusCode = &v
	return s
}

func (s *CloudListQueueSkillResponse) SetBody(v *CloudListQueueSkillResponseBody) *CloudListQueueSkillResponse {
	s.Body = v
	return s
}

func (s *CloudListQueueSkillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

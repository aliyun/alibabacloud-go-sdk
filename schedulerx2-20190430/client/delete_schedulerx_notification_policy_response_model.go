// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSchedulerxNotificationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSchedulerxNotificationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSchedulerxNotificationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSchedulerxNotificationPolicyResponseBody) *DeleteSchedulerxNotificationPolicyResponse
	GetBody() *DeleteSchedulerxNotificationPolicyResponseBody
}

type DeleteSchedulerxNotificationPolicyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSchedulerxNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSchedulerxNotificationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSchedulerxNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *DeleteSchedulerxNotificationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSchedulerxNotificationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSchedulerxNotificationPolicyResponse) GetBody() *DeleteSchedulerxNotificationPolicyResponseBody {
	return s.Body
}

func (s *DeleteSchedulerxNotificationPolicyResponse) SetHeaders(v map[string]*string) *DeleteSchedulerxNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponse) SetStatusCode(v int32) *DeleteSchedulerxNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponse) SetBody(v *DeleteSchedulerxNotificationPolicyResponseBody) *DeleteSchedulerxNotificationPolicyResponse {
	s.Body = v
	return s
}

func (s *DeleteSchedulerxNotificationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchedulerxNotificationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSchedulerxNotificationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSchedulerxNotificationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *CreateSchedulerxNotificationPolicyResponseBody) *CreateSchedulerxNotificationPolicyResponse
	GetBody() *CreateSchedulerxNotificationPolicyResponseBody
}

type CreateSchedulerxNotificationPolicyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSchedulerxNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSchedulerxNotificationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSchedulerxNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *CreateSchedulerxNotificationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSchedulerxNotificationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSchedulerxNotificationPolicyResponse) GetBody() *CreateSchedulerxNotificationPolicyResponseBody {
	return s.Body
}

func (s *CreateSchedulerxNotificationPolicyResponse) SetHeaders(v map[string]*string) *CreateSchedulerxNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponse) SetStatusCode(v int32) *CreateSchedulerxNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponse) SetBody(v *CreateSchedulerxNotificationPolicyResponseBody) *CreateSchedulerxNotificationPolicyResponse {
	s.Body = v
	return s
}

func (s *CreateSchedulerxNotificationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

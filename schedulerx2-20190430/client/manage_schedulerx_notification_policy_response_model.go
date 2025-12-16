// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iManageSchedulerxNotificationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ManageSchedulerxNotificationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ManageSchedulerxNotificationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ManageSchedulerxNotificationPolicyResponseBody) *ManageSchedulerxNotificationPolicyResponse
	GetBody() *ManageSchedulerxNotificationPolicyResponseBody
}

type ManageSchedulerxNotificationPolicyResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ManageSchedulerxNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ManageSchedulerxNotificationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ManageSchedulerxNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *ManageSchedulerxNotificationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ManageSchedulerxNotificationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ManageSchedulerxNotificationPolicyResponse) GetBody() *ManageSchedulerxNotificationPolicyResponseBody {
	return s.Body
}

func (s *ManageSchedulerxNotificationPolicyResponse) SetHeaders(v map[string]*string) *ManageSchedulerxNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponse) SetStatusCode(v int32) *ManageSchedulerxNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponse) SetBody(v *ManageSchedulerxNotificationPolicyResponseBody) *ManageSchedulerxNotificationPolicyResponse {
	s.Body = v
	return s
}

func (s *ManageSchedulerxNotificationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

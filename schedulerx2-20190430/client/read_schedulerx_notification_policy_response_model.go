// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReadSchedulerxNotificationPolicyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReadSchedulerxNotificationPolicyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReadSchedulerxNotificationPolicyResponse
	GetStatusCode() *int32
	SetBody(v *ReadSchedulerxNotificationPolicyResponseBody) *ReadSchedulerxNotificationPolicyResponse
	GetBody() *ReadSchedulerxNotificationPolicyResponseBody
}

type ReadSchedulerxNotificationPolicyResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReadSchedulerxNotificationPolicyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReadSchedulerxNotificationPolicyResponse) String() string {
	return dara.Prettify(s)
}

func (s ReadSchedulerxNotificationPolicyResponse) GoString() string {
	return s.String()
}

func (s *ReadSchedulerxNotificationPolicyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReadSchedulerxNotificationPolicyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReadSchedulerxNotificationPolicyResponse) GetBody() *ReadSchedulerxNotificationPolicyResponseBody {
	return s.Body
}

func (s *ReadSchedulerxNotificationPolicyResponse) SetHeaders(v map[string]*string) *ReadSchedulerxNotificationPolicyResponse {
	s.Headers = v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponse) SetStatusCode(v int32) *ReadSchedulerxNotificationPolicyResponse {
	s.StatusCode = &v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponse) SetBody(v *ReadSchedulerxNotificationPolicyResponseBody) *ReadSchedulerxNotificationPolicyResponse {
	s.Body = v
	return s
}

func (s *ReadSchedulerxNotificationPolicyResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

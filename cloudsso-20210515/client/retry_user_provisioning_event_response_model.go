// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryUserProvisioningEventResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryUserProvisioningEventResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryUserProvisioningEventResponse
	GetStatusCode() *int32
	SetBody(v *RetryUserProvisioningEventResponseBody) *RetryUserProvisioningEventResponse
	GetBody() *RetryUserProvisioningEventResponseBody
}

type RetryUserProvisioningEventResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryUserProvisioningEventResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryUserProvisioningEventResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryUserProvisioningEventResponse) GoString() string {
	return s.String()
}

func (s *RetryUserProvisioningEventResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryUserProvisioningEventResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryUserProvisioningEventResponse) GetBody() *RetryUserProvisioningEventResponseBody {
	return s.Body
}

func (s *RetryUserProvisioningEventResponse) SetHeaders(v map[string]*string) *RetryUserProvisioningEventResponse {
	s.Headers = v
	return s
}

func (s *RetryUserProvisioningEventResponse) SetStatusCode(v int32) *RetryUserProvisioningEventResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryUserProvisioningEventResponse) SetBody(v *RetryUserProvisioningEventResponseBody) *RetryUserProvisioningEventResponse {
	s.Body = v
	return s
}

func (s *RetryUserProvisioningEventResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

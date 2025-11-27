// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSuspendTriggerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SuspendTriggerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SuspendTriggerResponse
	GetStatusCode() *int32
	SetBody(v *SuspendTriggerResponseBody) *SuspendTriggerResponse
	GetBody() *SuspendTriggerResponseBody
}

type SuspendTriggerResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SuspendTriggerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SuspendTriggerResponse) String() string {
	return dara.Prettify(s)
}

func (s SuspendTriggerResponse) GoString() string {
	return s.String()
}

func (s *SuspendTriggerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SuspendTriggerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SuspendTriggerResponse) GetBody() *SuspendTriggerResponseBody {
	return s.Body
}

func (s *SuspendTriggerResponse) SetHeaders(v map[string]*string) *SuspendTriggerResponse {
	s.Headers = v
	return s
}

func (s *SuspendTriggerResponse) SetStatusCode(v int32) *SuspendTriggerResponse {
	s.StatusCode = &v
	return s
}

func (s *SuspendTriggerResponse) SetBody(v *SuspendTriggerResponseBody) *SuspendTriggerResponse {
	s.Body = v
	return s
}

func (s *SuspendTriggerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerMmsTimerResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerMmsTimerResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerMmsTimerResponse
	GetStatusCode() *int32
	SetBody(v *TriggerMmsTimerResponseBody) *TriggerMmsTimerResponse
	GetBody() *TriggerMmsTimerResponseBody
}

type TriggerMmsTimerResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerMmsTimerResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerMmsTimerResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerMmsTimerResponse) GoString() string {
	return s.String()
}

func (s *TriggerMmsTimerResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerMmsTimerResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerMmsTimerResponse) GetBody() *TriggerMmsTimerResponseBody {
	return s.Body
}

func (s *TriggerMmsTimerResponse) SetHeaders(v map[string]*string) *TriggerMmsTimerResponse {
	s.Headers = v
	return s
}

func (s *TriggerMmsTimerResponse) SetStatusCode(v int32) *TriggerMmsTimerResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerMmsTimerResponse) SetBody(v *TriggerMmsTimerResponseBody) *TriggerMmsTimerResponse {
	s.Body = v
	return s
}

func (s *TriggerMmsTimerResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

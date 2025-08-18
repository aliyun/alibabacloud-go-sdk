// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopScheduledPreloadExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopScheduledPreloadExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopScheduledPreloadExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StopScheduledPreloadExecutionResponseBody) *StopScheduledPreloadExecutionResponse
	GetBody() *StopScheduledPreloadExecutionResponseBody
}

type StopScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopScheduledPreloadExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StopScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *StopScheduledPreloadExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopScheduledPreloadExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopScheduledPreloadExecutionResponse) GetBody() *StopScheduledPreloadExecutionResponseBody {
	return s.Body
}

func (s *StopScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *StopScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *StopScheduledPreloadExecutionResponse) SetStatusCode(v int32) *StopScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StopScheduledPreloadExecutionResponse) SetBody(v *StopScheduledPreloadExecutionResponseBody) *StopScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

func (s *StopScheduledPreloadExecutionResponse) Validate() error {
	return dara.Validate(s)
}

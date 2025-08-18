// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartScheduledPreloadExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartScheduledPreloadExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartScheduledPreloadExecutionResponse
	GetStatusCode() *int32
	SetBody(v *StartScheduledPreloadExecutionResponseBody) *StartScheduledPreloadExecutionResponse
	GetBody() *StartScheduledPreloadExecutionResponseBody
}

type StartScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartScheduledPreloadExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s StartScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *StartScheduledPreloadExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartScheduledPreloadExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartScheduledPreloadExecutionResponse) GetBody() *StartScheduledPreloadExecutionResponseBody {
	return s.Body
}

func (s *StartScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *StartScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *StartScheduledPreloadExecutionResponse) SetStatusCode(v int32) *StartScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *StartScheduledPreloadExecutionResponse) SetBody(v *StartScheduledPreloadExecutionResponseBody) *StartScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

func (s *StartScheduledPreloadExecutionResponse) Validate() error {
	return dara.Validate(s)
}

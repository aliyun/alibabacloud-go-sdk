// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledPreloadExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduledPreloadExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduledPreloadExecutionResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduledPreloadExecutionResponseBody) *UpdateScheduledPreloadExecutionResponse
	GetBody() *UpdateScheduledPreloadExecutionResponseBody
}

type UpdateScheduledPreloadExecutionResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduledPreloadExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledPreloadExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledPreloadExecutionResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledPreloadExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduledPreloadExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduledPreloadExecutionResponse) GetBody() *UpdateScheduledPreloadExecutionResponseBody {
	return s.Body
}

func (s *UpdateScheduledPreloadExecutionResponse) SetHeaders(v map[string]*string) *UpdateScheduledPreloadExecutionResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponse) SetStatusCode(v int32) *UpdateScheduledPreloadExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponse) SetBody(v *UpdateScheduledPreloadExecutionResponseBody) *UpdateScheduledPreloadExecutionResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduledPreloadExecutionResponse) Validate() error {
	return dara.Validate(s)
}

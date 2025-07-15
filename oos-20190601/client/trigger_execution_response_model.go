// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerExecutionResponse
	GetStatusCode() *int32
	SetBody(v *TriggerExecutionResponseBody) *TriggerExecutionResponse
	GetBody() *TriggerExecutionResponseBody
}

type TriggerExecutionResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerExecutionResponse) GoString() string {
	return s.String()
}

func (s *TriggerExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerExecutionResponse) GetBody() *TriggerExecutionResponseBody {
	return s.Body
}

func (s *TriggerExecutionResponse) SetHeaders(v map[string]*string) *TriggerExecutionResponse {
	s.Headers = v
	return s
}

func (s *TriggerExecutionResponse) SetStatusCode(v int32) *TriggerExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerExecutionResponse) SetBody(v *TriggerExecutionResponseBody) *TriggerExecutionResponse {
	s.Body = v
	return s
}

func (s *TriggerExecutionResponse) Validate() error {
	return dara.Validate(s)
}

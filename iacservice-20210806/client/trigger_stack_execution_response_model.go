// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerStackExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TriggerStackExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TriggerStackExecutionResponse
	GetStatusCode() *int32
	SetBody(v *TriggerStackExecutionResponseBody) *TriggerStackExecutionResponse
	GetBody() *TriggerStackExecutionResponseBody
}

type TriggerStackExecutionResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TriggerStackExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TriggerStackExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s TriggerStackExecutionResponse) GoString() string {
	return s.String()
}

func (s *TriggerStackExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TriggerStackExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TriggerStackExecutionResponse) GetBody() *TriggerStackExecutionResponseBody {
	return s.Body
}

func (s *TriggerStackExecutionResponse) SetHeaders(v map[string]*string) *TriggerStackExecutionResponse {
	s.Headers = v
	return s
}

func (s *TriggerStackExecutionResponse) SetStatusCode(v int32) *TriggerStackExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *TriggerStackExecutionResponse) SetBody(v *TriggerStackExecutionResponseBody) *TriggerStackExecutionResponse {
	s.Body = v
	return s
}

func (s *TriggerStackExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

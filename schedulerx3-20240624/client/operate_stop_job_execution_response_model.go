// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateStopJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateStopJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateStopJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateStopJobExecutionResponseBody) *OperateStopJobExecutionResponse
	GetBody() *OperateStopJobExecutionResponseBody
}

type OperateStopJobExecutionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateStopJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateStopJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateStopJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateStopJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateStopJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateStopJobExecutionResponse) GetBody() *OperateStopJobExecutionResponseBody {
	return s.Body
}

func (s *OperateStopJobExecutionResponse) SetHeaders(v map[string]*string) *OperateStopJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateStopJobExecutionResponse) SetStatusCode(v int32) *OperateStopJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateStopJobExecutionResponse) SetBody(v *OperateStopJobExecutionResponseBody) *OperateStopJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateStopJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

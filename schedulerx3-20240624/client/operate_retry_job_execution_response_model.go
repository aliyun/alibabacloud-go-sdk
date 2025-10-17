// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateRetryJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateRetryJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateRetryJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateRetryJobExecutionResponseBody) *OperateRetryJobExecutionResponse
	GetBody() *OperateRetryJobExecutionResponseBody
}

type OperateRetryJobExecutionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateRetryJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateRetryJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateRetryJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateRetryJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateRetryJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateRetryJobExecutionResponse) GetBody() *OperateRetryJobExecutionResponseBody {
	return s.Body
}

func (s *OperateRetryJobExecutionResponse) SetHeaders(v map[string]*string) *OperateRetryJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateRetryJobExecutionResponse) SetStatusCode(v int32) *OperateRetryJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateRetryJobExecutionResponse) SetBody(v *OperateRetryJobExecutionResponseBody) *OperateRetryJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateRetryJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

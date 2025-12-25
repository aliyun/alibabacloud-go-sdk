// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateSkipJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateSkipJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateSkipJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateSkipJobExecutionResponseBody) *OperateSkipJobExecutionResponse
	GetBody() *OperateSkipJobExecutionResponseBody
}

type OperateSkipJobExecutionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateSkipJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateSkipJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateSkipJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateSkipJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateSkipJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateSkipJobExecutionResponse) GetBody() *OperateSkipJobExecutionResponseBody {
	return s.Body
}

func (s *OperateSkipJobExecutionResponse) SetHeaders(v map[string]*string) *OperateSkipJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateSkipJobExecutionResponse) SetStatusCode(v int32) *OperateSkipJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateSkipJobExecutionResponse) SetBody(v *OperateSkipJobExecutionResponseBody) *OperateSkipJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateSkipJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

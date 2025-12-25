// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnskipJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateUnskipJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateUnskipJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateUnskipJobExecutionResponseBody) *OperateUnskipJobExecutionResponse
	GetBody() *OperateUnskipJobExecutionResponseBody
}

type OperateUnskipJobExecutionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateUnskipJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateUnskipJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateUnskipJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateUnskipJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateUnskipJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateUnskipJobExecutionResponse) GetBody() *OperateUnskipJobExecutionResponseBody {
	return s.Body
}

func (s *OperateUnskipJobExecutionResponse) SetHeaders(v map[string]*string) *OperateUnskipJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateUnskipJobExecutionResponse) SetStatusCode(v int32) *OperateUnskipJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateUnskipJobExecutionResponse) SetBody(v *OperateUnskipJobExecutionResponseBody) *OperateUnskipJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateUnskipJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

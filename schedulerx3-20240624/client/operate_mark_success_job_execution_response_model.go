// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateMarkSuccessJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateMarkSuccessJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateMarkSuccessJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateMarkSuccessJobExecutionResponseBody) *OperateMarkSuccessJobExecutionResponse
	GetBody() *OperateMarkSuccessJobExecutionResponseBody
}

type OperateMarkSuccessJobExecutionResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateMarkSuccessJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateMarkSuccessJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateMarkSuccessJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateMarkSuccessJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateMarkSuccessJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateMarkSuccessJobExecutionResponse) GetBody() *OperateMarkSuccessJobExecutionResponseBody {
	return s.Body
}

func (s *OperateMarkSuccessJobExecutionResponse) SetHeaders(v map[string]*string) *OperateMarkSuccessJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponse) SetStatusCode(v int32) *OperateMarkSuccessJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponse) SetBody(v *OperateMarkSuccessJobExecutionResponseBody) *OperateMarkSuccessJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateMarkSuccessJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

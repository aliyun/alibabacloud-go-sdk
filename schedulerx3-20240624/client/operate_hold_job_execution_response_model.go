// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateHoldJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateHoldJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateHoldJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateHoldJobExecutionResponseBody) *OperateHoldJobExecutionResponse
	GetBody() *OperateHoldJobExecutionResponseBody
}

type OperateHoldJobExecutionResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateHoldJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateHoldJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateHoldJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateHoldJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateHoldJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateHoldJobExecutionResponse) GetBody() *OperateHoldJobExecutionResponseBody {
	return s.Body
}

func (s *OperateHoldJobExecutionResponse) SetHeaders(v map[string]*string) *OperateHoldJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateHoldJobExecutionResponse) SetStatusCode(v int32) *OperateHoldJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateHoldJobExecutionResponse) SetBody(v *OperateHoldJobExecutionResponseBody) *OperateHoldJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateHoldJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

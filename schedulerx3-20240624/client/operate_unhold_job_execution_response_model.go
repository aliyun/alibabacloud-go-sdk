// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateUnholdJobExecutionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateUnholdJobExecutionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateUnholdJobExecutionResponse
	GetStatusCode() *int32
	SetBody(v *OperateUnholdJobExecutionResponseBody) *OperateUnholdJobExecutionResponse
	GetBody() *OperateUnholdJobExecutionResponseBody
}

type OperateUnholdJobExecutionResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateUnholdJobExecutionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateUnholdJobExecutionResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateUnholdJobExecutionResponse) GoString() string {
	return s.String()
}

func (s *OperateUnholdJobExecutionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateUnholdJobExecutionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateUnholdJobExecutionResponse) GetBody() *OperateUnholdJobExecutionResponseBody {
	return s.Body
}

func (s *OperateUnholdJobExecutionResponse) SetHeaders(v map[string]*string) *OperateUnholdJobExecutionResponse {
	s.Headers = v
	return s
}

func (s *OperateUnholdJobExecutionResponse) SetStatusCode(v int32) *OperateUnholdJobExecutionResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateUnholdJobExecutionResponse) SetBody(v *OperateUnholdJobExecutionResponseBody) *OperateUnholdJobExecutionResponse {
	s.Body = v
	return s
}

func (s *OperateUnholdJobExecutionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

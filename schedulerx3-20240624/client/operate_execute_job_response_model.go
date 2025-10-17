// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOperateExecuteJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OperateExecuteJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OperateExecuteJobResponse
	GetStatusCode() *int32
	SetBody(v *OperateExecuteJobResponseBody) *OperateExecuteJobResponse
	GetBody() *OperateExecuteJobResponseBody
}

type OperateExecuteJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OperateExecuteJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OperateExecuteJobResponse) String() string {
	return dara.Prettify(s)
}

func (s OperateExecuteJobResponse) GoString() string {
	return s.String()
}

func (s *OperateExecuteJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OperateExecuteJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OperateExecuteJobResponse) GetBody() *OperateExecuteJobResponseBody {
	return s.Body
}

func (s *OperateExecuteJobResponse) SetHeaders(v map[string]*string) *OperateExecuteJobResponse {
	s.Headers = v
	return s
}

func (s *OperateExecuteJobResponse) SetStatusCode(v int32) *OperateExecuteJobResponse {
	s.StatusCode = &v
	return s
}

func (s *OperateExecuteJobResponse) SetBody(v *OperateExecuteJobResponseBody) *OperateExecuteJobResponse {
	s.Body = v
	return s
}

func (s *OperateExecuteJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetryVMDeployMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RetryVMDeployMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RetryVMDeployMachineResponse
	GetStatusCode() *int32
	SetBody(v *RetryVMDeployMachineResponseBody) *RetryVMDeployMachineResponse
	GetBody() *RetryVMDeployMachineResponseBody
}

type RetryVMDeployMachineResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RetryVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RetryVMDeployMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s RetryVMDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *RetryVMDeployMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RetryVMDeployMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RetryVMDeployMachineResponse) GetBody() *RetryVMDeployMachineResponseBody {
	return s.Body
}

func (s *RetryVMDeployMachineResponse) SetHeaders(v map[string]*string) *RetryVMDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *RetryVMDeployMachineResponse) SetStatusCode(v int32) *RetryVMDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *RetryVMDeployMachineResponse) SetBody(v *RetryVMDeployMachineResponseBody) *RetryVMDeployMachineResponse {
	s.Body = v
	return s
}

func (s *RetryVMDeployMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

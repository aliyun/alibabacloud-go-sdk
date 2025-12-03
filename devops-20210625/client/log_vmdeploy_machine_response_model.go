// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogVMDeployMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *LogVMDeployMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *LogVMDeployMachineResponse
	GetStatusCode() *int32
	SetBody(v *LogVMDeployMachineResponseBody) *LogVMDeployMachineResponse
	GetBody() *LogVMDeployMachineResponseBody
}

type LogVMDeployMachineResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *LogVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LogVMDeployMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s LogVMDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *LogVMDeployMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *LogVMDeployMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *LogVMDeployMachineResponse) GetBody() *LogVMDeployMachineResponseBody {
	return s.Body
}

func (s *LogVMDeployMachineResponse) SetHeaders(v map[string]*string) *LogVMDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *LogVMDeployMachineResponse) SetStatusCode(v int32) *LogVMDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *LogVMDeployMachineResponse) SetBody(v *LogVMDeployMachineResponseBody) *LogVMDeployMachineResponse {
	s.Body = v
	return s
}

func (s *LogVMDeployMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

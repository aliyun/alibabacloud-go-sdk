// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSkipVMDeployMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SkipVMDeployMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SkipVMDeployMachineResponse
	GetStatusCode() *int32
	SetBody(v *SkipVMDeployMachineResponseBody) *SkipVMDeployMachineResponse
	GetBody() *SkipVMDeployMachineResponseBody
}

type SkipVMDeployMachineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SkipVMDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SkipVMDeployMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s SkipVMDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *SkipVMDeployMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SkipVMDeployMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SkipVMDeployMachineResponse) GetBody() *SkipVMDeployMachineResponseBody {
	return s.Body
}

func (s *SkipVMDeployMachineResponse) SetHeaders(v map[string]*string) *SkipVMDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *SkipVMDeployMachineResponse) SetStatusCode(v int32) *SkipVMDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *SkipVMDeployMachineResponse) SetBody(v *SkipVMDeployMachineResponseBody) *SkipVMDeployMachineResponse {
	s.Body = v
	return s
}

func (s *SkipVMDeployMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

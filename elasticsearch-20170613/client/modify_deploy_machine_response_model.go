// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeployMachineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDeployMachineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDeployMachineResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDeployMachineResponseBody) *ModifyDeployMachineResponse
	GetBody() *ModifyDeployMachineResponseBody
}

type ModifyDeployMachineResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDeployMachineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDeployMachineResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeployMachineResponse) GoString() string {
	return s.String()
}

func (s *ModifyDeployMachineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDeployMachineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDeployMachineResponse) GetBody() *ModifyDeployMachineResponseBody {
	return s.Body
}

func (s *ModifyDeployMachineResponse) SetHeaders(v map[string]*string) *ModifyDeployMachineResponse {
	s.Headers = v
	return s
}

func (s *ModifyDeployMachineResponse) SetStatusCode(v int32) *ModifyDeployMachineResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDeployMachineResponse) SetBody(v *ModifyDeployMachineResponseBody) *ModifyDeployMachineResponse {
	s.Body = v
	return s
}

func (s *ModifyDeployMachineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

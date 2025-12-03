// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopVMDeployOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopVMDeployOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopVMDeployOrderResponse
	GetStatusCode() *int32
	SetBody(v *StopVMDeployOrderResponseBody) *StopVMDeployOrderResponse
	GetBody() *StopVMDeployOrderResponseBody
}

type StopVMDeployOrderResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopVMDeployOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s StopVMDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *StopVMDeployOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopVMDeployOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopVMDeployOrderResponse) GetBody() *StopVMDeployOrderResponseBody {
	return s.Body
}

func (s *StopVMDeployOrderResponse) SetHeaders(v map[string]*string) *StopVMDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *StopVMDeployOrderResponse) SetStatusCode(v int32) *StopVMDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *StopVMDeployOrderResponse) SetBody(v *StopVMDeployOrderResponseBody) *StopVMDeployOrderResponse {
	s.Body = v
	return s
}

func (s *StopVMDeployOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

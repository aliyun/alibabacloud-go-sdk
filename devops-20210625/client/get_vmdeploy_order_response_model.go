// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVMDeployOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetVMDeployOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetVMDeployOrderResponse
	GetStatusCode() *int32
	SetBody(v *GetVMDeployOrderResponseBody) *GetVMDeployOrderResponse
	GetBody() *GetVMDeployOrderResponseBody
}

type GetVMDeployOrderResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVMDeployOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s GetVMDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *GetVMDeployOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetVMDeployOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetVMDeployOrderResponse) GetBody() *GetVMDeployOrderResponseBody {
	return s.Body
}

func (s *GetVMDeployOrderResponse) SetHeaders(v map[string]*string) *GetVMDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *GetVMDeployOrderResponse) SetStatusCode(v int32) *GetVMDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVMDeployOrderResponse) SetBody(v *GetVMDeployOrderResponseBody) *GetVMDeployOrderResponse {
	s.Body = v
	return s
}

func (s *GetVMDeployOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGlobalAccelerationInstanceIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGlobalAccelerationInstanceIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGlobalAccelerationInstanceIpResponse
	GetStatusCode() *int32
	SetBody(v *AddGlobalAccelerationInstanceIpResponseBody) *AddGlobalAccelerationInstanceIpResponse
	GetBody() *AddGlobalAccelerationInstanceIpResponseBody
}

type AddGlobalAccelerationInstanceIpResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGlobalAccelerationInstanceIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGlobalAccelerationInstanceIpResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGlobalAccelerationInstanceIpResponse) GoString() string {
	return s.String()
}

func (s *AddGlobalAccelerationInstanceIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGlobalAccelerationInstanceIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGlobalAccelerationInstanceIpResponse) GetBody() *AddGlobalAccelerationInstanceIpResponseBody {
	return s.Body
}

func (s *AddGlobalAccelerationInstanceIpResponse) SetHeaders(v map[string]*string) *AddGlobalAccelerationInstanceIpResponse {
	s.Headers = v
	return s
}

func (s *AddGlobalAccelerationInstanceIpResponse) SetStatusCode(v int32) *AddGlobalAccelerationInstanceIpResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGlobalAccelerationInstanceIpResponse) SetBody(v *AddGlobalAccelerationInstanceIpResponseBody) *AddGlobalAccelerationInstanceIpResponse {
	s.Body = v
	return s
}

func (s *AddGlobalAccelerationInstanceIpResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

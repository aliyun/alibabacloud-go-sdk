// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveGlobalAccelerationInstanceIpResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveGlobalAccelerationInstanceIpResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveGlobalAccelerationInstanceIpResponse
	GetStatusCode() *int32
	SetBody(v *RemoveGlobalAccelerationInstanceIpResponseBody) *RemoveGlobalAccelerationInstanceIpResponse
	GetBody() *RemoveGlobalAccelerationInstanceIpResponseBody
}

type RemoveGlobalAccelerationInstanceIpResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveGlobalAccelerationInstanceIpResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveGlobalAccelerationInstanceIpResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveGlobalAccelerationInstanceIpResponse) GoString() string {
	return s.String()
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) GetBody() *RemoveGlobalAccelerationInstanceIpResponseBody {
	return s.Body
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) SetHeaders(v map[string]*string) *RemoveGlobalAccelerationInstanceIpResponse {
	s.Headers = v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) SetStatusCode(v int32) *RemoveGlobalAccelerationInstanceIpResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) SetBody(v *RemoveGlobalAccelerationInstanceIpResponseBody) *RemoveGlobalAccelerationInstanceIpResponse {
	s.Body = v
	return s
}

func (s *RemoveGlobalAccelerationInstanceIpResponse) Validate() error {
	return dara.Validate(s)
}

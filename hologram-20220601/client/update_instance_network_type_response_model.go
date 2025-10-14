// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNetworkTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceNetworkTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceNetworkTypeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceNetworkTypeResponseBody) *UpdateInstanceNetworkTypeResponse
	GetBody() *UpdateInstanceNetworkTypeResponseBody
}

type UpdateInstanceNetworkTypeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNetworkTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNetworkTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNetworkTypeResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceNetworkTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceNetworkTypeResponse) GetBody() *UpdateInstanceNetworkTypeResponseBody {
	return s.Body
}

func (s *UpdateInstanceNetworkTypeResponse) SetHeaders(v map[string]*string) *UpdateInstanceNetworkTypeResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) SetStatusCode(v int32) *UpdateInstanceNetworkTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) SetBody(v *UpdateInstanceNetworkTypeResponseBody) *UpdateInstanceNetworkTypeResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceNetworkTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

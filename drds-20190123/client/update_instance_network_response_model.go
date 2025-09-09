// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateInstanceNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateInstanceNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateInstanceNetworkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateInstanceNetworkResponseBody) *UpdateInstanceNetworkResponse
	GetBody() *UpdateInstanceNetworkResponseBody
}

type UpdateInstanceNetworkResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateInstanceNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateInstanceNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateInstanceNetworkResponse) GoString() string {
	return s.String()
}

func (s *UpdateInstanceNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateInstanceNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateInstanceNetworkResponse) GetBody() *UpdateInstanceNetworkResponseBody {
	return s.Body
}

func (s *UpdateInstanceNetworkResponse) SetHeaders(v map[string]*string) *UpdateInstanceNetworkResponse {
	s.Headers = v
	return s
}

func (s *UpdateInstanceNetworkResponse) SetStatusCode(v int32) *UpdateInstanceNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateInstanceNetworkResponse) SetBody(v *UpdateInstanceNetworkResponseBody) *UpdateInstanceNetworkResponse {
	s.Body = v
	return s
}

func (s *UpdateInstanceNetworkResponse) Validate() error {
	return dara.Validate(s)
}

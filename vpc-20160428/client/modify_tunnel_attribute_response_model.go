// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTunnelAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyTunnelAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyTunnelAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyTunnelAttributeResponseBody) *ModifyTunnelAttributeResponse
	GetBody() *ModifyTunnelAttributeResponseBody
}

type ModifyTunnelAttributeResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyTunnelAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyTunnelAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyTunnelAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyTunnelAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyTunnelAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyTunnelAttributeResponse) GetBody() *ModifyTunnelAttributeResponseBody {
	return s.Body
}

func (s *ModifyTunnelAttributeResponse) SetHeaders(v map[string]*string) *ModifyTunnelAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyTunnelAttributeResponse) SetStatusCode(v int32) *ModifyTunnelAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyTunnelAttributeResponse) SetBody(v *ModifyTunnelAttributeResponseBody) *ModifyTunnelAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyTunnelAttributeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateEgressProxyIpPoolAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateEgressProxyIpPoolAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateEgressProxyIpPoolAttributeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateEgressProxyIpPoolAttributeResponseBody) *UpdateEgressProxyIpPoolAttributeResponse
	GetBody() *UpdateEgressProxyIpPoolAttributeResponseBody
}

type UpdateEgressProxyIpPoolAttributeResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateEgressProxyIpPoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateEgressProxyIpPoolAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateEgressProxyIpPoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) GetBody() *UpdateEgressProxyIpPoolAttributeResponseBody {
	return s.Body
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) SetHeaders(v map[string]*string) *UpdateEgressProxyIpPoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) SetStatusCode(v int32) *UpdateEgressProxyIpPoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) SetBody(v *UpdateEgressProxyIpPoolAttributeResponseBody) *UpdateEgressProxyIpPoolAttributeResponse {
	s.Body = v
	return s
}

func (s *UpdateEgressProxyIpPoolAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

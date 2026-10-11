// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEgressProxyIpPoolAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetEgressProxyIpPoolAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetEgressProxyIpPoolAttributeResponse
	GetStatusCode() *int32
	SetBody(v *GetEgressProxyIpPoolAttributeResponseBody) *GetEgressProxyIpPoolAttributeResponse
	GetBody() *GetEgressProxyIpPoolAttributeResponseBody
}

type GetEgressProxyIpPoolAttributeResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetEgressProxyIpPoolAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetEgressProxyIpPoolAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetEgressProxyIpPoolAttributeResponse) GoString() string {
	return s.String()
}

func (s *GetEgressProxyIpPoolAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetEgressProxyIpPoolAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetEgressProxyIpPoolAttributeResponse) GetBody() *GetEgressProxyIpPoolAttributeResponseBody {
	return s.Body
}

func (s *GetEgressProxyIpPoolAttributeResponse) SetHeaders(v map[string]*string) *GetEgressProxyIpPoolAttributeResponse {
	s.Headers = v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponse) SetStatusCode(v int32) *GetEgressProxyIpPoolAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponse) SetBody(v *GetEgressProxyIpPoolAttributeResponseBody) *GetEgressProxyIpPoolAttributeResponse {
	s.Body = v
	return s
}

func (s *GetEgressProxyIpPoolAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

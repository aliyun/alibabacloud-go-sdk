// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEgressProxyIpPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEgressProxyIpPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEgressProxyIpPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateEgressProxyIpPoolResponseBody) *CreateEgressProxyIpPoolResponse
	GetBody() *CreateEgressProxyIpPoolResponseBody
}

type CreateEgressProxyIpPoolResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEgressProxyIpPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEgressProxyIpPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEgressProxyIpPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateEgressProxyIpPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEgressProxyIpPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEgressProxyIpPoolResponse) GetBody() *CreateEgressProxyIpPoolResponseBody {
	return s.Body
}

func (s *CreateEgressProxyIpPoolResponse) SetHeaders(v map[string]*string) *CreateEgressProxyIpPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateEgressProxyIpPoolResponse) SetStatusCode(v int32) *CreateEgressProxyIpPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEgressProxyIpPoolResponse) SetBody(v *CreateEgressProxyIpPoolResponseBody) *CreateEgressProxyIpPoolResponse {
	s.Body = v
	return s
}

func (s *CreateEgressProxyIpPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

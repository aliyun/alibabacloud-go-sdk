// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEgressProxyIpPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEgressProxyIpPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEgressProxyIpPoolResponse
	GetStatusCode() *int32
	SetBody(v *DeleteEgressProxyIpPoolResponseBody) *DeleteEgressProxyIpPoolResponse
	GetBody() *DeleteEgressProxyIpPoolResponseBody
}

type DeleteEgressProxyIpPoolResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteEgressProxyIpPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteEgressProxyIpPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEgressProxyIpPoolResponse) GoString() string {
	return s.String()
}

func (s *DeleteEgressProxyIpPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEgressProxyIpPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEgressProxyIpPoolResponse) GetBody() *DeleteEgressProxyIpPoolResponseBody {
	return s.Body
}

func (s *DeleteEgressProxyIpPoolResponse) SetHeaders(v map[string]*string) *DeleteEgressProxyIpPoolResponse {
	s.Headers = v
	return s
}

func (s *DeleteEgressProxyIpPoolResponse) SetStatusCode(v int32) *DeleteEgressProxyIpPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEgressProxyIpPoolResponse) SetBody(v *DeleteEgressProxyIpPoolResponseBody) *DeleteEgressProxyIpPoolResponse {
	s.Body = v
	return s
}

func (s *DeleteEgressProxyIpPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

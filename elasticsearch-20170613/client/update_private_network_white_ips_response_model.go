// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdatePrivateNetworkWhiteIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdatePrivateNetworkWhiteIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdatePrivateNetworkWhiteIpsResponse
	GetStatusCode() *int32
	SetBody(v *UpdatePrivateNetworkWhiteIpsResponseBody) *UpdatePrivateNetworkWhiteIpsResponse
	GetBody() *UpdatePrivateNetworkWhiteIpsResponseBody
}

type UpdatePrivateNetworkWhiteIpsResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdatePrivateNetworkWhiteIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdatePrivateNetworkWhiteIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdatePrivateNetworkWhiteIpsResponse) GoString() string {
	return s.String()
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) GetBody() *UpdatePrivateNetworkWhiteIpsResponseBody {
	return s.Body
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) SetHeaders(v map[string]*string) *UpdatePrivateNetworkWhiteIpsResponse {
	s.Headers = v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) SetStatusCode(v int32) *UpdatePrivateNetworkWhiteIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) SetBody(v *UpdatePrivateNetworkWhiteIpsResponseBody) *UpdatePrivateNetworkWhiteIpsResponse {
	s.Body = v
	return s
}

func (s *UpdatePrivateNetworkWhiteIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

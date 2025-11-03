// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySslVpnClientCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifySslVpnClientCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifySslVpnClientCertResponse
	GetStatusCode() *int32
	SetBody(v *ModifySslVpnClientCertResponseBody) *ModifySslVpnClientCertResponse
	GetBody() *ModifySslVpnClientCertResponseBody
}

type ModifySslVpnClientCertResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifySslVpnClientCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifySslVpnClientCertResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifySslVpnClientCertResponse) GoString() string {
	return s.String()
}

func (s *ModifySslVpnClientCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifySslVpnClientCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifySslVpnClientCertResponse) GetBody() *ModifySslVpnClientCertResponseBody {
	return s.Body
}

func (s *ModifySslVpnClientCertResponse) SetHeaders(v map[string]*string) *ModifySslVpnClientCertResponse {
	s.Headers = v
	return s
}

func (s *ModifySslVpnClientCertResponse) SetStatusCode(v int32) *ModifySslVpnClientCertResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifySslVpnClientCertResponse) SetBody(v *ModifySslVpnClientCertResponseBody) *ModifySslVpnClientCertResponse {
	s.Body = v
	return s
}

func (s *ModifySslVpnClientCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

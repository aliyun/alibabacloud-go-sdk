// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientCertResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSslVpnClientCertResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSslVpnClientCertResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSslVpnClientCertResponseBody) *DescribeSslVpnClientCertResponse
	GetBody() *DescribeSslVpnClientCertResponseBody
}

type DescribeSslVpnClientCertResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSslVpnClientCertResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSslVpnClientCertResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertResponse) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSslVpnClientCertResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSslVpnClientCertResponse) GetBody() *DescribeSslVpnClientCertResponseBody {
	return s.Body
}

func (s *DescribeSslVpnClientCertResponse) SetHeaders(v map[string]*string) *DescribeSslVpnClientCertResponse {
	s.Headers = v
	return s
}

func (s *DescribeSslVpnClientCertResponse) SetStatusCode(v int32) *DescribeSslVpnClientCertResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSslVpnClientCertResponse) SetBody(v *DescribeSslVpnClientCertResponseBody) *DescribeSslVpnClientCertResponse {
	s.Body = v
	return s
}

func (s *DescribeSslVpnClientCertResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

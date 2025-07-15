// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnClientCertsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSslVpnClientCertsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSslVpnClientCertsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSslVpnClientCertsResponseBody) *DescribeSslVpnClientCertsResponse
	GetBody() *DescribeSslVpnClientCertsResponseBody
}

type DescribeSslVpnClientCertsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSslVpnClientCertsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSslVpnClientCertsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnClientCertsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnClientCertsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSslVpnClientCertsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSslVpnClientCertsResponse) GetBody() *DescribeSslVpnClientCertsResponseBody {
	return s.Body
}

func (s *DescribeSslVpnClientCertsResponse) SetHeaders(v map[string]*string) *DescribeSslVpnClientCertsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSslVpnClientCertsResponse) SetStatusCode(v int32) *DescribeSslVpnClientCertsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSslVpnClientCertsResponse) SetBody(v *DescribeSslVpnClientCertsResponseBody) *DescribeSslVpnClientCertsResponse {
	s.Body = v
	return s
}

func (s *DescribeSslVpnClientCertsResponse) Validate() error {
	return dara.Validate(s)
}

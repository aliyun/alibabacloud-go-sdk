// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSslVpnServersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSslVpnServersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSslVpnServersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSslVpnServersResponseBody) *DescribeSslVpnServersResponse
	GetBody() *DescribeSslVpnServersResponseBody
}

type DescribeSslVpnServersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSslVpnServersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSslVpnServersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSslVpnServersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSslVpnServersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSslVpnServersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSslVpnServersResponse) GetBody() *DescribeSslVpnServersResponseBody {
	return s.Body
}

func (s *DescribeSslVpnServersResponse) SetHeaders(v map[string]*string) *DescribeSslVpnServersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSslVpnServersResponse) SetStatusCode(v int32) *DescribeSslVpnServersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSslVpnServersResponse) SetBody(v *DescribeSslVpnServersResponseBody) *DescribeSslVpnServersResponse {
	s.Body = v
	return s
}

func (s *DescribeSslVpnServersResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfuseSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfuseSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfuseSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfuseSecurityIpsResponseBody) *DescribePxfuseSecurityIpsResponse
	GetBody() *DescribePxfuseSecurityIpsResponseBody
}

type DescribePxfuseSecurityIpsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfuseSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfuseSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfuseSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfuseSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfuseSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfuseSecurityIpsResponse) GetBody() *DescribePxfuseSecurityIpsResponseBody {
	return s.Body
}

func (s *DescribePxfuseSecurityIpsResponse) SetHeaders(v map[string]*string) *DescribePxfuseSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfuseSecurityIpsResponse) SetStatusCode(v int32) *DescribePxfuseSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfuseSecurityIpsResponse) SetBody(v *DescribePxfuseSecurityIpsResponseBody) *DescribePxfuseSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *DescribePxfuseSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

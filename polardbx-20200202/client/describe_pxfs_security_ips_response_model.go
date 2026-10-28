// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePxfsSecurityIpsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePxfsSecurityIpsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePxfsSecurityIpsResponse
	GetStatusCode() *int32
	SetBody(v *DescribePxfsSecurityIpsResponseBody) *DescribePxfsSecurityIpsResponse
	GetBody() *DescribePxfsSecurityIpsResponseBody
}

type DescribePxfsSecurityIpsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePxfsSecurityIpsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePxfsSecurityIpsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePxfsSecurityIpsResponse) GoString() string {
	return s.String()
}

func (s *DescribePxfsSecurityIpsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePxfsSecurityIpsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePxfsSecurityIpsResponse) GetBody() *DescribePxfsSecurityIpsResponseBody {
	return s.Body
}

func (s *DescribePxfsSecurityIpsResponse) SetHeaders(v map[string]*string) *DescribePxfsSecurityIpsResponse {
	s.Headers = v
	return s
}

func (s *DescribePxfsSecurityIpsResponse) SetStatusCode(v int32) *DescribePxfsSecurityIpsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePxfsSecurityIpsResponse) SetBody(v *DescribePxfsSecurityIpsResponseBody) *DescribePxfsSecurityIpsResponse {
	s.Body = v
	return s
}

func (s *DescribePxfsSecurityIpsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

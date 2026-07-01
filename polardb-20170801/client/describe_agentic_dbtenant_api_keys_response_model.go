// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAgenticDBTenantApiKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAgenticDBTenantApiKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAgenticDBTenantApiKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAgenticDBTenantApiKeysResponseBody) *DescribeAgenticDBTenantApiKeysResponse
	GetBody() *DescribeAgenticDBTenantApiKeysResponseBody
}

type DescribeAgenticDBTenantApiKeysResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAgenticDBTenantApiKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAgenticDBTenantApiKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAgenticDBTenantApiKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeAgenticDBTenantApiKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAgenticDBTenantApiKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAgenticDBTenantApiKeysResponse) GetBody() *DescribeAgenticDBTenantApiKeysResponseBody {
	return s.Body
}

func (s *DescribeAgenticDBTenantApiKeysResponse) SetHeaders(v map[string]*string) *DescribeAgenticDBTenantApiKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponse) SetStatusCode(v int32) *DescribeAgenticDBTenantApiKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponse) SetBody(v *DescribeAgenticDBTenantApiKeysResponseBody) *DescribeAgenticDBTenantApiKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeAgenticDBTenantApiKeysResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

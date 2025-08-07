// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDatabaseNetworksResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalDatabaseNetworksResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalDatabaseNetworksResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalDatabaseNetworksResponseBody) *DescribeGlobalDatabaseNetworksResponse
	GetBody() *DescribeGlobalDatabaseNetworksResponseBody
}

type DescribeGlobalDatabaseNetworksResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDatabaseNetworksResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDatabaseNetworksResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworksResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworksResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalDatabaseNetworksResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalDatabaseNetworksResponse) GetBody() *DescribeGlobalDatabaseNetworksResponseBody {
	return s.Body
}

func (s *DescribeGlobalDatabaseNetworksResponse) SetHeaders(v map[string]*string) *DescribeGlobalDatabaseNetworksResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponse) SetStatusCode(v int32) *DescribeGlobalDatabaseNetworksResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponse) SetBody(v *DescribeGlobalDatabaseNetworksResponseBody) *DescribeGlobalDatabaseNetworksResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalDatabaseNetworksResponse) Validate() error {
	return dara.Validate(s)
}

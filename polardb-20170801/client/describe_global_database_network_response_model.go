// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalDatabaseNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalDatabaseNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalDatabaseNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalDatabaseNetworkResponseBody) *DescribeGlobalDatabaseNetworkResponse
	GetBody() *DescribeGlobalDatabaseNetworkResponseBody
}

type DescribeGlobalDatabaseNetworkResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalDatabaseNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalDatabaseNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalDatabaseNetworkResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalDatabaseNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalDatabaseNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalDatabaseNetworkResponse) GetBody() *DescribeGlobalDatabaseNetworkResponseBody {
	return s.Body
}

func (s *DescribeGlobalDatabaseNetworkResponse) SetHeaders(v map[string]*string) *DescribeGlobalDatabaseNetworkResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponse) SetStatusCode(v int32) *DescribeGlobalDatabaseNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponse) SetBody(v *DescribeGlobalDatabaseNetworkResponseBody) *DescribeGlobalDatabaseNetworkResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalDatabaseNetworkResponse) Validate() error {
	return dara.Validate(s)
}

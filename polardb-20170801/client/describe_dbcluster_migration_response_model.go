// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterMigrationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterMigrationResponseBody) *DescribeDBClusterMigrationResponse
	GetBody() *DescribeDBClusterMigrationResponseBody
}

type DescribeDBClusterMigrationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterMigrationResponse) GetBody() *DescribeDBClusterMigrationResponseBody {
	return s.Body
}

func (s *DescribeDBClusterMigrationResponse) SetHeaders(v map[string]*string) *DescribeDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetStatusCode(v int32) *DescribeDBClusterMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterMigrationResponse) SetBody(v *DescribeDBClusterMigrationResponseBody) *DescribeDBClusterMigrationResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterMigrationResponse) Validate() error {
	return dara.Validate(s)
}

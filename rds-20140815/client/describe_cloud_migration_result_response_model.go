// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMigrationResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudMigrationResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudMigrationResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudMigrationResultResponseBody) *DescribeCloudMigrationResultResponse
	GetBody() *DescribeCloudMigrationResultResponseBody
}

type DescribeCloudMigrationResultResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudMigrationResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudMigrationResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudMigrationResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudMigrationResultResponse) GetBody() *DescribeCloudMigrationResultResponseBody {
	return s.Body
}

func (s *DescribeCloudMigrationResultResponse) SetHeaders(v map[string]*string) *DescribeCloudMigrationResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudMigrationResultResponse) SetStatusCode(v int32) *DescribeCloudMigrationResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudMigrationResultResponse) SetBody(v *DescribeCloudMigrationResultResponseBody) *DescribeCloudMigrationResultResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudMigrationResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

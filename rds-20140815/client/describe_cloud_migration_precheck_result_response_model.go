// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudMigrationPrecheckResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCloudMigrationPrecheckResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCloudMigrationPrecheckResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCloudMigrationPrecheckResultResponseBody) *DescribeCloudMigrationPrecheckResultResponse
	GetBody() *DescribeCloudMigrationPrecheckResultResponseBody
}

type DescribeCloudMigrationPrecheckResultResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCloudMigrationPrecheckResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCloudMigrationPrecheckResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudMigrationPrecheckResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeCloudMigrationPrecheckResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCloudMigrationPrecheckResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCloudMigrationPrecheckResultResponse) GetBody() *DescribeCloudMigrationPrecheckResultResponseBody {
	return s.Body
}

func (s *DescribeCloudMigrationPrecheckResultResponse) SetHeaders(v map[string]*string) *DescribeCloudMigrationPrecheckResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponse) SetStatusCode(v int32) *DescribeCloudMigrationPrecheckResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponse) SetBody(v *DescribeCloudMigrationPrecheckResultResponseBody) *DescribeCloudMigrationPrecheckResultResponse {
	s.Body = v
	return s
}

func (s *DescribeCloudMigrationPrecheckResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

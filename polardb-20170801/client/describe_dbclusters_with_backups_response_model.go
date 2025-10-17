// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClustersWithBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClustersWithBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClustersWithBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClustersWithBackupsResponseBody) *DescribeDBClustersWithBackupsResponse
	GetBody() *DescribeDBClustersWithBackupsResponseBody
}

type DescribeDBClustersWithBackupsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClustersWithBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClustersWithBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClustersWithBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClustersWithBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClustersWithBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClustersWithBackupsResponse) GetBody() *DescribeDBClustersWithBackupsResponseBody {
	return s.Body
}

func (s *DescribeDBClustersWithBackupsResponse) SetHeaders(v map[string]*string) *DescribeDBClustersWithBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetStatusCode(v int32) *DescribeDBClustersWithBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) SetBody(v *DescribeDBClustersWithBackupsResponseBody) *DescribeDBClustersWithBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClustersWithBackupsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

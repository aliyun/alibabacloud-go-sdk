// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClusterBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterBackupsResponseBody) *DescribeClusterBackupsResponse
	GetBody() *DescribeClusterBackupsResponseBody
}

type DescribeClusterBackupsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterBackupsResponse) GetBody() *DescribeClusterBackupsResponseBody {
	return s.Body
}

func (s *DescribeClusterBackupsResponse) SetHeaders(v map[string]*string) *DescribeClusterBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterBackupsResponse) SetStatusCode(v int32) *DescribeClusterBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterBackupsResponse) SetBody(v *DescribeClusterBackupsResponseBody) *DescribeClusterBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeClusterBackupsResponse) Validate() error {
	return dara.Validate(s)
}

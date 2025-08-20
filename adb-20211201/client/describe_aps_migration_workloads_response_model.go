// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApsMigrationWorkloadsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApsMigrationWorkloadsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApsMigrationWorkloadsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApsMigrationWorkloadsResponseBody) *DescribeApsMigrationWorkloadsResponse
	GetBody() *DescribeApsMigrationWorkloadsResponseBody
}

type DescribeApsMigrationWorkloadsResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApsMigrationWorkloadsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApsMigrationWorkloadsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApsMigrationWorkloadsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApsMigrationWorkloadsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApsMigrationWorkloadsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApsMigrationWorkloadsResponse) GetBody() *DescribeApsMigrationWorkloadsResponseBody {
	return s.Body
}

func (s *DescribeApsMigrationWorkloadsResponse) SetHeaders(v map[string]*string) *DescribeApsMigrationWorkloadsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponse) SetStatusCode(v int32) *DescribeApsMigrationWorkloadsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponse) SetBody(v *DescribeApsMigrationWorkloadsResponseBody) *DescribeApsMigrationWorkloadsResponse {
	s.Body = v
	return s
}

func (s *DescribeApsMigrationWorkloadsResponse) Validate() error {
	return dara.Validate(s)
}

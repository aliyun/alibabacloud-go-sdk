// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataBackupsResponseBody) *DescribeDataBackupsResponse
	GetBody() *DescribeDataBackupsResponseBody
}

type DescribeDataBackupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataBackupsResponse) GetBody() *DescribeDataBackupsResponseBody {
	return s.Body
}

func (s *DescribeDataBackupsResponse) SetHeaders(v map[string]*string) *DescribeDataBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataBackupsResponse) SetStatusCode(v int32) *DescribeDataBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataBackupsResponse) SetBody(v *DescribeDataBackupsResponseBody) *DescribeDataBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDataBackupsResponse) Validate() error {
	return dara.Validate(s)
}

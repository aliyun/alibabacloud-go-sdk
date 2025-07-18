// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogBackupsResponseBody) *DescribeLogBackupsResponse
	GetBody() *DescribeLogBackupsResponseBody
}

type DescribeLogBackupsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogBackupsResponse) GetBody() *DescribeLogBackupsResponseBody {
	return s.Body
}

func (s *DescribeLogBackupsResponse) SetHeaders(v map[string]*string) *DescribeLogBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogBackupsResponse) SetStatusCode(v int32) *DescribeLogBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogBackupsResponse) SetBody(v *DescribeLogBackupsResponseBody) *DescribeLogBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeLogBackupsResponse) Validate() error {
	return dara.Validate(s)
}

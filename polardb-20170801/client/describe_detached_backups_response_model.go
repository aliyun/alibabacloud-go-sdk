// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDetachedBackupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDetachedBackupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDetachedBackupsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDetachedBackupsResponseBody) *DescribeDetachedBackupsResponse
	GetBody() *DescribeDetachedBackupsResponseBody
}

type DescribeDetachedBackupsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDetachedBackupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDetachedBackupsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDetachedBackupsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDetachedBackupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDetachedBackupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDetachedBackupsResponse) GetBody() *DescribeDetachedBackupsResponseBody {
	return s.Body
}

func (s *DescribeDetachedBackupsResponse) SetHeaders(v map[string]*string) *DescribeDetachedBackupsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetStatusCode(v int32) *DescribeDetachedBackupsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDetachedBackupsResponse) SetBody(v *DescribeDetachedBackupsResponseBody) *DescribeDetachedBackupsResponse {
	s.Body = v
	return s
}

func (s *DescribeDetachedBackupsResponse) Validate() error {
	return dara.Validate(s)
}

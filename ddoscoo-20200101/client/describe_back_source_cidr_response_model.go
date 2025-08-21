// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBackSourceCidrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBackSourceCidrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBackSourceCidrResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBackSourceCidrResponseBody) *DescribeBackSourceCidrResponse
	GetBody() *DescribeBackSourceCidrResponseBody
}

type DescribeBackSourceCidrResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBackSourceCidrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBackSourceCidrResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBackSourceCidrResponse) GoString() string {
	return s.String()
}

func (s *DescribeBackSourceCidrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBackSourceCidrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBackSourceCidrResponse) GetBody() *DescribeBackSourceCidrResponseBody {
	return s.Body
}

func (s *DescribeBackSourceCidrResponse) SetHeaders(v map[string]*string) *DescribeBackSourceCidrResponse {
	s.Headers = v
	return s
}

func (s *DescribeBackSourceCidrResponse) SetStatusCode(v int32) *DescribeBackSourceCidrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBackSourceCidrResponse) SetBody(v *DescribeBackSourceCidrResponseBody) *DescribeBackSourceCidrResponse {
	s.Body = v
	return s
}

func (s *DescribeBackSourceCidrResponse) Validate() error {
	return dara.Validate(s)
}

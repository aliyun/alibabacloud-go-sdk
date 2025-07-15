// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagKeysResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagKeysResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagKeysResponseBody) *DescribeTagKeysResponse
	GetBody() *DescribeTagKeysResponseBody
}

type DescribeTagKeysResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagKeysResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagKeysResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagKeysResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagKeysResponse) GetBody() *DescribeTagKeysResponseBody {
	return s.Body
}

func (s *DescribeTagKeysResponse) SetHeaders(v map[string]*string) *DescribeTagKeysResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagKeysResponse) SetStatusCode(v int32) *DescribeTagKeysResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagKeysResponse) SetBody(v *DescribeTagKeysResponseBody) *DescribeTagKeysResponse {
	s.Body = v
	return s
}

func (s *DescribeTagKeysResponse) Validate() error {
	return dara.Validate(s)
}

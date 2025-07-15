// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppsByApiProductResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppsByApiProductResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppsByApiProductResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppsByApiProductResponseBody) *DescribeAppsByApiProductResponse
	GetBody() *DescribeAppsByApiProductResponseBody
}

type DescribeAppsByApiProductResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppsByApiProductResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppsByApiProductResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppsByApiProductResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppsByApiProductResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppsByApiProductResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppsByApiProductResponse) GetBody() *DescribeAppsByApiProductResponseBody {
	return s.Body
}

func (s *DescribeAppsByApiProductResponse) SetHeaders(v map[string]*string) *DescribeAppsByApiProductResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppsByApiProductResponse) SetStatusCode(v int32) *DescribeAppsByApiProductResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppsByApiProductResponse) SetBody(v *DescribeAppsByApiProductResponseBody) *DescribeAppsByApiProductResponse {
	s.Body = v
	return s
}

func (s *DescribeAppsByApiProductResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoDelConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAutoDelConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAutoDelConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAutoDelConfigResponseBody) *DescribeAutoDelConfigResponse
	GetBody() *DescribeAutoDelConfigResponseBody
}

type DescribeAutoDelConfigResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAutoDelConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAutoDelConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoDelConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeAutoDelConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAutoDelConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAutoDelConfigResponse) GetBody() *DescribeAutoDelConfigResponseBody {
	return s.Body
}

func (s *DescribeAutoDelConfigResponse) SetHeaders(v map[string]*string) *DescribeAutoDelConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeAutoDelConfigResponse) SetStatusCode(v int32) *DescribeAutoDelConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAutoDelConfigResponse) SetBody(v *DescribeAutoDelConfigResponseBody) *DescribeAutoDelConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeAutoDelConfigResponse) Validate() error {
	return dara.Validate(s)
}

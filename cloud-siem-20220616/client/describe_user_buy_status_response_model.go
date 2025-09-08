// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBuyStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUserBuyStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUserBuyStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUserBuyStatusResponseBody) *DescribeUserBuyStatusResponse
	GetBody() *DescribeUserBuyStatusResponseBody
}

type DescribeUserBuyStatusResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUserBuyStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUserBuyStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBuyStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeUserBuyStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUserBuyStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUserBuyStatusResponse) GetBody() *DescribeUserBuyStatusResponseBody {
	return s.Body
}

func (s *DescribeUserBuyStatusResponse) SetHeaders(v map[string]*string) *DescribeUserBuyStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeUserBuyStatusResponse) SetStatusCode(v int32) *DescribeUserBuyStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUserBuyStatusResponse) SetBody(v *DescribeUserBuyStatusResponseBody) *DescribeUserBuyStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeUserBuyStatusResponse) Validate() error {
	return dara.Validate(s)
}

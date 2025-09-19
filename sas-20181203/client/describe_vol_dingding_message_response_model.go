// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVolDingdingMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVolDingdingMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVolDingdingMessageResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVolDingdingMessageResponseBody) *DescribeVolDingdingMessageResponse
	GetBody() *DescribeVolDingdingMessageResponseBody
}

type DescribeVolDingdingMessageResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVolDingdingMessageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVolDingdingMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVolDingdingMessageResponse) GoString() string {
	return s.String()
}

func (s *DescribeVolDingdingMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVolDingdingMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVolDingdingMessageResponse) GetBody() *DescribeVolDingdingMessageResponseBody {
	return s.Body
}

func (s *DescribeVolDingdingMessageResponse) SetHeaders(v map[string]*string) *DescribeVolDingdingMessageResponse {
	s.Headers = v
	return s
}

func (s *DescribeVolDingdingMessageResponse) SetStatusCode(v int32) *DescribeVolDingdingMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVolDingdingMessageResponse) SetBody(v *DescribeVolDingdingMessageResponseBody) *DescribeVolDingdingMessageResponse {
	s.Body = v
	return s
}

func (s *DescribeVolDingdingMessageResponse) Validate() error {
	return dara.Validate(s)
}

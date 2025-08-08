// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProjectMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeProjectMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeProjectMessagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeProjectMessagesResponseBody) *DescribeProjectMessagesResponse
	GetBody() *DescribeProjectMessagesResponseBody
}

type DescribeProjectMessagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeProjectMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeProjectMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeProjectMessagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeProjectMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeProjectMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeProjectMessagesResponse) GetBody() *DescribeProjectMessagesResponseBody {
	return s.Body
}

func (s *DescribeProjectMessagesResponse) SetHeaders(v map[string]*string) *DescribeProjectMessagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeProjectMessagesResponse) SetStatusCode(v int32) *DescribeProjectMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeProjectMessagesResponse) SetBody(v *DescribeProjectMessagesResponseBody) *DescribeProjectMessagesResponse {
	s.Body = v
	return s
}

func (s *DescribeProjectMessagesResponse) Validate() error {
	return dara.Validate(s)
}

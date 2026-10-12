// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConsoleMessagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConsoleMessagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConsoleMessagesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConsoleMessagesResponseBody) *DescribeConsoleMessagesResponse
	GetBody() *DescribeConsoleMessagesResponseBody
}

type DescribeConsoleMessagesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConsoleMessagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConsoleMessagesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConsoleMessagesResponse) GoString() string {
	return s.String()
}

func (s *DescribeConsoleMessagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConsoleMessagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConsoleMessagesResponse) GetBody() *DescribeConsoleMessagesResponseBody {
	return s.Body
}

func (s *DescribeConsoleMessagesResponse) SetHeaders(v map[string]*string) *DescribeConsoleMessagesResponse {
	s.Headers = v
	return s
}

func (s *DescribeConsoleMessagesResponse) SetStatusCode(v int32) *DescribeConsoleMessagesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConsoleMessagesResponse) SetBody(v *DescribeConsoleMessagesResponseBody) *DescribeConsoleMessagesResponse {
	s.Body = v
	return s
}

func (s *DescribeConsoleMessagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

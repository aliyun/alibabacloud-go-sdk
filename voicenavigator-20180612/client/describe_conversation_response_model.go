// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConversationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConversationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConversationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConversationResponseBody) *DescribeConversationResponse
	GetBody() *DescribeConversationResponseBody
}

type DescribeConversationResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConversationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConversationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConversationResponse) GoString() string {
	return s.String()
}

func (s *DescribeConversationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConversationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConversationResponse) GetBody() *DescribeConversationResponseBody {
	return s.Body
}

func (s *DescribeConversationResponse) SetHeaders(v map[string]*string) *DescribeConversationResponse {
	s.Headers = v
	return s
}

func (s *DescribeConversationResponse) SetStatusCode(v int32) *DescribeConversationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConversationResponse) SetBody(v *DescribeConversationResponseBody) *DescribeConversationResponse {
	s.Body = v
	return s
}

func (s *DescribeConversationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

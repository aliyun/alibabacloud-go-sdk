// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeConversationContextResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeConversationContextResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeConversationContextResponse
	GetStatusCode() *int32
	SetBody(v *DescribeConversationContextResponseBody) *DescribeConversationContextResponse
	GetBody() *DescribeConversationContextResponseBody
}

type DescribeConversationContextResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeConversationContextResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeConversationContextResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeConversationContextResponse) GoString() string {
	return s.String()
}

func (s *DescribeConversationContextResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeConversationContextResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeConversationContextResponse) GetBody() *DescribeConversationContextResponseBody {
	return s.Body
}

func (s *DescribeConversationContextResponse) SetHeaders(v map[string]*string) *DescribeConversationContextResponse {
	s.Headers = v
	return s
}

func (s *DescribeConversationContextResponse) SetStatusCode(v int32) *DescribeConversationContextResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeConversationContextResponse) SetBody(v *DescribeConversationContextResponseBody) *DescribeConversationContextResponse {
	s.Body = v
	return s
}

func (s *DescribeConversationContextResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

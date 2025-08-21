// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIntentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIntentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIntentResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIntentResponseBody) *DescribeIntentResponse
	GetBody() *DescribeIntentResponseBody
}

type DescribeIntentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIntentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIntentResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIntentResponse) GoString() string {
	return s.String()
}

func (s *DescribeIntentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIntentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIntentResponse) GetBody() *DescribeIntentResponseBody {
	return s.Body
}

func (s *DescribeIntentResponse) SetHeaders(v map[string]*string) *DescribeIntentResponse {
	s.Headers = v
	return s
}

func (s *DescribeIntentResponse) SetStatusCode(v int32) *DescribeIntentResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIntentResponse) SetBody(v *DescribeIntentResponseBody) *DescribeIntentResponse {
	s.Body = v
	return s
}

func (s *DescribeIntentResponse) Validate() error {
	return dara.Validate(s)
}

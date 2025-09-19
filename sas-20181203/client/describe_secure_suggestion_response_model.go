// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecureSuggestionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSecureSuggestionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSecureSuggestionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSecureSuggestionResponseBody) *DescribeSecureSuggestionResponse
	GetBody() *DescribeSecureSuggestionResponseBody
}

type DescribeSecureSuggestionResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSecureSuggestionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSecureSuggestionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecureSuggestionResponse) GoString() string {
	return s.String()
}

func (s *DescribeSecureSuggestionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSecureSuggestionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSecureSuggestionResponse) GetBody() *DescribeSecureSuggestionResponseBody {
	return s.Body
}

func (s *DescribeSecureSuggestionResponse) SetHeaders(v map[string]*string) *DescribeSecureSuggestionResponse {
	s.Headers = v
	return s
}

func (s *DescribeSecureSuggestionResponse) SetStatusCode(v int32) *DescribeSecureSuggestionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSecureSuggestionResponse) SetBody(v *DescribeSecureSuggestionResponseBody) *DescribeSecureSuggestionResponse {
	s.Body = v
	return s
}

func (s *DescribeSecureSuggestionResponse) Validate() error {
	return dara.Validate(s)
}

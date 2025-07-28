// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSuggestionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecSuggestionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecSuggestionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecSuggestionsResponseBody) *DescribeApisecSuggestionsResponse
	GetBody() *DescribeApisecSuggestionsResponseBody
}

type DescribeApisecSuggestionsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecSuggestionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecSuggestionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSuggestionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecSuggestionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecSuggestionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecSuggestionsResponse) GetBody() *DescribeApisecSuggestionsResponseBody {
	return s.Body
}

func (s *DescribeApisecSuggestionsResponse) SetHeaders(v map[string]*string) *DescribeApisecSuggestionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecSuggestionsResponse) SetStatusCode(v int32) *DescribeApisecSuggestionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecSuggestionsResponse) SetBody(v *DescribeApisecSuggestionsResponseBody) *DescribeApisecSuggestionsResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecSuggestionsResponse) Validate() error {
	return dara.Validate(s)
}

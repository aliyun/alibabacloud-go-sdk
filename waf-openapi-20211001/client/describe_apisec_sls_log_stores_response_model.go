// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApisecSlsLogStoresResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApisecSlsLogStoresResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApisecSlsLogStoresResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApisecSlsLogStoresResponseBody) *DescribeApisecSlsLogStoresResponse
	GetBody() *DescribeApisecSlsLogStoresResponseBody
}

type DescribeApisecSlsLogStoresResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApisecSlsLogStoresResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApisecSlsLogStoresResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApisecSlsLogStoresResponse) GoString() string {
	return s.String()
}

func (s *DescribeApisecSlsLogStoresResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApisecSlsLogStoresResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApisecSlsLogStoresResponse) GetBody() *DescribeApisecSlsLogStoresResponseBody {
	return s.Body
}

func (s *DescribeApisecSlsLogStoresResponse) SetHeaders(v map[string]*string) *DescribeApisecSlsLogStoresResponse {
	s.Headers = v
	return s
}

func (s *DescribeApisecSlsLogStoresResponse) SetStatusCode(v int32) *DescribeApisecSlsLogStoresResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApisecSlsLogStoresResponse) SetBody(v *DescribeApisecSlsLogStoresResponseBody) *DescribeApisecSlsLogStoresResponse {
	s.Body = v
	return s
}

func (s *DescribeApisecSlsLogStoresResponse) Validate() error {
	return dara.Validate(s)
}

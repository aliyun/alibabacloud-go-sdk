// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCrossAccountsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCrossAccountsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCrossAccountsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCrossAccountsResponseBody) *DescribeCrossAccountsResponse
	GetBody() *DescribeCrossAccountsResponseBody
}

type DescribeCrossAccountsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCrossAccountsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCrossAccountsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCrossAccountsResponse) GoString() string {
	return s.String()
}

func (s *DescribeCrossAccountsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCrossAccountsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCrossAccountsResponse) GetBody() *DescribeCrossAccountsResponseBody {
	return s.Body
}

func (s *DescribeCrossAccountsResponse) SetHeaders(v map[string]*string) *DescribeCrossAccountsResponse {
	s.Headers = v
	return s
}

func (s *DescribeCrossAccountsResponse) SetStatusCode(v int32) *DescribeCrossAccountsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCrossAccountsResponse) SetBody(v *DescribeCrossAccountsResponseBody) *DescribeCrossAccountsResponse {
	s.Body = v
	return s
}

func (s *DescribeCrossAccountsResponse) Validate() error {
	return dara.Validate(s)
}

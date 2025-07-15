// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiHistoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeApiHistoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeApiHistoriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeApiHistoriesResponseBody) *DescribeApiHistoriesResponse
	GetBody() *DescribeApiHistoriesResponseBody
}

type DescribeApiHistoriesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeApiHistoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeApiHistoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiHistoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeApiHistoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeApiHistoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeApiHistoriesResponse) GetBody() *DescribeApiHistoriesResponseBody {
	return s.Body
}

func (s *DescribeApiHistoriesResponse) SetHeaders(v map[string]*string) *DescribeApiHistoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeApiHistoriesResponse) SetStatusCode(v int32) *DescribeApiHistoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeApiHistoriesResponse) SetBody(v *DescribeApiHistoriesResponseBody) *DescribeApiHistoriesResponse {
	s.Body = v
	return s
}

func (s *DescribeApiHistoriesResponse) Validate() error {
	return dara.Validate(s)
}

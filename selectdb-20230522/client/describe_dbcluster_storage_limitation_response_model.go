// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterStorageLimitationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBClusterStorageLimitationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBClusterStorageLimitationResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBClusterStorageLimitationResponseBody) *DescribeDBClusterStorageLimitationResponse
	GetBody() *DescribeDBClusterStorageLimitationResponseBody
}

type DescribeDBClusterStorageLimitationResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBClusterStorageLimitationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBClusterStorageLimitationResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterStorageLimitationResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterStorageLimitationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBClusterStorageLimitationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBClusterStorageLimitationResponse) GetBody() *DescribeDBClusterStorageLimitationResponseBody {
	return s.Body
}

func (s *DescribeDBClusterStorageLimitationResponse) SetHeaders(v map[string]*string) *DescribeDBClusterStorageLimitationResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponse) SetStatusCode(v int32) *DescribeDBClusterStorageLimitationResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponse) SetBody(v *DescribeDBClusterStorageLimitationResponseBody) *DescribeDBClusterStorageLimitationResponse {
	s.Body = v
	return s
}

func (s *DescribeDBClusterStorageLimitationResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCollectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCollectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCollectionResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCollectionResponseBody) *DescribeCollectionResponse
	GetBody() *DescribeCollectionResponseBody
}

type DescribeCollectionResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCollectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCollectionResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCollectionResponse) GoString() string {
	return s.String()
}

func (s *DescribeCollectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCollectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCollectionResponse) GetBody() *DescribeCollectionResponseBody {
	return s.Body
}

func (s *DescribeCollectionResponse) SetHeaders(v map[string]*string) *DescribeCollectionResponse {
	s.Headers = v
	return s
}

func (s *DescribeCollectionResponse) SetStatusCode(v int32) *DescribeCollectionResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCollectionResponse) SetBody(v *DescribeCollectionResponseBody) *DescribeCollectionResponse {
	s.Body = v
	return s
}

func (s *DescribeCollectionResponse) Validate() error {
	return dara.Validate(s)
}

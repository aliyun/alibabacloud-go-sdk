// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEnumItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEnumItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEnumItemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEnumItemsResponseBody) *DescribeEnumItemsResponse
	GetBody() *DescribeEnumItemsResponseBody
}

type DescribeEnumItemsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEnumItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEnumItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEnumItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeEnumItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEnumItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEnumItemsResponse) GetBody() *DescribeEnumItemsResponseBody {
	return s.Body
}

func (s *DescribeEnumItemsResponse) SetHeaders(v map[string]*string) *DescribeEnumItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeEnumItemsResponse) SetStatusCode(v int32) *DescribeEnumItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEnumItemsResponse) SetBody(v *DescribeEnumItemsResponseBody) *DescribeEnumItemsResponse {
	s.Body = v
	return s
}

func (s *DescribeEnumItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

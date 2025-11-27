// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceByTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDBInstanceByTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDBInstanceByTagsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDBInstanceByTagsResponseBody) *DescribeDBInstanceByTagsResponse
	GetBody() *DescribeDBInstanceByTagsResponseBody
}

type DescribeDBInstanceByTagsResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDBInstanceByTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDBInstanceByTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceByTagsResponse) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceByTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDBInstanceByTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDBInstanceByTagsResponse) GetBody() *DescribeDBInstanceByTagsResponseBody {
	return s.Body
}

func (s *DescribeDBInstanceByTagsResponse) SetHeaders(v map[string]*string) *DescribeDBInstanceByTagsResponse {
	s.Headers = v
	return s
}

func (s *DescribeDBInstanceByTagsResponse) SetStatusCode(v int32) *DescribeDBInstanceByTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDBInstanceByTagsResponse) SetBody(v *DescribeDBInstanceByTagsResponseBody) *DescribeDBInstanceByTagsResponse {
	s.Body = v
	return s
}

func (s *DescribeDBInstanceByTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

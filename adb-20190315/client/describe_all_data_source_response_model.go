// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAllDataSourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAllDataSourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAllDataSourceResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAllDataSourceResponseBody) *DescribeAllDataSourceResponse
	GetBody() *DescribeAllDataSourceResponseBody
}

type DescribeAllDataSourceResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAllDataSourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAllDataSourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAllDataSourceResponse) GoString() string {
	return s.String()
}

func (s *DescribeAllDataSourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAllDataSourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAllDataSourceResponse) GetBody() *DescribeAllDataSourceResponseBody {
	return s.Body
}

func (s *DescribeAllDataSourceResponse) SetHeaders(v map[string]*string) *DescribeAllDataSourceResponse {
	s.Headers = v
	return s
}

func (s *DescribeAllDataSourceResponse) SetStatusCode(v int32) *DescribeAllDataSourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAllDataSourceResponse) SetBody(v *DescribeAllDataSourceResponseBody) *DescribeAllDataSourceResponse {
	s.Body = v
	return s
}

func (s *DescribeAllDataSourceResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

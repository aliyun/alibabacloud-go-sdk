// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorGroupCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorGroupCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorGroupCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorGroupCategoriesResponseBody) *DescribeMonitorGroupCategoriesResponse
	GetBody() *DescribeMonitorGroupCategoriesResponseBody
}

type DescribeMonitorGroupCategoriesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorGroupCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorGroupCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorGroupCategoriesResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorGroupCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorGroupCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorGroupCategoriesResponse) GetBody() *DescribeMonitorGroupCategoriesResponseBody {
	return s.Body
}

func (s *DescribeMonitorGroupCategoriesResponse) SetHeaders(v map[string]*string) *DescribeMonitorGroupCategoriesResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponse) SetStatusCode(v int32) *DescribeMonitorGroupCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponse) SetBody(v *DescribeMonitorGroupCategoriesResponseBody) *DescribeMonitorGroupCategoriesResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorGroupCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

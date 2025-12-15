// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMonitorItemsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMonitorItemsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMonitorItemsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMonitorItemsResponseBody) *DescribeMonitorItemsResponse
	GetBody() *DescribeMonitorItemsResponseBody
}

type DescribeMonitorItemsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMonitorItemsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMonitorItemsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMonitorItemsResponse) GoString() string {
	return s.String()
}

func (s *DescribeMonitorItemsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMonitorItemsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMonitorItemsResponse) GetBody() *DescribeMonitorItemsResponseBody {
	return s.Body
}

func (s *DescribeMonitorItemsResponse) SetHeaders(v map[string]*string) *DescribeMonitorItemsResponse {
	s.Headers = v
	return s
}

func (s *DescribeMonitorItemsResponse) SetStatusCode(v int32) *DescribeMonitorItemsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMonitorItemsResponse) SetBody(v *DescribeMonitorItemsResponseBody) *DescribeMonitorItemsResponse {
	s.Body = v
	return s
}

func (s *DescribeMonitorItemsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEndPointEventListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEndPointEventListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEndPointEventListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEndPointEventListResponseBody) *DescribeEndPointEventListResponse
	GetBody() *DescribeEndPointEventListResponseBody
}

type DescribeEndPointEventListResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEndPointEventListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEndPointEventListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEndPointEventListResponse) GoString() string {
	return s.String()
}

func (s *DescribeEndPointEventListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEndPointEventListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEndPointEventListResponse) GetBody() *DescribeEndPointEventListResponseBody {
	return s.Body
}

func (s *DescribeEndPointEventListResponse) SetHeaders(v map[string]*string) *DescribeEndPointEventListResponse {
	s.Headers = v
	return s
}

func (s *DescribeEndPointEventListResponse) SetStatusCode(v int32) *DescribeEndPointEventListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEndPointEventListResponse) SetBody(v *DescribeEndPointEventListResponseBody) *DescribeEndPointEventListResponse {
	s.Body = v
	return s
}

func (s *DescribeEndPointEventListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

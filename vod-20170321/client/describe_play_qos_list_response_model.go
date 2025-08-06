// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePlayQosListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePlayQosListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePlayQosListResponse
	GetStatusCode() *int32
	SetBody(v *DescribePlayQosListResponseBody) *DescribePlayQosListResponse
	GetBody() *DescribePlayQosListResponseBody
}

type DescribePlayQosListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePlayQosListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePlayQosListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePlayQosListResponse) GoString() string {
	return s.String()
}

func (s *DescribePlayQosListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePlayQosListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePlayQosListResponse) GetBody() *DescribePlayQosListResponseBody {
	return s.Body
}

func (s *DescribePlayQosListResponse) SetHeaders(v map[string]*string) *DescribePlayQosListResponse {
	s.Headers = v
	return s
}

func (s *DescribePlayQosListResponse) SetStatusCode(v int32) *DescribePlayQosListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePlayQosListResponse) SetBody(v *DescribePlayQosListResponseBody) *DescribePlayQosListResponse {
	s.Body = v
	return s
}

func (s *DescribePlayQosListResponse) Validate() error {
	return dara.Validate(s)
}

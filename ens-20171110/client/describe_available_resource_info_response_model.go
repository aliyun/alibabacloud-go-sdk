// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAvailableResourceInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAvailableResourceInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAvailableResourceInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAvailableResourceInfoResponseBody) *DescribeAvailableResourceInfoResponse
	GetBody() *DescribeAvailableResourceInfoResponseBody
}

type DescribeAvailableResourceInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAvailableResourceInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAvailableResourceInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAvailableResourceInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeAvailableResourceInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAvailableResourceInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAvailableResourceInfoResponse) GetBody() *DescribeAvailableResourceInfoResponseBody {
	return s.Body
}

func (s *DescribeAvailableResourceInfoResponse) SetHeaders(v map[string]*string) *DescribeAvailableResourceInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeAvailableResourceInfoResponse) SetStatusCode(v int32) *DescribeAvailableResourceInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAvailableResourceInfoResponse) SetBody(v *DescribeAvailableResourceInfoResponseBody) *DescribeAvailableResourceInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeAvailableResourceInfoResponse) Validate() error {
	return dara.Validate(s)
}

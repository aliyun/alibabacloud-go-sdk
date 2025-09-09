// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHotDbListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHotDbListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHotDbListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHotDbListResponseBody) *DescribeHotDbListResponse
	GetBody() *DescribeHotDbListResponseBody
}

type DescribeHotDbListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHotDbListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHotDbListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHotDbListResponse) GoString() string {
	return s.String()
}

func (s *DescribeHotDbListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHotDbListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHotDbListResponse) GetBody() *DescribeHotDbListResponseBody {
	return s.Body
}

func (s *DescribeHotDbListResponse) SetHeaders(v map[string]*string) *DescribeHotDbListResponse {
	s.Headers = v
	return s
}

func (s *DescribeHotDbListResponse) SetStatusCode(v int32) *DescribeHotDbListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHotDbListResponse) SetBody(v *DescribeHotDbListResponseBody) *DescribeHotDbListResponse {
	s.Body = v
	return s
}

func (s *DescribeHotDbListResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlsEmptyCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlsEmptyCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlsEmptyCountResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlsEmptyCountResponseBody) *DescribeSlsEmptyCountResponse
	GetBody() *DescribeSlsEmptyCountResponseBody
}

type DescribeSlsEmptyCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlsEmptyCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlsEmptyCountResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlsEmptyCountResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlsEmptyCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlsEmptyCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlsEmptyCountResponse) GetBody() *DescribeSlsEmptyCountResponseBody {
	return s.Body
}

func (s *DescribeSlsEmptyCountResponse) SetHeaders(v map[string]*string) *DescribeSlsEmptyCountResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlsEmptyCountResponse) SetStatusCode(v int32) *DescribeSlsEmptyCountResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlsEmptyCountResponse) SetBody(v *DescribeSlsEmptyCountResponseBody) *DescribeSlsEmptyCountResponse {
	s.Body = v
	return s
}

func (s *DescribeSlsEmptyCountResponse) Validate() error {
	return dara.Validate(s)
}

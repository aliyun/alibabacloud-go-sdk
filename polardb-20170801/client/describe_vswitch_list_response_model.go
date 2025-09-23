// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVSwitchListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVSwitchListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVSwitchListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVSwitchListResponseBody) *DescribeVSwitchListResponse
	GetBody() *DescribeVSwitchListResponseBody
}

type DescribeVSwitchListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVSwitchListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVSwitchListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVSwitchListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVSwitchListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVSwitchListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVSwitchListResponse) GetBody() *DescribeVSwitchListResponseBody {
	return s.Body
}

func (s *DescribeVSwitchListResponse) SetHeaders(v map[string]*string) *DescribeVSwitchListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVSwitchListResponse) SetStatusCode(v int32) *DescribeVSwitchListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVSwitchListResponse) SetBody(v *DescribeVSwitchListResponseBody) *DescribeVSwitchListResponse {
	s.Body = v
	return s
}

func (s *DescribeVSwitchListResponse) Validate() error {
	return dara.Validate(s)
}

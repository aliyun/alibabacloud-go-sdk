// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdvanceSearchPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAdvanceSearchPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAdvanceSearchPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAdvanceSearchPageListResponseBody) *DescribeAdvanceSearchPageListResponse
	GetBody() *DescribeAdvanceSearchPageListResponseBody
}

type DescribeAdvanceSearchPageListResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAdvanceSearchPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAdvanceSearchPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdvanceSearchPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAdvanceSearchPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAdvanceSearchPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAdvanceSearchPageListResponse) GetBody() *DescribeAdvanceSearchPageListResponseBody {
	return s.Body
}

func (s *DescribeAdvanceSearchPageListResponse) SetHeaders(v map[string]*string) *DescribeAdvanceSearchPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAdvanceSearchPageListResponse) SetStatusCode(v int32) *DescribeAdvanceSearchPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAdvanceSearchPageListResponse) SetBody(v *DescribeAdvanceSearchPageListResponseBody) *DescribeAdvanceSearchPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeAdvanceSearchPageListResponse) Validate() error {
	return dara.Validate(s)
}

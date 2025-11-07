// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySearchPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVerifySearchPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVerifySearchPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVerifySearchPageListResponseBody) *DescribeVerifySearchPageListResponse
	GetBody() *DescribeVerifySearchPageListResponseBody
}

type DescribeVerifySearchPageListResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVerifySearchPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVerifySearchPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySearchPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeVerifySearchPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVerifySearchPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVerifySearchPageListResponse) GetBody() *DescribeVerifySearchPageListResponseBody {
	return s.Body
}

func (s *DescribeVerifySearchPageListResponse) SetHeaders(v map[string]*string) *DescribeVerifySearchPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeVerifySearchPageListResponse) SetStatusCode(v int32) *DescribeVerifySearchPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVerifySearchPageListResponse) SetBody(v *DescribeVerifySearchPageListResponseBody) *DescribeVerifySearchPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeVerifySearchPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

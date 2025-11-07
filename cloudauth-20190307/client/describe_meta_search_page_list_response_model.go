// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMetaSearchPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeMetaSearchPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeMetaSearchPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeMetaSearchPageListResponseBody) *DescribeMetaSearchPageListResponse
	GetBody() *DescribeMetaSearchPageListResponseBody
}

type DescribeMetaSearchPageListResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeMetaSearchPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeMetaSearchPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeMetaSearchPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeMetaSearchPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeMetaSearchPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeMetaSearchPageListResponse) GetBody() *DescribeMetaSearchPageListResponseBody {
	return s.Body
}

func (s *DescribeMetaSearchPageListResponse) SetHeaders(v map[string]*string) *DescribeMetaSearchPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeMetaSearchPageListResponse) SetStatusCode(v int32) *DescribeMetaSearchPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeMetaSearchPageListResponse) SetBody(v *DescribeMetaSearchPageListResponseBody) *DescribeMetaSearchPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeMetaSearchPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

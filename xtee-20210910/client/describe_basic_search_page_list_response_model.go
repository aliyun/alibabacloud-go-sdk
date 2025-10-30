// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBasicSearchPageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeBasicSearchPageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeBasicSearchPageListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeBasicSearchPageListResponseBody) *DescribeBasicSearchPageListResponse
	GetBody() *DescribeBasicSearchPageListResponseBody
}

type DescribeBasicSearchPageListResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeBasicSearchPageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeBasicSearchPageListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeBasicSearchPageListResponse) GoString() string {
	return s.String()
}

func (s *DescribeBasicSearchPageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeBasicSearchPageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeBasicSearchPageListResponse) GetBody() *DescribeBasicSearchPageListResponseBody {
	return s.Body
}

func (s *DescribeBasicSearchPageListResponse) SetHeaders(v map[string]*string) *DescribeBasicSearchPageListResponse {
	s.Headers = v
	return s
}

func (s *DescribeBasicSearchPageListResponse) SetStatusCode(v int32) *DescribeBasicSearchPageListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeBasicSearchPageListResponse) SetBody(v *DescribeBasicSearchPageListResponseBody) *DescribeBasicSearchPageListResponse {
	s.Body = v
	return s
}

func (s *DescribeBasicSearchPageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

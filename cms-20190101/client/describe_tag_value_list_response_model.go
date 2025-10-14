// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagValueListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTagValueListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTagValueListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTagValueListResponseBody) *DescribeTagValueListResponse
	GetBody() *DescribeTagValueListResponseBody
}

type DescribeTagValueListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTagValueListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTagValueListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagValueListResponse) GoString() string {
	return s.String()
}

func (s *DescribeTagValueListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTagValueListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTagValueListResponse) GetBody() *DescribeTagValueListResponseBody {
	return s.Body
}

func (s *DescribeTagValueListResponse) SetHeaders(v map[string]*string) *DescribeTagValueListResponse {
	s.Headers = v
	return s
}

func (s *DescribeTagValueListResponse) SetStatusCode(v int32) *DescribeTagValueListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTagValueListResponse) SetBody(v *DescribeTagValueListResponseBody) *DescribeTagValueListResponse {
	s.Body = v
	return s
}

func (s *DescribeTagValueListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

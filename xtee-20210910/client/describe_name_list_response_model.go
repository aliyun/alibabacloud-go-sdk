// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNameListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNameListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNameListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNameListResponseBody) *DescribeNameListResponse
	GetBody() *DescribeNameListResponseBody
}

type DescribeNameListResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNameListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNameListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNameListResponse) GoString() string {
	return s.String()
}

func (s *DescribeNameListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNameListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNameListResponse) GetBody() *DescribeNameListResponseBody {
	return s.Body
}

func (s *DescribeNameListResponse) SetHeaders(v map[string]*string) *DescribeNameListResponse {
	s.Headers = v
	return s
}

func (s *DescribeNameListResponse) SetStatusCode(v int32) *DescribeNameListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNameListResponse) SetBody(v *DescribeNameListResponseBody) *DescribeNameListResponse {
	s.Body = v
	return s
}

func (s *DescribeNameListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

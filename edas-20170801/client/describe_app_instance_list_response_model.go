// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppInstanceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppInstanceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppInstanceListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppInstanceListResponseBody) *DescribeAppInstanceListResponse
	GetBody() *DescribeAppInstanceListResponseBody
}

type DescribeAppInstanceListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppInstanceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppInstanceListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppInstanceListResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppInstanceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppInstanceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppInstanceListResponse) GetBody() *DescribeAppInstanceListResponseBody {
	return s.Body
}

func (s *DescribeAppInstanceListResponse) SetHeaders(v map[string]*string) *DescribeAppInstanceListResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppInstanceListResponse) SetStatusCode(v int32) *DescribeAppInstanceListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppInstanceListResponse) SetBody(v *DescribeAppInstanceListResponseBody) *DescribeAppInstanceListResponse {
	s.Body = v
	return s
}

func (s *DescribeAppInstanceListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

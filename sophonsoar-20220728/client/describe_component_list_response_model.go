// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeComponentListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeComponentListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeComponentListResponse
	GetStatusCode() *int32
	SetBody(v *DescribeComponentListResponseBody) *DescribeComponentListResponse
	GetBody() *DescribeComponentListResponseBody
}

type DescribeComponentListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeComponentListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeComponentListResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeComponentListResponse) GoString() string {
	return s.String()
}

func (s *DescribeComponentListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeComponentListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeComponentListResponse) GetBody() *DescribeComponentListResponseBody {
	return s.Body
}

func (s *DescribeComponentListResponse) SetHeaders(v map[string]*string) *DescribeComponentListResponse {
	s.Headers = v
	return s
}

func (s *DescribeComponentListResponse) SetStatusCode(v int32) *DescribeComponentListResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeComponentListResponse) SetBody(v *DescribeComponentListResponseBody) *DescribeComponentListResponse {
	s.Body = v
	return s
}

func (s *DescribeComponentListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlrResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlrResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlrResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlrResponseBody) *DescribeSlrResponse
	GetBody() *DescribeSlrResponseBody
}

type DescribeSlrResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlrResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlrResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlrResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlrResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlrResponse) GetBody() *DescribeSlrResponseBody {
	return s.Body
}

func (s *DescribeSlrResponse) SetHeaders(v map[string]*string) *DescribeSlrResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlrResponse) SetStatusCode(v int32) *DescribeSlrResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlrResponse) SetBody(v *DescribeSlrResponseBody) *DescribeSlrResponse {
	s.Body = v
	return s
}

func (s *DescribeSlrResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSlrGrantResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSlrGrantResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSlrGrantResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSlrGrantResponseBody) *DescribeSlrGrantResponse
	GetBody() *DescribeSlrGrantResponseBody
}

type DescribeSlrGrantResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSlrGrantResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSlrGrantResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSlrGrantResponse) GoString() string {
	return s.String()
}

func (s *DescribeSlrGrantResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSlrGrantResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSlrGrantResponse) GetBody() *DescribeSlrGrantResponseBody {
	return s.Body
}

func (s *DescribeSlrGrantResponse) SetHeaders(v map[string]*string) *DescribeSlrGrantResponse {
	s.Headers = v
	return s
}

func (s *DescribeSlrGrantResponse) SetStatusCode(v int32) *DescribeSlrGrantResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSlrGrantResponse) SetBody(v *DescribeSlrGrantResponseBody) *DescribeSlrGrantResponse {
	s.Body = v
	return s
}

func (s *DescribeSlrGrantResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

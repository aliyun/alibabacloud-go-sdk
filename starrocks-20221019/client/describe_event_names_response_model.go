// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventNamesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventNamesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventNamesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventNamesResponseBody) *DescribeEventNamesResponse
	GetBody() *DescribeEventNamesResponseBody
}

type DescribeEventNamesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventNamesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventNamesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventNamesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventNamesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventNamesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventNamesResponse) GetBody() *DescribeEventNamesResponseBody {
	return s.Body
}

func (s *DescribeEventNamesResponse) SetHeaders(v map[string]*string) *DescribeEventNamesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventNamesResponse) SetStatusCode(v int32) *DescribeEventNamesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventNamesResponse) SetBody(v *DescribeEventNamesResponseBody) *DescribeEventNamesResponse {
	s.Body = v
	return s
}

func (s *DescribeEventNamesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

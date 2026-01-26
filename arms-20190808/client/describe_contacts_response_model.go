// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeContactsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeContactsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeContactsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeContactsResponseBody) *DescribeContactsResponse
	GetBody() *DescribeContactsResponseBody
}

type DescribeContactsResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeContactsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeContactsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeContactsResponse) GoString() string {
	return s.String()
}

func (s *DescribeContactsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeContactsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeContactsResponse) GetBody() *DescribeContactsResponseBody {
	return s.Body
}

func (s *DescribeContactsResponse) SetHeaders(v map[string]*string) *DescribeContactsResponse {
	s.Headers = v
	return s
}

func (s *DescribeContactsResponse) SetStatusCode(v int32) *DescribeContactsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeContactsResponse) SetBody(v *DescribeContactsResponseBody) *DescribeContactsResponse {
	s.Body = v
	return s
}

func (s *DescribeContactsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

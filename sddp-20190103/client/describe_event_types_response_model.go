// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEventTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEventTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEventTypesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEventTypesResponseBody) *DescribeEventTypesResponse
	GetBody() *DescribeEventTypesResponseBody
}

type DescribeEventTypesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEventTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEventTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEventTypesResponse) GoString() string {
	return s.String()
}

func (s *DescribeEventTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEventTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEventTypesResponse) GetBody() *DescribeEventTypesResponseBody {
	return s.Body
}

func (s *DescribeEventTypesResponse) SetHeaders(v map[string]*string) *DescribeEventTypesResponse {
	s.Headers = v
	return s
}

func (s *DescribeEventTypesResponse) SetStatusCode(v int32) *DescribeEventTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEventTypesResponse) SetBody(v *DescribeEventTypesResponseBody) *DescribeEventTypesResponse {
	s.Body = v
	return s
}

func (s *DescribeEventTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

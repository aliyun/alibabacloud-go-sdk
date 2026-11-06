// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSearchParametersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSearchParametersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSearchParametersResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSearchParametersResponseBody) *DescribeSearchParametersResponse
	GetBody() *DescribeSearchParametersResponseBody
}

type DescribeSearchParametersResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSearchParametersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSearchParametersResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSearchParametersResponse) GoString() string {
	return s.String()
}

func (s *DescribeSearchParametersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSearchParametersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSearchParametersResponse) GetBody() *DescribeSearchParametersResponseBody {
	return s.Body
}

func (s *DescribeSearchParametersResponse) SetHeaders(v map[string]*string) *DescribeSearchParametersResponse {
	s.Headers = v
	return s
}

func (s *DescribeSearchParametersResponse) SetStatusCode(v int32) *DescribeSearchParametersResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSearchParametersResponse) SetBody(v *DescribeSearchParametersResponseBody) *DescribeSearchParametersResponse {
	s.Body = v
	return s
}

func (s *DescribeSearchParametersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

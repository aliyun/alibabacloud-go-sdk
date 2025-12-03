// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEaisResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEaisResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEaisResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEaisResponseBody) *DescribeEaisResponse
	GetBody() *DescribeEaisResponseBody
}

type DescribeEaisResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEaisResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEaisResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEaisResponse) GoString() string {
	return s.String()
}

func (s *DescribeEaisResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEaisResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEaisResponse) GetBody() *DescribeEaisResponseBody {
	return s.Body
}

func (s *DescribeEaisResponse) SetHeaders(v map[string]*string) *DescribeEaisResponse {
	s.Headers = v
	return s
}

func (s *DescribeEaisResponse) SetStatusCode(v int32) *DescribeEaisResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEaisResponse) SetBody(v *DescribeEaisResponseBody) *DescribeEaisResponse {
	s.Body = v
	return s
}

func (s *DescribeEaisResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

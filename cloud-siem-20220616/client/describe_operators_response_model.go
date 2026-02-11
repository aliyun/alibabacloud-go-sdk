// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperatorsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperatorsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperatorsResponseBody) *DescribeOperatorsResponse
	GetBody() *DescribeOperatorsResponseBody
}

type DescribeOperatorsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperatorsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperatorsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorsResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperatorsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperatorsResponse) GetBody() *DescribeOperatorsResponseBody {
	return s.Body
}

func (s *DescribeOperatorsResponse) SetHeaders(v map[string]*string) *DescribeOperatorsResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorsResponse) SetStatusCode(v int32) *DescribeOperatorsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorsResponse) SetBody(v *DescribeOperatorsResponseBody) *DescribeOperatorsResponse {
	s.Body = v
	return s
}

func (s *DescribeOperatorsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

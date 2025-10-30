// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOperatorListByTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOperatorListByTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOperatorListByTypeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOperatorListByTypeResponseBody) *DescribeOperatorListByTypeResponse
	GetBody() *DescribeOperatorListByTypeResponseBody
}

type DescribeOperatorListByTypeResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOperatorListByTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOperatorListByTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOperatorListByTypeResponse) GoString() string {
	return s.String()
}

func (s *DescribeOperatorListByTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOperatorListByTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOperatorListByTypeResponse) GetBody() *DescribeOperatorListByTypeResponseBody {
	return s.Body
}

func (s *DescribeOperatorListByTypeResponse) SetHeaders(v map[string]*string) *DescribeOperatorListByTypeResponse {
	s.Headers = v
	return s
}

func (s *DescribeOperatorListByTypeResponse) SetStatusCode(v int32) *DescribeOperatorListByTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOperatorListByTypeResponse) SetBody(v *DescribeOperatorListByTypeResponseBody) *DescribeOperatorListByTypeResponse {
	s.Body = v
	return s
}

func (s *DescribeOperatorListByTypeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

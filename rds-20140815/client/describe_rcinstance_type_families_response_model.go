// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRCInstanceTypeFamiliesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRCInstanceTypeFamiliesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRCInstanceTypeFamiliesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRCInstanceTypeFamiliesResponseBody) *DescribeRCInstanceTypeFamiliesResponse
	GetBody() *DescribeRCInstanceTypeFamiliesResponseBody
}

type DescribeRCInstanceTypeFamiliesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRCInstanceTypeFamiliesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRCInstanceTypeFamiliesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRCInstanceTypeFamiliesResponse) GoString() string {
	return s.String()
}

func (s *DescribeRCInstanceTypeFamiliesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRCInstanceTypeFamiliesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRCInstanceTypeFamiliesResponse) GetBody() *DescribeRCInstanceTypeFamiliesResponseBody {
	return s.Body
}

func (s *DescribeRCInstanceTypeFamiliesResponse) SetHeaders(v map[string]*string) *DescribeRCInstanceTypeFamiliesResponse {
	s.Headers = v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponse) SetStatusCode(v int32) *DescribeRCInstanceTypeFamiliesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponse) SetBody(v *DescribeRCInstanceTypeFamiliesResponseBody) *DescribeRCInstanceTypeFamiliesResponse {
	s.Body = v
	return s
}

func (s *DescribeRCInstanceTypeFamiliesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePackgeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePackgeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePackgeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribePackgeInfoResponseBody) *DescribePackgeInfoResponse
	GetBody() *DescribePackgeInfoResponseBody
}

type DescribePackgeInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePackgeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePackgeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePackgeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribePackgeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePackgeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePackgeInfoResponse) GetBody() *DescribePackgeInfoResponseBody {
	return s.Body
}

func (s *DescribePackgeInfoResponse) SetHeaders(v map[string]*string) *DescribePackgeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribePackgeInfoResponse) SetStatusCode(v int32) *DescribePackgeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePackgeInfoResponse) SetBody(v *DescribePackgeInfoResponseBody) *DescribePackgeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribePackgeInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

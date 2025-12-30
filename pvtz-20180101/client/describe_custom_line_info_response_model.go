// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomLineInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeCustomLineInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeCustomLineInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeCustomLineInfoResponseBody) *DescribeCustomLineInfoResponse
	GetBody() *DescribeCustomLineInfoResponseBody
}

type DescribeCustomLineInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeCustomLineInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeCustomLineInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomLineInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeCustomLineInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeCustomLineInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeCustomLineInfoResponse) GetBody() *DescribeCustomLineInfoResponseBody {
	return s.Body
}

func (s *DescribeCustomLineInfoResponse) SetHeaders(v map[string]*string) *DescribeCustomLineInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeCustomLineInfoResponse) SetStatusCode(v int32) *DescribeCustomLineInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeCustomLineInfoResponse) SetBody(v *DescribeCustomLineInfoResponseBody) *DescribeCustomLineInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeCustomLineInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

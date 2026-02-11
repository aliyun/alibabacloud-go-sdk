// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEntityInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeEntityInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeEntityInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeEntityInfoResponseBody) *DescribeEntityInfoResponse
	GetBody() *DescribeEntityInfoResponseBody
}

type DescribeEntityInfoResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeEntityInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeEntityInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeEntityInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeEntityInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeEntityInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeEntityInfoResponse) GetBody() *DescribeEntityInfoResponseBody {
	return s.Body
}

func (s *DescribeEntityInfoResponse) SetHeaders(v map[string]*string) *DescribeEntityInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeEntityInfoResponse) SetStatusCode(v int32) *DescribeEntityInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeEntityInfoResponse) SetBody(v *DescribeEntityInfoResponseBody) *DescribeEntityInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeEntityInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

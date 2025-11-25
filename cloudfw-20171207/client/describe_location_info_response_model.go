// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLocationInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLocationInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLocationInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLocationInfoResponseBody) *DescribeLocationInfoResponse
	GetBody() *DescribeLocationInfoResponseBody
}

type DescribeLocationInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLocationInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLocationInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLocationInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLocationInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLocationInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLocationInfoResponse) GetBody() *DescribeLocationInfoResponseBody {
	return s.Body
}

func (s *DescribeLocationInfoResponse) SetHeaders(v map[string]*string) *DescribeLocationInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLocationInfoResponse) SetStatusCode(v int32) *DescribeLocationInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLocationInfoResponse) SetBody(v *DescribeLocationInfoResponseBody) *DescribeLocationInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLocationInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

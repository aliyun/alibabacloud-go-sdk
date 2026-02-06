// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGroupExecutingInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGroupExecutingInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGroupExecutingInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGroupExecutingInfoResponseBody) *DescribeGroupExecutingInfoResponse
	GetBody() *DescribeGroupExecutingInfoResponseBody
}

type DescribeGroupExecutingInfoResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGroupExecutingInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGroupExecutingInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGroupExecutingInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeGroupExecutingInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGroupExecutingInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGroupExecutingInfoResponse) GetBody() *DescribeGroupExecutingInfoResponseBody {
	return s.Body
}

func (s *DescribeGroupExecutingInfoResponse) SetHeaders(v map[string]*string) *DescribeGroupExecutingInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeGroupExecutingInfoResponse) SetStatusCode(v int32) *DescribeGroupExecutingInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGroupExecutingInfoResponse) SetBody(v *DescribeGroupExecutingInfoResponseBody) *DescribeGroupExecutingInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeGroupExecutingInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

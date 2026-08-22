// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchInfoResponseBody) *DescribeOpenSearchInfoResponse
	GetBody() *DescribeOpenSearchInfoResponseBody
}

type DescribeOpenSearchInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchInfoResponse) GetBody() *DescribeOpenSearchInfoResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchInfoResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchInfoResponse) SetStatusCode(v int32) *DescribeOpenSearchInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchInfoResponse) SetBody(v *DescribeOpenSearchInfoResponseBody) *DescribeOpenSearchInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

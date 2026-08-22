// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeOpenSearchAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeOpenSearchAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeOpenSearchAccountInfoResponseBody) *DescribeOpenSearchAccountInfoResponse
	GetBody() *DescribeOpenSearchAccountInfoResponseBody
}

type DescribeOpenSearchAccountInfoResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeOpenSearchAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeOpenSearchAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeOpenSearchAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeOpenSearchAccountInfoResponse) GetBody() *DescribeOpenSearchAccountInfoResponseBody {
	return s.Body
}

func (s *DescribeOpenSearchAccountInfoResponse) SetHeaders(v map[string]*string) *DescribeOpenSearchAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponse) SetStatusCode(v int32) *DescribeOpenSearchAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponse) SetBody(v *DescribeOpenSearchAccountInfoResponseBody) *DescribeOpenSearchAccountInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeOpenSearchAccountInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

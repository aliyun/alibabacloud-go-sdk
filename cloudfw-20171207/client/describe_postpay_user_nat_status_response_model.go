// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePostpayUserNatStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribePostpayUserNatStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribePostpayUserNatStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribePostpayUserNatStatusResponseBody) *DescribePostpayUserNatStatusResponse
	GetBody() *DescribePostpayUserNatStatusResponseBody
}

type DescribePostpayUserNatStatusResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribePostpayUserNatStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribePostpayUserNatStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribePostpayUserNatStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribePostpayUserNatStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribePostpayUserNatStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribePostpayUserNatStatusResponse) GetBody() *DescribePostpayUserNatStatusResponseBody {
	return s.Body
}

func (s *DescribePostpayUserNatStatusResponse) SetHeaders(v map[string]*string) *DescribePostpayUserNatStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribePostpayUserNatStatusResponse) SetStatusCode(v int32) *DescribePostpayUserNatStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribePostpayUserNatStatusResponse) SetBody(v *DescribePostpayUserNatStatusResponseBody) *DescribePostpayUserNatStatusResponse {
	s.Body = v
	return s
}

func (s *DescribePostpayUserNatStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

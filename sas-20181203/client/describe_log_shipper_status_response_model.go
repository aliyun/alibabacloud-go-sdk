// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogShipperStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLogShipperStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLogShipperStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLogShipperStatusResponseBody) *DescribeLogShipperStatusResponse
	GetBody() *DescribeLogShipperStatusResponseBody
}

type DescribeLogShipperStatusResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLogShipperStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLogShipperStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogShipperStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeLogShipperStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLogShipperStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLogShipperStatusResponse) GetBody() *DescribeLogShipperStatusResponseBody {
	return s.Body
}

func (s *DescribeLogShipperStatusResponse) SetHeaders(v map[string]*string) *DescribeLogShipperStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeLogShipperStatusResponse) SetStatusCode(v int32) *DescribeLogShipperStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLogShipperStatusResponse) SetBody(v *DescribeLogShipperStatusResponseBody) *DescribeLogShipperStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeLogShipperStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

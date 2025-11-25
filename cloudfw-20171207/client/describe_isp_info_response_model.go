// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeIspInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeIspInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeIspInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeIspInfoResponseBody) *DescribeIspInfoResponse
	GetBody() *DescribeIspInfoResponseBody
}

type DescribeIspInfoResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeIspInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeIspInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeIspInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeIspInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeIspInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeIspInfoResponse) GetBody() *DescribeIspInfoResponseBody {
	return s.Body
}

func (s *DescribeIspInfoResponse) SetHeaders(v map[string]*string) *DescribeIspInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeIspInfoResponse) SetStatusCode(v int32) *DescribeIspInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeIspInfoResponse) SetBody(v *DescribeIspInfoResponseBody) *DescribeIspInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeIspInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

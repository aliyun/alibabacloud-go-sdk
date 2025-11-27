// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeADInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeADInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeADInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeADInfoResponseBody) *DescribeADInfoResponse
	GetBody() *DescribeADInfoResponseBody
}

type DescribeADInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeADInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeADInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeADInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeADInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeADInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeADInfoResponse) GetBody() *DescribeADInfoResponseBody {
	return s.Body
}

func (s *DescribeADInfoResponse) SetHeaders(v map[string]*string) *DescribeADInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeADInfoResponse) SetStatusCode(v int32) *DescribeADInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeADInfoResponse) SetBody(v *DescribeADInfoResponseBody) *DescribeADInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeADInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRdStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRdStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRdStatusResponseBody) *DescribeRdStatusResponse
	GetBody() *DescribeRdStatusResponseBody
}

type DescribeRdStatusResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRdStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRdStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRdStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRdStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRdStatusResponse) GetBody() *DescribeRdStatusResponseBody {
	return s.Body
}

func (s *DescribeRdStatusResponse) SetHeaders(v map[string]*string) *DescribeRdStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRdStatusResponse) SetStatusCode(v int32) *DescribeRdStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRdStatusResponse) SetBody(v *DescribeRdStatusResponseBody) *DescribeRdStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRdStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

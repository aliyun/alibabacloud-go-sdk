// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelDetailsByIdResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeModelDetailsByIdResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeModelDetailsByIdResponse
	GetStatusCode() *int32
	SetBody(v *DescribeModelDetailsByIdResponseBody) *DescribeModelDetailsByIdResponse
	GetBody() *DescribeModelDetailsByIdResponseBody
}

type DescribeModelDetailsByIdResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeModelDetailsByIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeModelDetailsByIdResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelDetailsByIdResponse) GoString() string {
	return s.String()
}

func (s *DescribeModelDetailsByIdResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeModelDetailsByIdResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeModelDetailsByIdResponse) GetBody() *DescribeModelDetailsByIdResponseBody {
	return s.Body
}

func (s *DescribeModelDetailsByIdResponse) SetHeaders(v map[string]*string) *DescribeModelDetailsByIdResponse {
	s.Headers = v
	return s
}

func (s *DescribeModelDetailsByIdResponse) SetStatusCode(v int32) *DescribeModelDetailsByIdResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeModelDetailsByIdResponse) SetBody(v *DescribeModelDetailsByIdResponseBody) *DescribeModelDetailsByIdResponse {
	s.Body = v
	return s
}

func (s *DescribeModelDetailsByIdResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

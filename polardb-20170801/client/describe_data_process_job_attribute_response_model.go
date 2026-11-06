// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataProcessJobAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataProcessJobAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataProcessJobAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataProcessJobAttributeResponseBody) *DescribeDataProcessJobAttributeResponse
	GetBody() *DescribeDataProcessJobAttributeResponseBody
}

type DescribeDataProcessJobAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataProcessJobAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataProcessJobAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataProcessJobAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataProcessJobAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataProcessJobAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataProcessJobAttributeResponse) GetBody() *DescribeDataProcessJobAttributeResponseBody {
	return s.Body
}

func (s *DescribeDataProcessJobAttributeResponse) SetHeaders(v map[string]*string) *DescribeDataProcessJobAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataProcessJobAttributeResponse) SetStatusCode(v int32) *DescribeDataProcessJobAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataProcessJobAttributeResponse) SetBody(v *DescribeDataProcessJobAttributeResponseBody) *DescribeDataProcessJobAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeDataProcessJobAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

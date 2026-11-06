// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceReportAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInstanceReportAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInstanceReportAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInstanceReportAttributeResponseBody) *DescribeInstanceReportAttributeResponse
	GetBody() *DescribeInstanceReportAttributeResponseBody
}

type DescribeInstanceReportAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInstanceReportAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInstanceReportAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceReportAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInstanceReportAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInstanceReportAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInstanceReportAttributeResponse) GetBody() *DescribeInstanceReportAttributeResponseBody {
	return s.Body
}

func (s *DescribeInstanceReportAttributeResponse) SetHeaders(v map[string]*string) *DescribeInstanceReportAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInstanceReportAttributeResponse) SetStatusCode(v int32) *DescribeInstanceReportAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInstanceReportAttributeResponse) SetBody(v *DescribeInstanceReportAttributeResponseBody) *DescribeInstanceReportAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeInstanceReportAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

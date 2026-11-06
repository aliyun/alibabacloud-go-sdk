// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionTaskAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionTaskAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionTaskAttributeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionTaskAttributeResponseBody) *DescribeInspectionTaskAttributeResponse
	GetBody() *DescribeInspectionTaskAttributeResponseBody
}

type DescribeInspectionTaskAttributeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionTaskAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionTaskAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionTaskAttributeResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionTaskAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionTaskAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionTaskAttributeResponse) GetBody() *DescribeInspectionTaskAttributeResponseBody {
	return s.Body
}

func (s *DescribeInspectionTaskAttributeResponse) SetHeaders(v map[string]*string) *DescribeInspectionTaskAttributeResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionTaskAttributeResponse) SetStatusCode(v int32) *DescribeInspectionTaskAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionTaskAttributeResponse) SetBody(v *DescribeInspectionTaskAttributeResponseBody) *DescribeInspectionTaskAttributeResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionTaskAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

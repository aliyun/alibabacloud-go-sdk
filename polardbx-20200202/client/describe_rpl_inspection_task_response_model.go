// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRplInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRplInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRplInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRplInspectionTaskResponseBody) *DescribeRplInspectionTaskResponse
	GetBody() *DescribeRplInspectionTaskResponseBody
}

type DescribeRplInspectionTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRplInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRplInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRplInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeRplInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRplInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRplInspectionTaskResponse) GetBody() *DescribeRplInspectionTaskResponseBody {
	return s.Body
}

func (s *DescribeRplInspectionTaskResponse) SetHeaders(v map[string]*string) *DescribeRplInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeRplInspectionTaskResponse) SetStatusCode(v int32) *DescribeRplInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRplInspectionTaskResponse) SetBody(v *DescribeRplInspectionTaskResponseBody) *DescribeRplInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeRplInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

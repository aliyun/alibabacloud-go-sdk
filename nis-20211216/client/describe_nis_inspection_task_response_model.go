// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNisInspectionTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNisInspectionTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNisInspectionTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNisInspectionTaskResponseBody) *DescribeNisInspectionTaskResponse
	GetBody() *DescribeNisInspectionTaskResponseBody
}

type DescribeNisInspectionTaskResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNisInspectionTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNisInspectionTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNisInspectionTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeNisInspectionTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNisInspectionTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNisInspectionTaskResponse) GetBody() *DescribeNisInspectionTaskResponseBody {
	return s.Body
}

func (s *DescribeNisInspectionTaskResponse) SetHeaders(v map[string]*string) *DescribeNisInspectionTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeNisInspectionTaskResponse) SetStatusCode(v int32) *DescribeNisInspectionTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNisInspectionTaskResponse) SetBody(v *DescribeNisInspectionTaskResponseBody) *DescribeNisInspectionTaskResponse {
	s.Body = v
	return s
}

func (s *DescribeNisInspectionTaskResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

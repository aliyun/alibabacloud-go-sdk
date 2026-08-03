// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInspectionSchedulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInspectionSchedulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInspectionSchedulesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInspectionSchedulesResponseBody) *DescribeInspectionSchedulesResponse
	GetBody() *DescribeInspectionSchedulesResponseBody
}

type DescribeInspectionSchedulesResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInspectionSchedulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInspectionSchedulesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInspectionSchedulesResponse) GoString() string {
	return s.String()
}

func (s *DescribeInspectionSchedulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInspectionSchedulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInspectionSchedulesResponse) GetBody() *DescribeInspectionSchedulesResponseBody {
	return s.Body
}

func (s *DescribeInspectionSchedulesResponse) SetHeaders(v map[string]*string) *DescribeInspectionSchedulesResponse {
	s.Headers = v
	return s
}

func (s *DescribeInspectionSchedulesResponse) SetStatusCode(v int32) *DescribeInspectionSchedulesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInspectionSchedulesResponse) SetBody(v *DescribeInspectionSchedulesResponseBody) *DescribeInspectionSchedulesResponse {
	s.Body = v
	return s
}

func (s *DescribeInspectionSchedulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

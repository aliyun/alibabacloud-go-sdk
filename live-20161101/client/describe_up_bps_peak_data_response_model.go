// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpBpsPeakDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpBpsPeakDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpBpsPeakDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpBpsPeakDataResponseBody) *DescribeUpBpsPeakDataResponse
	GetBody() *DescribeUpBpsPeakDataResponseBody
}

type DescribeUpBpsPeakDataResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpBpsPeakDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpBpsPeakDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpBpsPeakDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpBpsPeakDataResponse) GetBody() *DescribeUpBpsPeakDataResponseBody {
	return s.Body
}

func (s *DescribeUpBpsPeakDataResponse) SetHeaders(v map[string]*string) *DescribeUpBpsPeakDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpBpsPeakDataResponse) SetStatusCode(v int32) *DescribeUpBpsPeakDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpBpsPeakDataResponse) SetBody(v *DescribeUpBpsPeakDataResponseBody) *DescribeUpBpsPeakDataResponse {
	s.Body = v
	return s
}

func (s *DescribeUpBpsPeakDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

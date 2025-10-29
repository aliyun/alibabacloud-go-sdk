// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUpBpsPeakOfLineResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeUpBpsPeakOfLineResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeUpBpsPeakOfLineResponse
	GetStatusCode() *int32
	SetBody(v *DescribeUpBpsPeakOfLineResponseBody) *DescribeUpBpsPeakOfLineResponse
	GetBody() *DescribeUpBpsPeakOfLineResponseBody
}

type DescribeUpBpsPeakOfLineResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeUpBpsPeakOfLineResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeUpBpsPeakOfLineResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeUpBpsPeakOfLineResponse) GoString() string {
	return s.String()
}

func (s *DescribeUpBpsPeakOfLineResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeUpBpsPeakOfLineResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeUpBpsPeakOfLineResponse) GetBody() *DescribeUpBpsPeakOfLineResponseBody {
	return s.Body
}

func (s *DescribeUpBpsPeakOfLineResponse) SetHeaders(v map[string]*string) *DescribeUpBpsPeakOfLineResponse {
	s.Headers = v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponse) SetStatusCode(v int32) *DescribeUpBpsPeakOfLineResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponse) SetBody(v *DescribeUpBpsPeakOfLineResponseBody) *DescribeUpBpsPeakOfLineResponse {
	s.Body = v
	return s
}

func (s *DescribeUpBpsPeakOfLineResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

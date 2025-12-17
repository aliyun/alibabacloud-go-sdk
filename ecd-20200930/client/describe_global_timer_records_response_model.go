// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeGlobalTimerRecordsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeGlobalTimerRecordsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeGlobalTimerRecordsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeGlobalTimerRecordsResponseBody) *DescribeGlobalTimerRecordsResponse
	GetBody() *DescribeGlobalTimerRecordsResponseBody
}

type DescribeGlobalTimerRecordsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeGlobalTimerRecordsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeGlobalTimerRecordsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeGlobalTimerRecordsResponse) GoString() string {
	return s.String()
}

func (s *DescribeGlobalTimerRecordsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeGlobalTimerRecordsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeGlobalTimerRecordsResponse) GetBody() *DescribeGlobalTimerRecordsResponseBody {
	return s.Body
}

func (s *DescribeGlobalTimerRecordsResponse) SetHeaders(v map[string]*string) *DescribeGlobalTimerRecordsResponse {
	s.Headers = v
	return s
}

func (s *DescribeGlobalTimerRecordsResponse) SetStatusCode(v int32) *DescribeGlobalTimerRecordsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeGlobalTimerRecordsResponse) SetBody(v *DescribeGlobalTimerRecordsResponseBody) *DescribeGlobalTimerRecordsResponse {
	s.Body = v
	return s
}

func (s *DescribeGlobalTimerRecordsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDataPushResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDataPushResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDataPushResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDataPushResultResponseBody) *DescribeDataPushResultResponse
	GetBody() *DescribeDataPushResultResponseBody
}

type DescribeDataPushResultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDataPushResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDataPushResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDataPushResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDataPushResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDataPushResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDataPushResultResponse) GetBody() *DescribeDataPushResultResponseBody {
	return s.Body
}

func (s *DescribeDataPushResultResponse) SetHeaders(v map[string]*string) *DescribeDataPushResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDataPushResultResponse) SetStatusCode(v int32) *DescribeDataPushResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDataPushResultResponse) SetBody(v *DescribeDataPushResultResponseBody) *DescribeDataPushResultResponse {
	s.Body = v
	return s
}

func (s *DescribeDataPushResultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

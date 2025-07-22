// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcDurationDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcDurationDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcDurationDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcDurationDataResponseBody) *DescribeRtcDurationDataResponse
	GetBody() *DescribeRtcDurationDataResponseBody
}

type DescribeRtcDurationDataResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcDurationDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcDurationDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcDurationDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcDurationDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcDurationDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcDurationDataResponse) GetBody() *DescribeRtcDurationDataResponseBody {
	return s.Body
}

func (s *DescribeRtcDurationDataResponse) SetHeaders(v map[string]*string) *DescribeRtcDurationDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcDurationDataResponse) SetStatusCode(v int32) *DescribeRtcDurationDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcDurationDataResponse) SetBody(v *DescribeRtcDurationDataResponseBody) *DescribeRtcDurationDataResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcDurationDataResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKvRealTimeQpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeKvRealTimeQpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeKvRealTimeQpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeKvRealTimeQpsDataResponseBody) *DescribeKvRealTimeQpsDataResponse
	GetBody() *DescribeKvRealTimeQpsDataResponseBody
}

type DescribeKvRealTimeQpsDataResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeKvRealTimeQpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeKvRealTimeQpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeKvRealTimeQpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeKvRealTimeQpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeKvRealTimeQpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeKvRealTimeQpsDataResponse) GetBody() *DescribeKvRealTimeQpsDataResponseBody {
	return s.Body
}

func (s *DescribeKvRealTimeQpsDataResponse) SetHeaders(v map[string]*string) *DescribeKvRealTimeQpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponse) SetStatusCode(v int32) *DescribeKvRealTimeQpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponse) SetBody(v *DescribeKvRealTimeQpsDataResponseBody) *DescribeKvRealTimeQpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeKvRealTimeQpsDataResponse) Validate() error {
	return dara.Validate(s)
}

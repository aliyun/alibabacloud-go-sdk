// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAppLiveStreamStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeAppLiveStreamStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeAppLiveStreamStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeAppLiveStreamStatusResponseBody) *DescribeAppLiveStreamStatusResponse
	GetBody() *DescribeAppLiveStreamStatusResponseBody
}

type DescribeAppLiveStreamStatusResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeAppLiveStreamStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeAppLiveStreamStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeAppLiveStreamStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeAppLiveStreamStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeAppLiveStreamStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeAppLiveStreamStatusResponse) GetBody() *DescribeAppLiveStreamStatusResponseBody {
	return s.Body
}

func (s *DescribeAppLiveStreamStatusResponse) SetHeaders(v map[string]*string) *DescribeAppLiveStreamStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeAppLiveStreamStatusResponse) SetStatusCode(v int32) *DescribeAppLiveStreamStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeAppLiveStreamStatusResponse) SetBody(v *DescribeAppLiveStreamStatusResponseBody) *DescribeAppLiveStreamStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeAppLiveStreamStatusResponse) Validate() error {
	return dara.Validate(s)
}

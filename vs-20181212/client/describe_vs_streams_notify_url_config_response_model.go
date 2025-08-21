// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVsStreamsNotifyUrlConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVsStreamsNotifyUrlConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVsStreamsNotifyUrlConfigResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVsStreamsNotifyUrlConfigResponseBody) *DescribeVsStreamsNotifyUrlConfigResponse
	GetBody() *DescribeVsStreamsNotifyUrlConfigResponseBody
}

type DescribeVsStreamsNotifyUrlConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVsStreamsNotifyUrlConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVsStreamsNotifyUrlConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVsStreamsNotifyUrlConfigResponse) GoString() string {
	return s.String()
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) GetBody() *DescribeVsStreamsNotifyUrlConfigResponseBody {
	return s.Body
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) SetHeaders(v map[string]*string) *DescribeVsStreamsNotifyUrlConfigResponse {
	s.Headers = v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) SetStatusCode(v int32) *DescribeVsStreamsNotifyUrlConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) SetBody(v *DescribeVsStreamsNotifyUrlConfigResponseBody) *DescribeVsStreamsNotifyUrlConfigResponse {
	s.Body = v
	return s
}

func (s *DescribeVsStreamsNotifyUrlConfigResponse) Validate() error {
	return dara.Validate(s)
}

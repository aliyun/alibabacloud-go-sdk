// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodRealtimeLogAuthorizedResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodRealtimeLogAuthorizedResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodRealtimeLogAuthorizedResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodRealtimeLogAuthorizedResponseBody) *DescribeVodRealtimeLogAuthorizedResponse
	GetBody() *DescribeVodRealtimeLogAuthorizedResponseBody
}

type DescribeVodRealtimeLogAuthorizedResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodRealtimeLogAuthorizedResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodRealtimeLogAuthorizedResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodRealtimeLogAuthorizedResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) GetBody() *DescribeVodRealtimeLogAuthorizedResponseBody {
	return s.Body
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) SetHeaders(v map[string]*string) *DescribeVodRealtimeLogAuthorizedResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) SetStatusCode(v int32) *DescribeVodRealtimeLogAuthorizedResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) SetBody(v *DescribeVodRealtimeLogAuthorizedResponseBody) *DescribeVodRealtimeLogAuthorizedResponse {
	s.Body = v
	return s
}

func (s *DescribeVodRealtimeLogAuthorizedResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRTSNativeSDKPlayTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRTSNativeSDKPlayTimeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRTSNativeSDKPlayTimeResponseBody) *DescribeRTSNativeSDKPlayTimeResponse
	GetBody() *DescribeRTSNativeSDKPlayTimeResponseBody
}

type DescribeRTSNativeSDKPlayTimeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRTSNativeSDKPlayTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRTSNativeSDKPlayTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayTimeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) GetBody() *DescribeRTSNativeSDKPlayTimeResponseBody {
	return s.Body
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) SetHeaders(v map[string]*string) *DescribeRTSNativeSDKPlayTimeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) SetStatusCode(v int32) *DescribeRTSNativeSDKPlayTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) SetBody(v *DescribeRTSNativeSDKPlayTimeResponseBody) *DescribeRTSNativeSDKPlayTimeResponse {
	s.Body = v
	return s
}

func (s *DescribeRTSNativeSDKPlayTimeResponse) Validate() error {
	return dara.Validate(s)
}

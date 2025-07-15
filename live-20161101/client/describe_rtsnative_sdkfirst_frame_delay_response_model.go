// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameDelayResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRTSNativeSDKFirstFrameDelayResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRTSNativeSDKFirstFrameDelayResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRTSNativeSDKFirstFrameDelayResponseBody) *DescribeRTSNativeSDKFirstFrameDelayResponse
	GetBody() *DescribeRTSNativeSDKFirstFrameDelayResponseBody
}

type DescribeRTSNativeSDKFirstFrameDelayResponse struct {
	Headers    map[string]*string                               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRTSNativeSDKFirstFrameDelayResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameDelayResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameDelayResponse) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) GetBody() *DescribeRTSNativeSDKFirstFrameDelayResponseBody {
	return s.Body
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) SetHeaders(v map[string]*string) *DescribeRTSNativeSDKFirstFrameDelayResponse {
	s.Headers = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) SetStatusCode(v int32) *DescribeRTSNativeSDKFirstFrameDelayResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) SetBody(v *DescribeRTSNativeSDKFirstFrameDelayResponseBody) *DescribeRTSNativeSDKFirstFrameDelayResponse {
	s.Body = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameDelayResponse) Validate() error {
	return dara.Validate(s)
}

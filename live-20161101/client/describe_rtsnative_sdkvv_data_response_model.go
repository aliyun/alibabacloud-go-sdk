// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKVvDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRTSNativeSDKVvDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRTSNativeSDKVvDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRTSNativeSDKVvDataResponseBody) *DescribeRTSNativeSDKVvDataResponse
	GetBody() *DescribeRTSNativeSDKVvDataResponseBody
}

type DescribeRTSNativeSDKVvDataResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRTSNativeSDKVvDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRTSNativeSDKVvDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKVvDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKVvDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRTSNativeSDKVvDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRTSNativeSDKVvDataResponse) GetBody() *DescribeRTSNativeSDKVvDataResponseBody {
	return s.Body
}

func (s *DescribeRTSNativeSDKVvDataResponse) SetHeaders(v map[string]*string) *DescribeRTSNativeSDKVvDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponse) SetStatusCode(v int32) *DescribeRTSNativeSDKVvDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponse) SetBody(v *DescribeRTSNativeSDKVvDataResponseBody) *DescribeRTSNativeSDKVvDataResponse {
	s.Body = v
	return s
}

func (s *DescribeRTSNativeSDKVvDataResponse) Validate() error {
	return dara.Validate(s)
}

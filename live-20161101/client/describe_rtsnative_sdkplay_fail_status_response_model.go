// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKPlayFailStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRTSNativeSDKPlayFailStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRTSNativeSDKPlayFailStatusResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRTSNativeSDKPlayFailStatusResponseBody) *DescribeRTSNativeSDKPlayFailStatusResponse
	GetBody() *DescribeRTSNativeSDKPlayFailStatusResponseBody
}

type DescribeRTSNativeSDKPlayFailStatusResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRTSNativeSDKPlayFailStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRTSNativeSDKPlayFailStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKPlayFailStatusResponse) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) GetBody() *DescribeRTSNativeSDKPlayFailStatusResponseBody {
	return s.Body
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) SetHeaders(v map[string]*string) *DescribeRTSNativeSDKPlayFailStatusResponse {
	s.Headers = v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) SetStatusCode(v int32) *DescribeRTSNativeSDKPlayFailStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) SetBody(v *DescribeRTSNativeSDKPlayFailStatusResponseBody) *DescribeRTSNativeSDKPlayFailStatusResponse {
	s.Body = v
	return s
}

func (s *DescribeRTSNativeSDKPlayFailStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

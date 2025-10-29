// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRTSNativeSDKFirstFrameCostResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRTSNativeSDKFirstFrameCostResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRTSNativeSDKFirstFrameCostResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRTSNativeSDKFirstFrameCostResponseBody) *DescribeRTSNativeSDKFirstFrameCostResponse
	GetBody() *DescribeRTSNativeSDKFirstFrameCostResponseBody
}

type DescribeRTSNativeSDKFirstFrameCostResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRTSNativeSDKFirstFrameCostResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRTSNativeSDKFirstFrameCostResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRTSNativeSDKFirstFrameCostResponse) GoString() string {
	return s.String()
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) GetBody() *DescribeRTSNativeSDKFirstFrameCostResponseBody {
	return s.Body
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) SetHeaders(v map[string]*string) *DescribeRTSNativeSDKFirstFrameCostResponse {
	s.Headers = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) SetStatusCode(v int32) *DescribeRTSNativeSDKFirstFrameCostResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) SetBody(v *DescribeRTSNativeSDKFirstFrameCostResponseBody) *DescribeRTSNativeSDKFirstFrameCostResponse {
	s.Body = v
	return s
}

func (s *DescribeRTSNativeSDKFirstFrameCostResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

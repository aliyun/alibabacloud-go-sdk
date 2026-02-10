// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRtcCloudTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRtcCloudTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRtcCloudTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRtcCloudTranscodeResponseBody) *DescribeRtcCloudTranscodeResponse
	GetBody() *DescribeRtcCloudTranscodeResponseBody
}

type DescribeRtcCloudTranscodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRtcCloudTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRtcCloudTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRtcCloudTranscodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeRtcCloudTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRtcCloudTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRtcCloudTranscodeResponse) GetBody() *DescribeRtcCloudTranscodeResponseBody {
	return s.Body
}

func (s *DescribeRtcCloudTranscodeResponse) SetHeaders(v map[string]*string) *DescribeRtcCloudTranscodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeRtcCloudTranscodeResponse) SetStatusCode(v int32) *DescribeRtcCloudTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRtcCloudTranscodeResponse) SetBody(v *DescribeRtcCloudTranscodeResponseBody) *DescribeRtcCloudTranscodeResponse {
	s.Body = v
	return s
}

func (s *DescribeRtcCloudTranscodeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHlsLiveStreamRealTimeBpsDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeHlsLiveStreamRealTimeBpsDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeHlsLiveStreamRealTimeBpsDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) *DescribeHlsLiveStreamRealTimeBpsDataResponse
	GetBody() *DescribeHlsLiveStreamRealTimeBpsDataResponseBody
}

type DescribeHlsLiveStreamRealTimeBpsDataResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeHlsLiveStreamRealTimeBpsDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeHlsLiveStreamRealTimeBpsDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) GetBody() *DescribeHlsLiveStreamRealTimeBpsDataResponseBody {
	return s.Body
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) SetHeaders(v map[string]*string) *DescribeHlsLiveStreamRealTimeBpsDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) SetStatusCode(v int32) *DescribeHlsLiveStreamRealTimeBpsDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) SetBody(v *DescribeHlsLiveStreamRealTimeBpsDataResponseBody) *DescribeHlsLiveStreamRealTimeBpsDataResponse {
	s.Body = v
	return s
}

func (s *DescribeHlsLiveStreamRealTimeBpsDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

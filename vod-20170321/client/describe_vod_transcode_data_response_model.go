// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodTranscodeDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVodTranscodeDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVodTranscodeDataResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVodTranscodeDataResponseBody) *DescribeVodTranscodeDataResponse
	GetBody() *DescribeVodTranscodeDataResponseBody
}

type DescribeVodTranscodeDataResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVodTranscodeDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVodTranscodeDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodTranscodeDataResponse) GoString() string {
	return s.String()
}

func (s *DescribeVodTranscodeDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVodTranscodeDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVodTranscodeDataResponse) GetBody() *DescribeVodTranscodeDataResponseBody {
	return s.Body
}

func (s *DescribeVodTranscodeDataResponse) SetHeaders(v map[string]*string) *DescribeVodTranscodeDataResponse {
	s.Headers = v
	return s
}

func (s *DescribeVodTranscodeDataResponse) SetStatusCode(v int32) *DescribeVodTranscodeDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVodTranscodeDataResponse) SetBody(v *DescribeVodTranscodeDataResponseBody) *DescribeVodTranscodeDataResponse {
	s.Body = v
	return s
}

func (s *DescribeVodTranscodeDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

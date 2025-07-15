// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamTranscodeInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamTranscodeInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamTranscodeInfoResponseBody) *DescribeLiveStreamTranscodeInfoResponse
	GetBody() *DescribeLiveStreamTranscodeInfoResponseBody
}

type DescribeLiveStreamTranscodeInfoResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamTranscodeInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamTranscodeInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamTranscodeInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamTranscodeInfoResponse) GetBody() *DescribeLiveStreamTranscodeInfoResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamTranscodeInfoResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamTranscodeInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponse) SetStatusCode(v int32) *DescribeLiveStreamTranscodeInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponse) SetBody(v *DescribeLiveStreamTranscodeInfoResponseBody) *DescribeLiveStreamTranscodeInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamTranscodeInfoResponse) Validate() error {
	return dara.Validate(s)
}

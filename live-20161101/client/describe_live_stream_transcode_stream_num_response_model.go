// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeStreamNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveStreamTranscodeStreamNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveStreamTranscodeStreamNumResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveStreamTranscodeStreamNumResponseBody) *DescribeLiveStreamTranscodeStreamNumResponse
	GetBody() *DescribeLiveStreamTranscodeStreamNumResponseBody
}

type DescribeLiveStreamTranscodeStreamNumResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveStreamTranscodeStreamNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveStreamTranscodeStreamNumResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeStreamNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) GetBody() *DescribeLiveStreamTranscodeStreamNumResponseBody {
	return s.Body
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetHeaders(v map[string]*string) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetStatusCode(v int32) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) SetBody(v *DescribeLiveStreamTranscodeStreamNumResponseBody) *DescribeLiveStreamTranscodeStreamNumResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponse) Validate() error {
	return dara.Validate(s)
}

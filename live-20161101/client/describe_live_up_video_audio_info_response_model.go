// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveUpVideoAudioInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveUpVideoAudioInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveUpVideoAudioInfoResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveUpVideoAudioInfoResponseBody) *DescribeLiveUpVideoAudioInfoResponse
	GetBody() *DescribeLiveUpVideoAudioInfoResponseBody
}

type DescribeLiveUpVideoAudioInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveUpVideoAudioInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveUpVideoAudioInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveUpVideoAudioInfoResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveUpVideoAudioInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveUpVideoAudioInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveUpVideoAudioInfoResponse) GetBody() *DescribeLiveUpVideoAudioInfoResponseBody {
	return s.Body
}

func (s *DescribeLiveUpVideoAudioInfoResponse) SetHeaders(v map[string]*string) *DescribeLiveUpVideoAudioInfoResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponse) SetStatusCode(v int32) *DescribeLiveUpVideoAudioInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponse) SetBody(v *DescribeLiveUpVideoAudioInfoResponseBody) *DescribeLiveUpVideoAudioInfoResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveUpVideoAudioInfoResponse) Validate() error {
	return dara.Validate(s)
}

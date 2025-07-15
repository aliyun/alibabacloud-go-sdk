// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveStreamTranscodeResponseBody) *AddLiveStreamTranscodeResponse
	GetBody() *AddLiveStreamTranscodeResponseBody
}

type AddLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *AddLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveStreamTranscodeResponse) GetBody() *AddLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *AddLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *AddLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *AddLiveStreamTranscodeResponse) SetStatusCode(v int32) *AddLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveStreamTranscodeResponse) SetBody(v *AddLiveStreamTranscodeResponseBody) *AddLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *AddLiveStreamTranscodeResponse) Validate() error {
	return dara.Validate(s)
}

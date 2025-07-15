// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCustomLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCustomLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCustomLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *AddCustomLiveStreamTranscodeResponseBody) *AddCustomLiveStreamTranscodeResponse
	GetBody() *AddCustomLiveStreamTranscodeResponseBody
}

type AddCustomLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCustomLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCustomLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCustomLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *AddCustomLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCustomLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCustomLiveStreamTranscodeResponse) GetBody() *AddCustomLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *AddCustomLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *AddCustomLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *AddCustomLiveStreamTranscodeResponse) SetStatusCode(v int32) *AddCustomLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCustomLiveStreamTranscodeResponse) SetBody(v *AddCustomLiveStreamTranscodeResponseBody) *AddCustomLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *AddCustomLiveStreamTranscodeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveStreamTranscodeResponseBody) *UpdateLiveStreamTranscodeResponse
	GetBody() *UpdateLiveStreamTranscodeResponseBody
}

type UpdateLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveStreamTranscodeResponse) GetBody() *UpdateLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *UpdateLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *UpdateLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveStreamTranscodeResponse) SetStatusCode(v int32) *UpdateLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveStreamTranscodeResponse) SetBody(v *UpdateLiveStreamTranscodeResponseBody) *UpdateLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveStreamTranscodeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateRtsLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateRtsLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateRtsLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateRtsLiveStreamTranscodeResponseBody) *UpdateRtsLiveStreamTranscodeResponse
	GetBody() *UpdateRtsLiveStreamTranscodeResponseBody
}

type UpdateRtsLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateRtsLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRtsLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateRtsLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateRtsLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateRtsLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateRtsLiveStreamTranscodeResponse) GetBody() *UpdateRtsLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *UpdateRtsLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *UpdateRtsLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeResponse) SetStatusCode(v int32) *UpdateRtsLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeResponse) SetBody(v *UpdateRtsLiveStreamTranscodeResponseBody) *UpdateRtsLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *UpdateRtsLiveStreamTranscodeResponse) Validate() error {
	return dara.Validate(s)
}

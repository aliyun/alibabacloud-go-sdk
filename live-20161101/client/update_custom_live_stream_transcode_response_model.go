// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustomLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCustomLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCustomLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCustomLiveStreamTranscodeResponseBody) *UpdateCustomLiveStreamTranscodeResponse
	GetBody() *UpdateCustomLiveStreamTranscodeResponseBody
}

type UpdateCustomLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCustomLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCustomLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustomLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *UpdateCustomLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCustomLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCustomLiveStreamTranscodeResponse) GetBody() *UpdateCustomLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *UpdateCustomLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *UpdateCustomLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeResponse) SetStatusCode(v int32) *UpdateCustomLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeResponse) SetBody(v *UpdateCustomLiveStreamTranscodeResponseBody) *UpdateCustomLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *UpdateCustomLiveStreamTranscodeResponse) Validate() error {
	return dara.Validate(s)
}

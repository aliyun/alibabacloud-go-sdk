// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamTranscodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamTranscodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamTranscodeResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamTranscodeResponseBody) *DeleteLiveStreamTranscodeResponse
	GetBody() *DeleteLiveStreamTranscodeResponseBody
}

type DeleteLiveStreamTranscodeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamTranscodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamTranscodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamTranscodeResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamTranscodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamTranscodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamTranscodeResponse) GetBody() *DeleteLiveStreamTranscodeResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamTranscodeResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamTranscodeResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamTranscodeResponse) SetStatusCode(v int32) *DeleteLiveStreamTranscodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamTranscodeResponse) SetBody(v *DeleteLiveStreamTranscodeResponseBody) *DeleteLiveStreamTranscodeResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamTranscodeResponse) Validate() error {
	return dara.Validate(s)
}

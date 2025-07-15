// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWebRtcInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveWebRtcInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveWebRtcInfoResponse
	GetStatusCode() *int32
	SetBody(v *SaveWebRtcInfoResponseBody) *SaveWebRtcInfoResponse
	GetBody() *SaveWebRtcInfoResponseBody
}

type SaveWebRtcInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWebRtcInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveWebRtcInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveWebRtcInfoResponse) GoString() string {
	return s.String()
}

func (s *SaveWebRtcInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveWebRtcInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveWebRtcInfoResponse) GetBody() *SaveWebRtcInfoResponseBody {
	return s.Body
}

func (s *SaveWebRtcInfoResponse) SetHeaders(v map[string]*string) *SaveWebRtcInfoResponse {
	s.Headers = v
	return s
}

func (s *SaveWebRtcInfoResponse) SetStatusCode(v int32) *SaveWebRtcInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWebRtcInfoResponse) SetBody(v *SaveWebRtcInfoResponseBody) *SaveWebRtcInfoResponse {
	s.Body = v
	return s
}

func (s *SaveWebRtcInfoResponse) Validate() error {
	return dara.Validate(s)
}

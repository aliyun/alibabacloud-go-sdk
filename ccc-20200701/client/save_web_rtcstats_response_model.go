// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveWebRTCStatsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SaveWebRTCStatsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SaveWebRTCStatsResponse
	GetStatusCode() *int32
	SetBody(v *SaveWebRTCStatsResponseBody) *SaveWebRTCStatsResponse
	GetBody() *SaveWebRTCStatsResponseBody
}

type SaveWebRTCStatsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SaveWebRTCStatsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SaveWebRTCStatsResponse) String() string {
	return dara.Prettify(s)
}

func (s SaveWebRTCStatsResponse) GoString() string {
	return s.String()
}

func (s *SaveWebRTCStatsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SaveWebRTCStatsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SaveWebRTCStatsResponse) GetBody() *SaveWebRTCStatsResponseBody {
	return s.Body
}

func (s *SaveWebRTCStatsResponse) SetHeaders(v map[string]*string) *SaveWebRTCStatsResponse {
	s.Headers = v
	return s
}

func (s *SaveWebRTCStatsResponse) SetStatusCode(v int32) *SaveWebRTCStatsResponse {
	s.StatusCode = &v
	return s
}

func (s *SaveWebRTCStatsResponse) SetBody(v *SaveWebRTCStatsResponseBody) *SaveWebRTCStatsResponse {
	s.Body = v
	return s
}

func (s *SaveWebRTCStatsResponse) Validate() error {
	return dara.Validate(s)
}

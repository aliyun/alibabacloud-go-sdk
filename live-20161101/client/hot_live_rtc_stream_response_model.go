// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHotLiveRtcStreamResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *HotLiveRtcStreamResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *HotLiveRtcStreamResponse
	GetStatusCode() *int32
	SetBody(v *HotLiveRtcStreamResponseBody) *HotLiveRtcStreamResponse
	GetBody() *HotLiveRtcStreamResponseBody
}

type HotLiveRtcStreamResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *HotLiveRtcStreamResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s HotLiveRtcStreamResponse) String() string {
	return dara.Prettify(s)
}

func (s HotLiveRtcStreamResponse) GoString() string {
	return s.String()
}

func (s *HotLiveRtcStreamResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *HotLiveRtcStreamResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *HotLiveRtcStreamResponse) GetBody() *HotLiveRtcStreamResponseBody {
	return s.Body
}

func (s *HotLiveRtcStreamResponse) SetHeaders(v map[string]*string) *HotLiveRtcStreamResponse {
	s.Headers = v
	return s
}

func (s *HotLiveRtcStreamResponse) SetStatusCode(v int32) *HotLiveRtcStreamResponse {
	s.StatusCode = &v
	return s
}

func (s *HotLiveRtcStreamResponse) SetBody(v *HotLiveRtcStreamResponseBody) *HotLiveRtcStreamResponse {
	s.Body = v
	return s
}

func (s *HotLiveRtcStreamResponse) Validate() error {
	return dara.Validate(s)
}

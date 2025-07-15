// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveStreamWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveStreamWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveStreamWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveStreamWatermarkResponseBody) *AddLiveStreamWatermarkResponse
	GetBody() *AddLiveStreamWatermarkResponseBody
}

type AddLiveStreamWatermarkResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveStreamWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveStreamWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveStreamWatermarkResponse) GoString() string {
	return s.String()
}

func (s *AddLiveStreamWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveStreamWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveStreamWatermarkResponse) GetBody() *AddLiveStreamWatermarkResponseBody {
	return s.Body
}

func (s *AddLiveStreamWatermarkResponse) SetHeaders(v map[string]*string) *AddLiveStreamWatermarkResponse {
	s.Headers = v
	return s
}

func (s *AddLiveStreamWatermarkResponse) SetStatusCode(v int32) *AddLiveStreamWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveStreamWatermarkResponse) SetBody(v *AddLiveStreamWatermarkResponseBody) *AddLiveStreamWatermarkResponse {
	s.Body = v
	return s
}

func (s *AddLiveStreamWatermarkResponse) Validate() error {
	return dara.Validate(s)
}

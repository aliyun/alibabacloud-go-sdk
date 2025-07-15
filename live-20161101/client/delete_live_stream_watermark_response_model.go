// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveStreamWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveStreamWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveStreamWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveStreamWatermarkResponseBody) *DeleteLiveStreamWatermarkResponse
	GetBody() *DeleteLiveStreamWatermarkResponseBody
}

type DeleteLiveStreamWatermarkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveStreamWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveStreamWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveStreamWatermarkResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveStreamWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveStreamWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveStreamWatermarkResponse) GetBody() *DeleteLiveStreamWatermarkResponseBody {
	return s.Body
}

func (s *DeleteLiveStreamWatermarkResponse) SetHeaders(v map[string]*string) *DeleteLiveStreamWatermarkResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveStreamWatermarkResponse) SetStatusCode(v int32) *DeleteLiveStreamWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveStreamWatermarkResponse) SetBody(v *DeleteLiveStreamWatermarkResponseBody) *DeleteLiveStreamWatermarkResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveStreamWatermarkResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveStreamWatermarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveStreamWatermarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveStreamWatermarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveStreamWatermarkResponseBody) *UpdateLiveStreamWatermarkResponse
	GetBody() *UpdateLiveStreamWatermarkResponseBody
}

type UpdateLiveStreamWatermarkResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveStreamWatermarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveStreamWatermarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveStreamWatermarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveStreamWatermarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveStreamWatermarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveStreamWatermarkResponse) GetBody() *UpdateLiveStreamWatermarkResponseBody {
	return s.Body
}

func (s *UpdateLiveStreamWatermarkResponse) SetHeaders(v map[string]*string) *UpdateLiveStreamWatermarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveStreamWatermarkResponse) SetStatusCode(v int32) *UpdateLiveStreamWatermarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveStreamWatermarkResponse) SetBody(v *UpdateLiveStreamWatermarkResponseBody) *UpdateLiveStreamWatermarkResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveStreamWatermarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

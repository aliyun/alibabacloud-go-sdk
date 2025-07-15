// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAudioAuditConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveAudioAuditConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveAudioAuditConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveAudioAuditConfigResponseBody) *UpdateLiveAudioAuditConfigResponse
	GetBody() *UpdateLiveAudioAuditConfigResponseBody
}

type UpdateLiveAudioAuditConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveAudioAuditConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveAudioAuditConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveAudioAuditConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveAudioAuditConfigResponse) GetBody() *UpdateLiveAudioAuditConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveAudioAuditConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveAudioAuditConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveAudioAuditConfigResponse) SetStatusCode(v int32) *UpdateLiveAudioAuditConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveAudioAuditConfigResponse) SetBody(v *UpdateLiveAudioAuditConfigResponseBody) *UpdateLiveAudioAuditConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveAudioAuditConfigResponse) Validate() error {
	return dara.Validate(s)
}

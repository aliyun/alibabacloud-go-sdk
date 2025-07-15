// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLiveAudioAuditNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLiveAudioAuditNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLiveAudioAuditNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLiveAudioAuditNotifyConfigResponseBody) *UpdateLiveAudioAuditNotifyConfigResponse
	GetBody() *UpdateLiveAudioAuditNotifyConfigResponseBody
}

type UpdateLiveAudioAuditNotifyConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLiveAudioAuditNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLiveAudioAuditNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) GetBody() *UpdateLiveAudioAuditNotifyConfigResponseBody {
	return s.Body
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) SetHeaders(v map[string]*string) *UpdateLiveAudioAuditNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) SetStatusCode(v int32) *UpdateLiveAudioAuditNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) SetBody(v *UpdateLiveAudioAuditNotifyConfigResponseBody) *UpdateLiveAudioAuditNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateLiveAudioAuditNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}

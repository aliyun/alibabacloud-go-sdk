// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAudioAuditNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAudioAuditNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAudioAuditNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAudioAuditNotifyConfigResponseBody) *DeleteLiveAudioAuditNotifyConfigResponse
	GetBody() *DeleteLiveAudioAuditNotifyConfigResponseBody
}

type DeleteLiveAudioAuditNotifyConfigResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAudioAuditNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAudioAuditNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) GetBody() *DeleteLiveAudioAuditNotifyConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveAudioAuditNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) SetStatusCode(v int32) *DeleteLiveAudioAuditNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) SetBody(v *DeleteLiveAudioAuditNotifyConfigResponseBody) *DeleteLiveAudioAuditNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAudioAuditNotifyConfigResponse) Validate() error {
	return dara.Validate(s)
}

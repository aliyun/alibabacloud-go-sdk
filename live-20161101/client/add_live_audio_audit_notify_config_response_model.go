// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAudioAuditNotifyConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveAudioAuditNotifyConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveAudioAuditNotifyConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveAudioAuditNotifyConfigResponseBody) *AddLiveAudioAuditNotifyConfigResponse
	GetBody() *AddLiveAudioAuditNotifyConfigResponseBody
}

type AddLiveAudioAuditNotifyConfigResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveAudioAuditNotifyConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveAudioAuditNotifyConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAudioAuditNotifyConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditNotifyConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveAudioAuditNotifyConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveAudioAuditNotifyConfigResponse) GetBody() *AddLiveAudioAuditNotifyConfigResponseBody {
	return s.Body
}

func (s *AddLiveAudioAuditNotifyConfigResponse) SetHeaders(v map[string]*string) *AddLiveAudioAuditNotifyConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigResponse) SetStatusCode(v int32) *AddLiveAudioAuditNotifyConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigResponse) SetBody(v *AddLiveAudioAuditNotifyConfigResponseBody) *AddLiveAudioAuditNotifyConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveAudioAuditNotifyConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

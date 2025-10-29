// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveAudioAuditConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLiveAudioAuditConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLiveAudioAuditConfigResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLiveAudioAuditConfigResponseBody) *DeleteLiveAudioAuditConfigResponse
	GetBody() *DeleteLiveAudioAuditConfigResponseBody
}

type DeleteLiveAudioAuditConfigResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLiveAudioAuditConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLiveAudioAuditConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *DeleteLiveAudioAuditConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLiveAudioAuditConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLiveAudioAuditConfigResponse) GetBody() *DeleteLiveAudioAuditConfigResponseBody {
	return s.Body
}

func (s *DeleteLiveAudioAuditConfigResponse) SetHeaders(v map[string]*string) *DeleteLiveAudioAuditConfigResponse {
	s.Headers = v
	return s
}

func (s *DeleteLiveAudioAuditConfigResponse) SetStatusCode(v int32) *DeleteLiveAudioAuditConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLiveAudioAuditConfigResponse) SetBody(v *DeleteLiveAudioAuditConfigResponseBody) *DeleteLiveAudioAuditConfigResponse {
	s.Body = v
	return s
}

func (s *DeleteLiveAudioAuditConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

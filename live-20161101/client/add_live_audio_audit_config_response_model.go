// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddLiveAudioAuditConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddLiveAudioAuditConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddLiveAudioAuditConfigResponse
	GetStatusCode() *int32
	SetBody(v *AddLiveAudioAuditConfigResponseBody) *AddLiveAudioAuditConfigResponse
	GetBody() *AddLiveAudioAuditConfigResponseBody
}

type AddLiveAudioAuditConfigResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddLiveAudioAuditConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddLiveAudioAuditConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s AddLiveAudioAuditConfigResponse) GoString() string {
	return s.String()
}

func (s *AddLiveAudioAuditConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddLiveAudioAuditConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddLiveAudioAuditConfigResponse) GetBody() *AddLiveAudioAuditConfigResponseBody {
	return s.Body
}

func (s *AddLiveAudioAuditConfigResponse) SetHeaders(v map[string]*string) *AddLiveAudioAuditConfigResponse {
	s.Headers = v
	return s
}

func (s *AddLiveAudioAuditConfigResponse) SetStatusCode(v int32) *AddLiveAudioAuditConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *AddLiveAudioAuditConfigResponse) SetBody(v *AddLiveAudioAuditConfigResponseBody) *AddLiveAudioAuditConfigResponse {
	s.Body = v
	return s
}

func (s *AddLiveAudioAuditConfigResponse) Validate() error {
	return dara.Validate(s)
}

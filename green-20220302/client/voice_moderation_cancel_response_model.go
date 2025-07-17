// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVoiceModerationCancelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *VoiceModerationCancelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *VoiceModerationCancelResponse
	GetStatusCode() *int32
	SetBody(v *VoiceModerationCancelResponseBody) *VoiceModerationCancelResponse
	GetBody() *VoiceModerationCancelResponseBody
}

type VoiceModerationCancelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *VoiceModerationCancelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s VoiceModerationCancelResponse) String() string {
	return dara.Prettify(s)
}

func (s VoiceModerationCancelResponse) GoString() string {
	return s.String()
}

func (s *VoiceModerationCancelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *VoiceModerationCancelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *VoiceModerationCancelResponse) GetBody() *VoiceModerationCancelResponseBody {
	return s.Body
}

func (s *VoiceModerationCancelResponse) SetHeaders(v map[string]*string) *VoiceModerationCancelResponse {
	s.Headers = v
	return s
}

func (s *VoiceModerationCancelResponse) SetStatusCode(v int32) *VoiceModerationCancelResponse {
	s.StatusCode = &v
	return s
}

func (s *VoiceModerationCancelResponse) SetBody(v *VoiceModerationCancelResponseBody) *VoiceModerationCancelResponse {
	s.Body = v
	return s
}

func (s *VoiceModerationCancelResponse) Validate() error {
	return dara.Validate(s)
}

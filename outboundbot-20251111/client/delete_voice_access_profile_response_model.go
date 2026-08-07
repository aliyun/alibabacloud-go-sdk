// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVoiceAccessProfileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVoiceAccessProfileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVoiceAccessProfileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVoiceAccessProfileResponseBody) *DeleteVoiceAccessProfileResponse
	GetBody() *DeleteVoiceAccessProfileResponseBody
}

type DeleteVoiceAccessProfileResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVoiceAccessProfileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVoiceAccessProfileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVoiceAccessProfileResponse) GoString() string {
	return s.String()
}

func (s *DeleteVoiceAccessProfileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVoiceAccessProfileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVoiceAccessProfileResponse) GetBody() *DeleteVoiceAccessProfileResponseBody {
	return s.Body
}

func (s *DeleteVoiceAccessProfileResponse) SetHeaders(v map[string]*string) *DeleteVoiceAccessProfileResponse {
	s.Headers = v
	return s
}

func (s *DeleteVoiceAccessProfileResponse) SetStatusCode(v int32) *DeleteVoiceAccessProfileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVoiceAccessProfileResponse) SetBody(v *DeleteVoiceAccessProfileResponseBody) *DeleteVoiceAccessProfileResponse {
	s.Body = v
	return s
}

func (s *DeleteVoiceAccessProfileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

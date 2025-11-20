// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVideoConferenceSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateVideoConferenceSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateVideoConferenceSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateVideoConferenceSettingResponseBody) *UpdateVideoConferenceSettingResponse
	GetBody() *UpdateVideoConferenceSettingResponseBody
}

type UpdateVideoConferenceSettingResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateVideoConferenceSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateVideoConferenceSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateVideoConferenceSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateVideoConferenceSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateVideoConferenceSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateVideoConferenceSettingResponse) GetBody() *UpdateVideoConferenceSettingResponseBody {
	return s.Body
}

func (s *UpdateVideoConferenceSettingResponse) SetHeaders(v map[string]*string) *UpdateVideoConferenceSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateVideoConferenceSettingResponse) SetStatusCode(v int32) *UpdateVideoConferenceSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateVideoConferenceSettingResponse) SetBody(v *UpdateVideoConferenceSettingResponseBody) *UpdateVideoConferenceSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateVideoConferenceSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

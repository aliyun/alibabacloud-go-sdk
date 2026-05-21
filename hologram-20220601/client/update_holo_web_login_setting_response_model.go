// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHoloWebLoginSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHoloWebLoginSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHoloWebLoginSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHoloWebLoginSettingResponseBody) *UpdateHoloWebLoginSettingResponse
	GetBody() *UpdateHoloWebLoginSettingResponseBody
}

type UpdateHoloWebLoginSettingResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHoloWebLoginSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHoloWebLoginSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHoloWebLoginSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateHoloWebLoginSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHoloWebLoginSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHoloWebLoginSettingResponse) GetBody() *UpdateHoloWebLoginSettingResponseBody {
	return s.Body
}

func (s *UpdateHoloWebLoginSettingResponse) SetHeaders(v map[string]*string) *UpdateHoloWebLoginSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateHoloWebLoginSettingResponse) SetStatusCode(v int32) *UpdateHoloWebLoginSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHoloWebLoginSettingResponse) SetBody(v *UpdateHoloWebLoginSettingResponseBody) *UpdateHoloWebLoginSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateHoloWebLoginSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

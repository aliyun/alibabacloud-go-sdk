// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoloWebLoginSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHoloWebLoginSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHoloWebLoginSettingResponse
	GetStatusCode() *int32
	SetBody(v *GetHoloWebLoginSettingResponseBody) *GetHoloWebLoginSettingResponse
	GetBody() *GetHoloWebLoginSettingResponseBody
}

type GetHoloWebLoginSettingResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHoloWebLoginSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHoloWebLoginSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHoloWebLoginSettingResponse) GoString() string {
	return s.String()
}

func (s *GetHoloWebLoginSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHoloWebLoginSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHoloWebLoginSettingResponse) GetBody() *GetHoloWebLoginSettingResponseBody {
	return s.Body
}

func (s *GetHoloWebLoginSettingResponse) SetHeaders(v map[string]*string) *GetHoloWebLoginSettingResponse {
	s.Headers = v
	return s
}

func (s *GetHoloWebLoginSettingResponse) SetStatusCode(v int32) *GetHoloWebLoginSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHoloWebLoginSettingResponse) SetBody(v *GetHoloWebLoginSettingResponseBody) *GetHoloWebLoginSettingResponse {
	s.Body = v
	return s
}

func (s *GetHoloWebLoginSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

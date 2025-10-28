// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLocalitySettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLocalitySettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLocalitySettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLocalitySettingResponseBody) *UpdateLocalitySettingResponse
	GetBody() *UpdateLocalitySettingResponseBody
}

type UpdateLocalitySettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLocalitySettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLocalitySettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLocalitySettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateLocalitySettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLocalitySettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLocalitySettingResponse) GetBody() *UpdateLocalitySettingResponseBody {
	return s.Body
}

func (s *UpdateLocalitySettingResponse) SetHeaders(v map[string]*string) *UpdateLocalitySettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateLocalitySettingResponse) SetStatusCode(v int32) *UpdateLocalitySettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLocalitySettingResponse) SetBody(v *UpdateLocalitySettingResponseBody) *UpdateLocalitySettingResponse {
	s.Body = v
	return s
}

func (s *UpdateLocalitySettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

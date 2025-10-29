// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAdvancedSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateAdvancedSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateAdvancedSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateAdvancedSettingResponseBody) *UpdateAdvancedSettingResponse
	GetBody() *UpdateAdvancedSettingResponseBody
}

type UpdateAdvancedSettingResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateAdvancedSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateAdvancedSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateAdvancedSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateAdvancedSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateAdvancedSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateAdvancedSettingResponse) GetBody() *UpdateAdvancedSettingResponseBody {
	return s.Body
}

func (s *UpdateAdvancedSettingResponse) SetHeaders(v map[string]*string) *UpdateAdvancedSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateAdvancedSettingResponse) SetStatusCode(v int32) *UpdateAdvancedSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateAdvancedSettingResponse) SetBody(v *UpdateAdvancedSettingResponseBody) *UpdateAdvancedSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateAdvancedSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

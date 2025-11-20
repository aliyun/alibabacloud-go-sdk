// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaRetentionSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHanaRetentionSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHanaRetentionSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHanaRetentionSettingResponseBody) *UpdateHanaRetentionSettingResponse
	GetBody() *UpdateHanaRetentionSettingResponseBody
}

type UpdateHanaRetentionSettingResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHanaRetentionSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHanaRetentionSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaRetentionSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaRetentionSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHanaRetentionSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHanaRetentionSettingResponse) GetBody() *UpdateHanaRetentionSettingResponseBody {
	return s.Body
}

func (s *UpdateHanaRetentionSettingResponse) SetHeaders(v map[string]*string) *UpdateHanaRetentionSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaRetentionSettingResponse) SetStatusCode(v int32) *UpdateHanaRetentionSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaRetentionSettingResponse) SetBody(v *UpdateHanaRetentionSettingResponseBody) *UpdateHanaRetentionSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateHanaRetentionSettingResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

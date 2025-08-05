// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaBackupSettingResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHanaBackupSettingResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHanaBackupSettingResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHanaBackupSettingResponseBody) *UpdateHanaBackupSettingResponse
	GetBody() *UpdateHanaBackupSettingResponseBody
}

type UpdateHanaBackupSettingResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHanaBackupSettingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHanaBackupSettingResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaBackupSettingResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupSettingResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHanaBackupSettingResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHanaBackupSettingResponse) GetBody() *UpdateHanaBackupSettingResponseBody {
	return s.Body
}

func (s *UpdateHanaBackupSettingResponse) SetHeaders(v map[string]*string) *UpdateHanaBackupSettingResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaBackupSettingResponse) SetStatusCode(v int32) *UpdateHanaBackupSettingResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaBackupSettingResponse) SetBody(v *UpdateHanaBackupSettingResponseBody) *UpdateHanaBackupSettingResponse {
	s.Body = v
	return s
}

func (s *UpdateHanaBackupSettingResponse) Validate() error {
	return dara.Validate(s)
}

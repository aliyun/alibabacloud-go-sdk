// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBackupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBackupConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBackupConfigResponseBody) *UpdateBackupConfigResponse
	GetBody() *UpdateBackupConfigResponseBody
}

type UpdateBackupConfigResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBackupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBackupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBackupConfigResponse) GetBody() *UpdateBackupConfigResponseBody {
	return s.Body
}

func (s *UpdateBackupConfigResponse) SetHeaders(v map[string]*string) *UpdateBackupConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupConfigResponse) SetStatusCode(v int32) *UpdateBackupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupConfigResponse) SetBody(v *UpdateBackupConfigResponseBody) *UpdateBackupConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateBackupConfigResponse) Validate() error {
	return dara.Validate(s)
}

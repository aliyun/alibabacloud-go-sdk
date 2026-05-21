// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateScheduledBackupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateScheduledBackupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateScheduledBackupConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateScheduledBackupConfigResponseBody) *UpdateScheduledBackupConfigResponse
	GetBody() *UpdateScheduledBackupConfigResponseBody
}

type UpdateScheduledBackupConfigResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateScheduledBackupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateScheduledBackupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateScheduledBackupConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateScheduledBackupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateScheduledBackupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateScheduledBackupConfigResponse) GetBody() *UpdateScheduledBackupConfigResponseBody {
	return s.Body
}

func (s *UpdateScheduledBackupConfigResponse) SetHeaders(v map[string]*string) *UpdateScheduledBackupConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateScheduledBackupConfigResponse) SetStatusCode(v int32) *UpdateScheduledBackupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateScheduledBackupConfigResponse) SetBody(v *UpdateScheduledBackupConfigResponseBody) *UpdateScheduledBackupConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateScheduledBackupConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

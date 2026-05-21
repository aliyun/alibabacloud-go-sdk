// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupDataDescResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBackupDataDescResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBackupDataDescResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBackupDataDescResponseBody) *UpdateBackupDataDescResponse
	GetBody() *UpdateBackupDataDescResponseBody
}

type UpdateBackupDataDescResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupDataDescResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBackupDataDescResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupDataDescResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupDataDescResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBackupDataDescResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBackupDataDescResponse) GetBody() *UpdateBackupDataDescResponseBody {
	return s.Body
}

func (s *UpdateBackupDataDescResponse) SetHeaders(v map[string]*string) *UpdateBackupDataDescResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupDataDescResponse) SetStatusCode(v int32) *UpdateBackupDataDescResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupDataDescResponse) SetBody(v *UpdateBackupDataDescResponseBody) *UpdateBackupDataDescResponse {
	s.Body = v
	return s
}

func (s *UpdateBackupDataDescResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

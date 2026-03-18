// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBackupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBackupResponseBody) *UpdateBackupResponse
	GetBody() *UpdateBackupResponseBody
}

type UpdateBackupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBackupResponse) GetBody() *UpdateBackupResponseBody {
	return s.Body
}

func (s *UpdateBackupResponse) SetHeaders(v map[string]*string) *UpdateBackupResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupResponse) SetStatusCode(v int32) *UpdateBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupResponse) SetBody(v *UpdateBackupResponseBody) *UpdateBackupResponse {
	s.Body = v
	return s
}

func (s *UpdateBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

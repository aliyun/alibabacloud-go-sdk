// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureBackupSetInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureBackupSetInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureBackupSetInfoResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureBackupSetInfoResponseBody) *ConfigureBackupSetInfoResponse
	GetBody() *ConfigureBackupSetInfoResponseBody
}

type ConfigureBackupSetInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureBackupSetInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureBackupSetInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupSetInfoResponse) GoString() string {
	return s.String()
}

func (s *ConfigureBackupSetInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureBackupSetInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureBackupSetInfoResponse) GetBody() *ConfigureBackupSetInfoResponseBody {
	return s.Body
}

func (s *ConfigureBackupSetInfoResponse) SetHeaders(v map[string]*string) *ConfigureBackupSetInfoResponse {
	s.Headers = v
	return s
}

func (s *ConfigureBackupSetInfoResponse) SetStatusCode(v int32) *ConfigureBackupSetInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureBackupSetInfoResponse) SetBody(v *ConfigureBackupSetInfoResponseBody) *ConfigureBackupSetInfoResponse {
	s.Body = v
	return s
}

func (s *ConfigureBackupSetInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

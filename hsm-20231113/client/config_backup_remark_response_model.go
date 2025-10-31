// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigBackupRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigBackupRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigBackupRemarkResponse
	GetStatusCode() *int32
	SetBody(v *ConfigBackupRemarkResponseBody) *ConfigBackupRemarkResponse
	GetBody() *ConfigBackupRemarkResponseBody
}

type ConfigBackupRemarkResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigBackupRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigBackupRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigBackupRemarkResponse) GoString() string {
	return s.String()
}

func (s *ConfigBackupRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigBackupRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigBackupRemarkResponse) GetBody() *ConfigBackupRemarkResponseBody {
	return s.Body
}

func (s *ConfigBackupRemarkResponse) SetHeaders(v map[string]*string) *ConfigBackupRemarkResponse {
	s.Headers = v
	return s
}

func (s *ConfigBackupRemarkResponse) SetStatusCode(v int32) *ConfigBackupRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigBackupRemarkResponse) SetBody(v *ConfigBackupRemarkResponseBody) *ConfigBackupRemarkResponse {
	s.Body = v
	return s
}

func (s *ConfigBackupRemarkResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

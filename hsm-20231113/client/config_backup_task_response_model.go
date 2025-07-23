// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigBackupTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigBackupTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigBackupTaskResponse
	GetStatusCode() *int32
	SetBody(v *ConfigBackupTaskResponseBody) *ConfigBackupTaskResponse
	GetBody() *ConfigBackupTaskResponseBody
}

type ConfigBackupTaskResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigBackupTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigBackupTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigBackupTaskResponse) GoString() string {
	return s.String()
}

func (s *ConfigBackupTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigBackupTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigBackupTaskResponse) GetBody() *ConfigBackupTaskResponseBody {
	return s.Body
}

func (s *ConfigBackupTaskResponse) SetHeaders(v map[string]*string) *ConfigBackupTaskResponse {
	s.Headers = v
	return s
}

func (s *ConfigBackupTaskResponse) SetStatusCode(v int32) *ConfigBackupTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigBackupTaskResponse) SetBody(v *ConfigBackupTaskResponseBody) *ConfigBackupTaskResponse {
	s.Body = v
	return s
}

func (s *ConfigBackupTaskResponse) Validate() error {
	return dara.Validate(s)
}

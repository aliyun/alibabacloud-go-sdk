// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupExpireTimeResponseBody) *ModifyBackupExpireTimeResponse
	GetBody() *ModifyBackupExpireTimeResponseBody
}

type ModifyBackupExpireTimeResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupExpireTimeResponse) GetBody() *ModifyBackupExpireTimeResponseBody {
	return s.Body
}

func (s *ModifyBackupExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyBackupExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupExpireTimeResponse) SetStatusCode(v int32) *ModifyBackupExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupExpireTimeResponse) SetBody(v *ModifyBackupExpireTimeResponseBody) *ModifyBackupExpireTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupExpireTimeResponse) Validate() error {
	return dara.Validate(s)
}

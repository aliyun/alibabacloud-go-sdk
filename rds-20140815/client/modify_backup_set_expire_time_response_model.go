// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupSetExpireTimeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupSetExpireTimeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupSetExpireTimeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupSetExpireTimeResponseBody) *ModifyBackupSetExpireTimeResponse
	GetBody() *ModifyBackupSetExpireTimeResponseBody
}

type ModifyBackupSetExpireTimeResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupSetExpireTimeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupSetExpireTimeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupSetExpireTimeResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupSetExpireTimeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupSetExpireTimeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupSetExpireTimeResponse) GetBody() *ModifyBackupSetExpireTimeResponseBody {
	return s.Body
}

func (s *ModifyBackupSetExpireTimeResponse) SetHeaders(v map[string]*string) *ModifyBackupSetExpireTimeResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupSetExpireTimeResponse) SetStatusCode(v int32) *ModifyBackupSetExpireTimeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupSetExpireTimeResponse) SetBody(v *ModifyBackupSetExpireTimeResponseBody) *ModifyBackupSetExpireTimeResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupSetExpireTimeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

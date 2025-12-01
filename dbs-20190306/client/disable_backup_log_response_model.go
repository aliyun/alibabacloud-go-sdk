// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupLogResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableBackupLogResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableBackupLogResponse
	GetStatusCode() *int32
	SetBody(v *DisableBackupLogResponseBody) *DisableBackupLogResponse
	GetBody() *DisableBackupLogResponseBody
}

type DisableBackupLogResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableBackupLogResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableBackupLogResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupLogResponse) GoString() string {
	return s.String()
}

func (s *DisableBackupLogResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableBackupLogResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableBackupLogResponse) GetBody() *DisableBackupLogResponseBody {
	return s.Body
}

func (s *DisableBackupLogResponse) SetHeaders(v map[string]*string) *DisableBackupLogResponse {
	s.Headers = v
	return s
}

func (s *DisableBackupLogResponse) SetStatusCode(v int32) *DisableBackupLogResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableBackupLogResponse) SetBody(v *DisableBackupLogResponseBody) *DisableBackupLogResponse {
	s.Body = v
	return s
}

func (s *DisableBackupLogResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

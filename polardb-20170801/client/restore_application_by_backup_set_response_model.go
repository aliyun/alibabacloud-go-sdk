// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestoreApplicationByBackupSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestoreApplicationByBackupSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestoreApplicationByBackupSetResponse
	GetStatusCode() *int32
	SetBody(v *RestoreApplicationByBackupSetResponseBody) *RestoreApplicationByBackupSetResponse
	GetBody() *RestoreApplicationByBackupSetResponseBody
}

type RestoreApplicationByBackupSetResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreApplicationByBackupSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreApplicationByBackupSetResponse) String() string {
	return dara.Prettify(s)
}

func (s RestoreApplicationByBackupSetResponse) GoString() string {
	return s.String()
}

func (s *RestoreApplicationByBackupSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestoreApplicationByBackupSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestoreApplicationByBackupSetResponse) GetBody() *RestoreApplicationByBackupSetResponseBody {
	return s.Body
}

func (s *RestoreApplicationByBackupSetResponse) SetHeaders(v map[string]*string) *RestoreApplicationByBackupSetResponse {
	s.Headers = v
	return s
}

func (s *RestoreApplicationByBackupSetResponse) SetStatusCode(v int32) *RestoreApplicationByBackupSetResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreApplicationByBackupSetResponse) SetBody(v *RestoreApplicationByBackupSetResponseBody) *RestoreApplicationByBackupSetResponse {
	s.Body = v
	return s
}

func (s *RestoreApplicationByBackupSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

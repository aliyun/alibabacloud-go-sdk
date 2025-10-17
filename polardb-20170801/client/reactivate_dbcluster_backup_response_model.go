// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReactivateDBClusterBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReactivateDBClusterBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReactivateDBClusterBackupResponse
	GetStatusCode() *int32
	SetBody(v *ReactivateDBClusterBackupResponseBody) *ReactivateDBClusterBackupResponse
	GetBody() *ReactivateDBClusterBackupResponseBody
}

type ReactivateDBClusterBackupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReactivateDBClusterBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReactivateDBClusterBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s ReactivateDBClusterBackupResponse) GoString() string {
	return s.String()
}

func (s *ReactivateDBClusterBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReactivateDBClusterBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReactivateDBClusterBackupResponse) GetBody() *ReactivateDBClusterBackupResponseBody {
	return s.Body
}

func (s *ReactivateDBClusterBackupResponse) SetHeaders(v map[string]*string) *ReactivateDBClusterBackupResponse {
	s.Headers = v
	return s
}

func (s *ReactivateDBClusterBackupResponse) SetStatusCode(v int32) *ReactivateDBClusterBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *ReactivateDBClusterBackupResponse) SetBody(v *ReactivateDBClusterBackupResponseBody) *ReactivateDBClusterBackupResponse {
	s.Body = v
	return s
}

func (s *ReactivateDBClusterBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUseAclBackupDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UseAclBackupDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UseAclBackupDataResponse
	GetStatusCode() *int32
	SetBody(v *UseAclBackupDataResponseBody) *UseAclBackupDataResponse
	GetBody() *UseAclBackupDataResponseBody
}

type UseAclBackupDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UseAclBackupDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UseAclBackupDataResponse) String() string {
	return dara.Prettify(s)
}

func (s UseAclBackupDataResponse) GoString() string {
	return s.String()
}

func (s *UseAclBackupDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UseAclBackupDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UseAclBackupDataResponse) GetBody() *UseAclBackupDataResponseBody {
	return s.Body
}

func (s *UseAclBackupDataResponse) SetHeaders(v map[string]*string) *UseAclBackupDataResponse {
	s.Headers = v
	return s
}

func (s *UseAclBackupDataResponse) SetStatusCode(v int32) *UseAclBackupDataResponse {
	s.StatusCode = &v
	return s
}

func (s *UseAclBackupDataResponse) SetBody(v *UseAclBackupDataResponseBody) *UseAclBackupDataResponse {
	s.Body = v
	return s
}

func (s *UseAclBackupDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

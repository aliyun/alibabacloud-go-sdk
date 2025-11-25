// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAclBackupDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAclBackupDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAclBackupDataResponse
	GetStatusCode() *int32
	SetBody(v *AddAclBackupDataResponseBody) *AddAclBackupDataResponse
	GetBody() *AddAclBackupDataResponseBody
}

type AddAclBackupDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAclBackupDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAclBackupDataResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAclBackupDataResponse) GoString() string {
	return s.String()
}

func (s *AddAclBackupDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAclBackupDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAclBackupDataResponse) GetBody() *AddAclBackupDataResponseBody {
	return s.Body
}

func (s *AddAclBackupDataResponse) SetHeaders(v map[string]*string) *AddAclBackupDataResponse {
	s.Headers = v
	return s
}

func (s *AddAclBackupDataResponse) SetStatusCode(v int32) *AddAclBackupDataResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAclBackupDataResponse) SetBody(v *AddAclBackupDataResponseBody) *AddAclBackupDataResponse {
	s.Body = v
	return s
}

func (s *AddAclBackupDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

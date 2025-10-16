// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAclBackupDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAclBackupDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAclBackupDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAclBackupDataResponseBody) *DeleteAclBackupDataResponse
	GetBody() *DeleteAclBackupDataResponseBody
}

type DeleteAclBackupDataResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAclBackupDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAclBackupDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAclBackupDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteAclBackupDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAclBackupDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAclBackupDataResponse) GetBody() *DeleteAclBackupDataResponseBody {
	return s.Body
}

func (s *DeleteAclBackupDataResponse) SetHeaders(v map[string]*string) *DeleteAclBackupDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteAclBackupDataResponse) SetStatusCode(v int32) *DeleteAclBackupDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAclBackupDataResponse) SetBody(v *DeleteAclBackupDataResponseBody) *DeleteAclBackupDataResponse {
	s.Body = v
	return s
}

func (s *DeleteAclBackupDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

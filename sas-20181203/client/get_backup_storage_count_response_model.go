// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupStorageCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBackupStorageCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBackupStorageCountResponse
	GetStatusCode() *int32
	SetBody(v *GetBackupStorageCountResponseBody) *GetBackupStorageCountResponse
	GetBody() *GetBackupStorageCountResponseBody
}

type GetBackupStorageCountResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupStorageCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupStorageCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBackupStorageCountResponse) GoString() string {
	return s.String()
}

func (s *GetBackupStorageCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBackupStorageCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBackupStorageCountResponse) GetBody() *GetBackupStorageCountResponseBody {
	return s.Body
}

func (s *GetBackupStorageCountResponse) SetHeaders(v map[string]*string) *GetBackupStorageCountResponse {
	s.Headers = v
	return s
}

func (s *GetBackupStorageCountResponse) SetStatusCode(v int32) *GetBackupStorageCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupStorageCountResponse) SetBody(v *GetBackupStorageCountResponseBody) *GetBackupStorageCountResponse {
	s.Body = v
	return s
}

func (s *GetBackupStorageCountResponse) Validate() error {
	return dara.Validate(s)
}

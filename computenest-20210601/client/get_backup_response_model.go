// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBackupResponse
	GetStatusCode() *int32
	SetBody(v *GetBackupResponseBody) *GetBackupResponse
	GetBody() *GetBackupResponseBody
}

type GetBackupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBackupResponse) GoString() string {
	return s.String()
}

func (s *GetBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBackupResponse) GetBody() *GetBackupResponseBody {
	return s.Body
}

func (s *GetBackupResponse) SetHeaders(v map[string]*string) *GetBackupResponse {
	s.Headers = v
	return s
}

func (s *GetBackupResponse) SetStatusCode(v int32) *GetBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupResponse) SetBody(v *GetBackupResponseBody) *GetBackupResponse {
	s.Body = v
	return s
}

func (s *GetBackupResponse) Validate() error {
	return dara.Validate(s)
}

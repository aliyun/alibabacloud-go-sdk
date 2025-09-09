// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetBackupLocalResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetBackupLocalResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetBackupLocalResponse
	GetStatusCode() *int32
	SetBody(v *SetBackupLocalResponseBody) *SetBackupLocalResponse
	GetBody() *SetBackupLocalResponseBody
}

type SetBackupLocalResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetBackupLocalResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetBackupLocalResponse) String() string {
	return dara.Prettify(s)
}

func (s SetBackupLocalResponse) GoString() string {
	return s.String()
}

func (s *SetBackupLocalResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetBackupLocalResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetBackupLocalResponse) GetBody() *SetBackupLocalResponseBody {
	return s.Body
}

func (s *SetBackupLocalResponse) SetHeaders(v map[string]*string) *SetBackupLocalResponse {
	s.Headers = v
	return s
}

func (s *SetBackupLocalResponse) SetStatusCode(v int32) *SetBackupLocalResponse {
	s.StatusCode = &v
	return s
}

func (s *SetBackupLocalResponse) SetBody(v *SetBackupLocalResponseBody) *SetBackupLocalResponse {
	s.Body = v
	return s
}

func (s *SetBackupLocalResponse) Validate() error {
	return dara.Validate(s)
}

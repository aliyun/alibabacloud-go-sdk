// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackupResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackupResponseBody) *CreateBackupResponse
	GetBody() *CreateBackupResponseBody
}

type CreateBackupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackupResponse) GetBody() *CreateBackupResponseBody {
	return s.Body
}

func (s *CreateBackupResponse) SetHeaders(v map[string]*string) *CreateBackupResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupResponse) SetStatusCode(v int32) *CreateBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupResponse) SetBody(v *CreateBackupResponseBody) *CreateBackupResponse {
	s.Body = v
	return s
}

func (s *CreateBackupResponse) Validate() error {
	return dara.Validate(s)
}

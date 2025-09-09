// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPutStartBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PutStartBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PutStartBackupResponse
	GetStatusCode() *int32
	SetBody(v *PutStartBackupResponseBody) *PutStartBackupResponse
	GetBody() *PutStartBackupResponseBody
}

type PutStartBackupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PutStartBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PutStartBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s PutStartBackupResponse) GoString() string {
	return s.String()
}

func (s *PutStartBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PutStartBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PutStartBackupResponse) GetBody() *PutStartBackupResponseBody {
	return s.Body
}

func (s *PutStartBackupResponse) SetHeaders(v map[string]*string) *PutStartBackupResponse {
	s.Headers = v
	return s
}

func (s *PutStartBackupResponse) SetStatusCode(v int32) *PutStartBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *PutStartBackupResponse) SetBody(v *PutStartBackupResponseBody) *PutStartBackupResponse {
	s.Body = v
	return s
}

func (s *PutStartBackupResponse) Validate() error {
	return dara.Validate(s)
}

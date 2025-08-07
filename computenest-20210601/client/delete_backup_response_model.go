// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupResponseBody) *DeleteBackupResponse
	GetBody() *DeleteBackupResponseBody
}

type DeleteBackupResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupResponse) GetBody() *DeleteBackupResponseBody {
	return s.Body
}

func (s *DeleteBackupResponse) SetHeaders(v map[string]*string) *DeleteBackupResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupResponse) SetStatusCode(v int32) *DeleteBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupResponse) SetBody(v *DeleteBackupResponseBody) *DeleteBackupResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupClientResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupClientResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupClientResponseBody) *DeleteBackupClientResponse
	GetBody() *DeleteBackupClientResponseBody
}

type DeleteBackupClientResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupClientResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupClientResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupClientResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupClientResponse) GetBody() *DeleteBackupClientResponseBody {
	return s.Body
}

func (s *DeleteBackupClientResponse) SetHeaders(v map[string]*string) *DeleteBackupClientResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupClientResponse) SetStatusCode(v int32) *DeleteBackupClientResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupClientResponse) SetBody(v *DeleteBackupClientResponseBody) *DeleteBackupClientResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupClientResponse) Validate() error {
	return dara.Validate(s)
}

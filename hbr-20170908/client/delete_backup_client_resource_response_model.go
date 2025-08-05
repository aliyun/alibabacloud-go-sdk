// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupClientResourceResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupClientResourceResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupClientResourceResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupClientResourceResponseBody) *DeleteBackupClientResourceResponse
	GetBody() *DeleteBackupClientResourceResponseBody
}

type DeleteBackupClientResourceResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupClientResourceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupClientResourceResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupClientResourceResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupClientResourceResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupClientResourceResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupClientResourceResponse) GetBody() *DeleteBackupClientResourceResponseBody {
	return s.Body
}

func (s *DeleteBackupClientResourceResponse) SetHeaders(v map[string]*string) *DeleteBackupClientResourceResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupClientResourceResponse) SetStatusCode(v int32) *DeleteBackupClientResourceResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupClientResourceResponse) SetBody(v *DeleteBackupClientResourceResponseBody) *DeleteBackupClientResourceResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupClientResourceResponse) Validate() error {
	return dara.Validate(s)
}

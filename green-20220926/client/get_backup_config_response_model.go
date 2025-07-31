// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBackupConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBackupConfigResponse
	GetStatusCode() *int32
	SetBody(v *GetBackupConfigResponseBody) *GetBackupConfigResponse
	GetBody() *GetBackupConfigResponseBody
}

type GetBackupConfigResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBackupConfigResponse) GoString() string {
	return s.String()
}

func (s *GetBackupConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBackupConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBackupConfigResponse) GetBody() *GetBackupConfigResponseBody {
	return s.Body
}

func (s *GetBackupConfigResponse) SetHeaders(v map[string]*string) *GetBackupConfigResponse {
	s.Headers = v
	return s
}

func (s *GetBackupConfigResponse) SetStatusCode(v int32) *GetBackupConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupConfigResponse) SetBody(v *GetBackupConfigResponseBody) *GetBackupConfigResponse {
	s.Body = v
	return s
}

func (s *GetBackupConfigResponse) Validate() error {
	return dara.Validate(s)
}

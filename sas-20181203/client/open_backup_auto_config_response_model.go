// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupAutoConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenBackupAutoConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenBackupAutoConfigResponse
	GetStatusCode() *int32
	SetBody(v *OpenBackupAutoConfigResponseBody) *OpenBackupAutoConfigResponse
	GetBody() *OpenBackupAutoConfigResponseBody
}

type OpenBackupAutoConfigResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBackupAutoConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBackupAutoConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupAutoConfigResponse) GoString() string {
	return s.String()
}

func (s *OpenBackupAutoConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenBackupAutoConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenBackupAutoConfigResponse) GetBody() *OpenBackupAutoConfigResponseBody {
	return s.Body
}

func (s *OpenBackupAutoConfigResponse) SetHeaders(v map[string]*string) *OpenBackupAutoConfigResponse {
	s.Headers = v
	return s
}

func (s *OpenBackupAutoConfigResponse) SetStatusCode(v int32) *OpenBackupAutoConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBackupAutoConfigResponse) SetBody(v *OpenBackupAutoConfigResponseBody) *OpenBackupAutoConfigResponse {
	s.Body = v
	return s
}

func (s *OpenBackupAutoConfigResponse) Validate() error {
	return dara.Validate(s)
}

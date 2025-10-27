// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupAutoConfigStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBackupAutoConfigStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBackupAutoConfigStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetBackupAutoConfigStatusResponseBody) *GetBackupAutoConfigStatusResponse
	GetBody() *GetBackupAutoConfigStatusResponseBody
}

type GetBackupAutoConfigStatusResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupAutoConfigStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupAutoConfigStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBackupAutoConfigStatusResponse) GoString() string {
	return s.String()
}

func (s *GetBackupAutoConfigStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBackupAutoConfigStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBackupAutoConfigStatusResponse) GetBody() *GetBackupAutoConfigStatusResponseBody {
	return s.Body
}

func (s *GetBackupAutoConfigStatusResponse) SetHeaders(v map[string]*string) *GetBackupAutoConfigStatusResponse {
	s.Headers = v
	return s
}

func (s *GetBackupAutoConfigStatusResponse) SetStatusCode(v int32) *GetBackupAutoConfigStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupAutoConfigStatusResponse) SetBody(v *GetBackupAutoConfigStatusResponseBody) *GetBackupAutoConfigStatusResponse {
	s.Body = v
	return s
}

func (s *GetBackupAutoConfigStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

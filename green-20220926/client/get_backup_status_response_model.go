// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBackupStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBackupStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBackupStatusResponse
	GetStatusCode() *int32
	SetBody(v *GetBackupStatusResponseBody) *GetBackupStatusResponse
	GetBody() *GetBackupStatusResponseBody
}

type GetBackupStatusResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBackupStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBackupStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBackupStatusResponse) GoString() string {
	return s.String()
}

func (s *GetBackupStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBackupStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBackupStatusResponse) GetBody() *GetBackupStatusResponseBody {
	return s.Body
}

func (s *GetBackupStatusResponse) SetHeaders(v map[string]*string) *GetBackupStatusResponse {
	s.Headers = v
	return s
}

func (s *GetBackupStatusResponse) SetStatusCode(v int32) *GetBackupStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBackupStatusResponse) SetBody(v *GetBackupStatusResponseBody) *GetBackupStatusResponse {
	s.Body = v
	return s
}

func (s *GetBackupStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

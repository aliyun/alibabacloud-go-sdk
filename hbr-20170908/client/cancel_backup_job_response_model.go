// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelBackupJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CancelBackupJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CancelBackupJobResponse
	GetStatusCode() *int32
	SetBody(v *CancelBackupJobResponseBody) *CancelBackupJobResponse
	GetBody() *CancelBackupJobResponseBody
}

type CancelBackupJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CancelBackupJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CancelBackupJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CancelBackupJobResponse) GoString() string {
	return s.String()
}

func (s *CancelBackupJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CancelBackupJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CancelBackupJobResponse) GetBody() *CancelBackupJobResponseBody {
	return s.Body
}

func (s *CancelBackupJobResponse) SetHeaders(v map[string]*string) *CancelBackupJobResponse {
	s.Headers = v
	return s
}

func (s *CancelBackupJobResponse) SetStatusCode(v int32) *CancelBackupJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CancelBackupJobResponse) SetBody(v *CancelBackupJobResponseBody) *CancelBackupJobResponse {
	s.Body = v
	return s
}

func (s *CancelBackupJobResponse) Validate() error {
	return dara.Validate(s)
}

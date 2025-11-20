// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackupJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackupJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackupJobResponseBody) *CreateBackupJobResponse
	GetBody() *CreateBackupJobResponseBody
}

type CreateBackupJobResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackupJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupJobResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackupJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackupJobResponse) GetBody() *CreateBackupJobResponseBody {
	return s.Body
}

func (s *CreateBackupJobResponse) SetHeaders(v map[string]*string) *CreateBackupJobResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupJobResponse) SetStatusCode(v int32) *CreateBackupJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupJobResponse) SetBody(v *CreateBackupJobResponseBody) *CreateBackupJobResponse {
	s.Body = v
	return s
}

func (s *CreateBackupJobResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

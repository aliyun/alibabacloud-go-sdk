// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateApplicationBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateApplicationBackupResponse
	GetStatusCode() *int32
	SetBody(v *CreateApplicationBackupResponseBody) *CreateApplicationBackupResponse
	GetBody() *CreateApplicationBackupResponseBody
}

type CreateApplicationBackupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateApplicationBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateApplicationBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateApplicationBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateApplicationBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateApplicationBackupResponse) GetBody() *CreateApplicationBackupResponseBody {
	return s.Body
}

func (s *CreateApplicationBackupResponse) SetHeaders(v map[string]*string) *CreateApplicationBackupResponse {
	s.Headers = v
	return s
}

func (s *CreateApplicationBackupResponse) SetStatusCode(v int32) *CreateApplicationBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateApplicationBackupResponse) SetBody(v *CreateApplicationBackupResponseBody) *CreateApplicationBackupResponse {
	s.Body = v
	return s
}

func (s *CreateApplicationBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

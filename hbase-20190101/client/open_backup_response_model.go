// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OpenBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OpenBackupResponse
	GetStatusCode() *int32
	SetBody(v *OpenBackupResponseBody) *OpenBackupResponse
	GetBody() *OpenBackupResponseBody
}

type OpenBackupResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OpenBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OpenBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s OpenBackupResponse) GoString() string {
	return s.String()
}

func (s *OpenBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OpenBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OpenBackupResponse) GetBody() *OpenBackupResponseBody {
	return s.Body
}

func (s *OpenBackupResponse) SetHeaders(v map[string]*string) *OpenBackupResponse {
	s.Headers = v
	return s
}

func (s *OpenBackupResponse) SetStatusCode(v int32) *OpenBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *OpenBackupResponse) SetBody(v *OpenBackupResponseBody) *OpenBackupResponse {
	s.Body = v
	return s
}

func (s *OpenBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

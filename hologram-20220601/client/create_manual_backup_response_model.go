// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateManualBackupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateManualBackupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateManualBackupResponse
	GetStatusCode() *int32
	SetBody(v *CreateManualBackupResponseBody) *CreateManualBackupResponse
	GetBody() *CreateManualBackupResponseBody
}

type CreateManualBackupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateManualBackupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateManualBackupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateManualBackupResponse) GoString() string {
	return s.String()
}

func (s *CreateManualBackupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateManualBackupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateManualBackupResponse) GetBody() *CreateManualBackupResponseBody {
	return s.Body
}

func (s *CreateManualBackupResponse) SetHeaders(v map[string]*string) *CreateManualBackupResponse {
	s.Headers = v
	return s
}

func (s *CreateManualBackupResponse) SetStatusCode(v int32) *CreateManualBackupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateManualBackupResponse) SetBody(v *CreateManualBackupResponseBody) *CreateManualBackupResponse {
	s.Body = v
	return s
}

func (s *CreateManualBackupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

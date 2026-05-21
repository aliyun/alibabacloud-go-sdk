// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupDataResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupDataResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupDataResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupDataResponseBody) *DeleteBackupDataResponse
	GetBody() *DeleteBackupDataResponseBody
}

type DeleteBackupDataResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupDataResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupDataResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupDataResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupDataResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupDataResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupDataResponse) GetBody() *DeleteBackupDataResponseBody {
	return s.Body
}

func (s *DeleteBackupDataResponse) SetHeaders(v map[string]*string) *DeleteBackupDataResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupDataResponse) SetStatusCode(v int32) *DeleteBackupDataResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupDataResponse) SetBody(v *DeleteBackupDataResponseBody) *DeleteBackupDataResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupDataResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

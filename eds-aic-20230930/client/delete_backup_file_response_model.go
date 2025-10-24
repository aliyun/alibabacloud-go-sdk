// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupFileResponseBody) *DeleteBackupFileResponse
	GetBody() *DeleteBackupFileResponseBody
}

type DeleteBackupFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupFileResponse) GetBody() *DeleteBackupFileResponseBody {
	return s.Body
}

func (s *DeleteBackupFileResponse) SetHeaders(v map[string]*string) *DeleteBackupFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupFileResponse) SetStatusCode(v int32) *DeleteBackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupFileResponse) SetBody(v *DeleteBackupFileResponseBody) *DeleteBackupFileResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

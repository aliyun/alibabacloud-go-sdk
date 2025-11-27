// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteUserBackupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteUserBackupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteUserBackupFileResponse
	GetStatusCode() *int32
	SetBody(v *DeleteUserBackupFileResponseBody) *DeleteUserBackupFileResponse
	GetBody() *DeleteUserBackupFileResponseBody
}

type DeleteUserBackupFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteUserBackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteUserBackupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteUserBackupFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserBackupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteUserBackupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteUserBackupFileResponse) GetBody() *DeleteUserBackupFileResponseBody {
	return s.Body
}

func (s *DeleteUserBackupFileResponse) SetHeaders(v map[string]*string) *DeleteUserBackupFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserBackupFileResponse) SetStatusCode(v int32) *DeleteUserBackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteUserBackupFileResponse) SetBody(v *DeleteUserBackupFileResponseBody) *DeleteUserBackupFileResponse {
	s.Body = v
	return s
}

func (s *DeleteUserBackupFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

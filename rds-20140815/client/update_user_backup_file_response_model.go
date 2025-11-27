// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateUserBackupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateUserBackupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateUserBackupFileResponse
	GetStatusCode() *int32
	SetBody(v *UpdateUserBackupFileResponseBody) *UpdateUserBackupFileResponse
	GetBody() *UpdateUserBackupFileResponseBody
}

type UpdateUserBackupFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateUserBackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserBackupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateUserBackupFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserBackupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateUserBackupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateUserBackupFileResponse) GetBody() *UpdateUserBackupFileResponseBody {
	return s.Body
}

func (s *UpdateUserBackupFileResponse) SetHeaders(v map[string]*string) *UpdateUserBackupFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserBackupFileResponse) SetStatusCode(v int32) *UpdateUserBackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserBackupFileResponse) SetBody(v *UpdateUserBackupFileResponseBody) *UpdateUserBackupFileResponse {
	s.Body = v
	return s
}

func (s *UpdateUserBackupFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

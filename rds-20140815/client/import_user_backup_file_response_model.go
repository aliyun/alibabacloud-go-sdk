// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportUserBackupFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ImportUserBackupFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ImportUserBackupFileResponse
	GetStatusCode() *int32
	SetBody(v *ImportUserBackupFileResponseBody) *ImportUserBackupFileResponse
	GetBody() *ImportUserBackupFileResponseBody
}

type ImportUserBackupFileResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ImportUserBackupFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportUserBackupFileResponse) String() string {
	return dara.Prettify(s)
}

func (s ImportUserBackupFileResponse) GoString() string {
	return s.String()
}

func (s *ImportUserBackupFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ImportUserBackupFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ImportUserBackupFileResponse) GetBody() *ImportUserBackupFileResponseBody {
	return s.Body
}

func (s *ImportUserBackupFileResponse) SetHeaders(v map[string]*string) *ImportUserBackupFileResponse {
	s.Headers = v
	return s
}

func (s *ImportUserBackupFileResponse) SetStatusCode(v int32) *ImportUserBackupFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportUserBackupFileResponse) SetBody(v *ImportUserBackupFileResponseBody) *ImportUserBackupFileResponse {
	s.Body = v
	return s
}

func (s *ImportUserBackupFileResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

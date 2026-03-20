// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRegionSupportBackupEncryptionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CheckRegionSupportBackupEncryptionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CheckRegionSupportBackupEncryptionResponse
	GetStatusCode() *int32
	SetBody(v *CheckRegionSupportBackupEncryptionResponseBody) *CheckRegionSupportBackupEncryptionResponse
	GetBody() *CheckRegionSupportBackupEncryptionResponseBody
}

type CheckRegionSupportBackupEncryptionResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CheckRegionSupportBackupEncryptionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CheckRegionSupportBackupEncryptionResponse) String() string {
	return dara.Prettify(s)
}

func (s CheckRegionSupportBackupEncryptionResponse) GoString() string {
	return s.String()
}

func (s *CheckRegionSupportBackupEncryptionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CheckRegionSupportBackupEncryptionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CheckRegionSupportBackupEncryptionResponse) GetBody() *CheckRegionSupportBackupEncryptionResponseBody {
	return s.Body
}

func (s *CheckRegionSupportBackupEncryptionResponse) SetHeaders(v map[string]*string) *CheckRegionSupportBackupEncryptionResponse {
	s.Headers = v
	return s
}

func (s *CheckRegionSupportBackupEncryptionResponse) SetStatusCode(v int32) *CheckRegionSupportBackupEncryptionResponse {
	s.StatusCode = &v
	return s
}

func (s *CheckRegionSupportBackupEncryptionResponse) SetBody(v *CheckRegionSupportBackupEncryptionResponseBody) *CheckRegionSupportBackupEncryptionResponse {
	s.Body = v
	return s
}

func (s *CheckRegionSupportBackupEncryptionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

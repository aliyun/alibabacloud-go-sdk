// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupClientsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeBackupClientsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeBackupClientsResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeBackupClientsResponseBody) *UpgradeBackupClientsResponse
	GetBody() *UpgradeBackupClientsResponseBody
}

type UpgradeBackupClientsResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeBackupClientsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeBackupClientsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupClientsResponse) GoString() string {
	return s.String()
}

func (s *UpgradeBackupClientsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeBackupClientsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeBackupClientsResponse) GetBody() *UpgradeBackupClientsResponseBody {
	return s.Body
}

func (s *UpgradeBackupClientsResponse) SetHeaders(v map[string]*string) *UpgradeBackupClientsResponse {
	s.Headers = v
	return s
}

func (s *UpgradeBackupClientsResponse) SetStatusCode(v int32) *UpgradeBackupClientsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeBackupClientsResponse) SetBody(v *UpgradeBackupClientsResponseBody) *UpgradeBackupClientsResponse {
	s.Body = v
	return s
}

func (s *UpgradeBackupClientsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

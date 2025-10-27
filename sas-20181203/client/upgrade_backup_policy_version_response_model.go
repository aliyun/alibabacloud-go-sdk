// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupPolicyVersionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeBackupPolicyVersionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeBackupPolicyVersionResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeBackupPolicyVersionResponseBody) *UpgradeBackupPolicyVersionResponse
	GetBody() *UpgradeBackupPolicyVersionResponseBody
}

type UpgradeBackupPolicyVersionResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeBackupPolicyVersionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeBackupPolicyVersionResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupPolicyVersionResponse) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPolicyVersionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeBackupPolicyVersionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeBackupPolicyVersionResponse) GetBody() *UpgradeBackupPolicyVersionResponseBody {
	return s.Body
}

func (s *UpgradeBackupPolicyVersionResponse) SetHeaders(v map[string]*string) *UpgradeBackupPolicyVersionResponse {
	s.Headers = v
	return s
}

func (s *UpgradeBackupPolicyVersionResponse) SetStatusCode(v int32) *UpgradeBackupPolicyVersionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeBackupPolicyVersionResponse) SetBody(v *UpgradeBackupPolicyVersionResponseBody) *UpgradeBackupPolicyVersionResponse {
	s.Body = v
	return s
}

func (s *UpgradeBackupPolicyVersionResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

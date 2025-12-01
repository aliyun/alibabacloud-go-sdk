// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpgradeBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpgradeBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpgradeBackupPlanResponseBody) *UpgradeBackupPlanResponse
	GetBody() *UpgradeBackupPlanResponseBody
}

type UpgradeBackupPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpgradeBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpgradeBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpgradeBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *UpgradeBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpgradeBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpgradeBackupPlanResponse) GetBody() *UpgradeBackupPlanResponseBody {
	return s.Body
}

func (s *UpgradeBackupPlanResponse) SetHeaders(v map[string]*string) *UpgradeBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *UpgradeBackupPlanResponse) SetStatusCode(v int32) *UpgradeBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpgradeBackupPlanResponse) SetBody(v *UpgradeBackupPlanResponseBody) *UpgradeBackupPlanResponse {
	s.Body = v
	return s
}

func (s *UpgradeBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableHanaBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableHanaBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableHanaBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *DisableHanaBackupPlanResponseBody) *DisableHanaBackupPlanResponse
	GetBody() *DisableHanaBackupPlanResponseBody
}

type DisableHanaBackupPlanResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableHanaBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableHanaBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableHanaBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableHanaBackupPlanResponse) GetBody() *DisableHanaBackupPlanResponseBody {
	return s.Body
}

func (s *DisableHanaBackupPlanResponse) SetHeaders(v map[string]*string) *DisableHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableHanaBackupPlanResponse) SetStatusCode(v int32) *DisableHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableHanaBackupPlanResponse) SetBody(v *DisableHanaBackupPlanResponseBody) *DisableHanaBackupPlanResponse {
	s.Body = v
	return s
}

func (s *DisableHanaBackupPlanResponse) Validate() error {
	return dara.Validate(s)
}

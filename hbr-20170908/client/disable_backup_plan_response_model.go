// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *DisableBackupPlanResponseBody) *DisableBackupPlanResponse
	GetBody() *DisableBackupPlanResponseBody
}

type DisableBackupPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DisableBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableBackupPlanResponse) GetBody() *DisableBackupPlanResponseBody {
	return s.Body
}

func (s *DisableBackupPlanResponse) SetHeaders(v map[string]*string) *DisableBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DisableBackupPlanResponse) SetStatusCode(v int32) *DisableBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableBackupPlanResponse) SetBody(v *DisableBackupPlanResponseBody) *DisableBackupPlanResponse {
	s.Body = v
	return s
}

func (s *DisableBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

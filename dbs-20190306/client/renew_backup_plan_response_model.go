// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RenewBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RenewBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *RenewBackupPlanResponseBody) *RenewBackupPlanResponse
	GetBody() *RenewBackupPlanResponseBody
}

type RenewBackupPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RenewBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RenewBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s RenewBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *RenewBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RenewBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RenewBackupPlanResponse) GetBody() *RenewBackupPlanResponseBody {
	return s.Body
}

func (s *RenewBackupPlanResponse) SetHeaders(v map[string]*string) *RenewBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *RenewBackupPlanResponse) SetStatusCode(v int32) *RenewBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *RenewBackupPlanResponse) SetBody(v *RenewBackupPlanResponseBody) *RenewBackupPlanResponse {
	s.Body = v
	return s
}

func (s *RenewBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

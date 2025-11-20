// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBackupPlanResponseBody) *UpdateBackupPlanResponse
	GetBody() *UpdateBackupPlanResponseBody
}

type UpdateBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBackupPlanResponse) GetBody() *UpdateBackupPlanResponseBody {
	return s.Body
}

func (s *UpdateBackupPlanResponse) SetHeaders(v map[string]*string) *UpdateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateBackupPlanResponse) SetStatusCode(v int32) *UpdateBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBackupPlanResponse) SetBody(v *UpdateBackupPlanResponseBody) *UpdateBackupPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

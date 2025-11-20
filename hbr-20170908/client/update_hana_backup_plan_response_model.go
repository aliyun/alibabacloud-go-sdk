// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateHanaBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateHanaBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateHanaBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *UpdateHanaBackupPlanResponseBody) *UpdateHanaBackupPlanResponse
	GetBody() *UpdateHanaBackupPlanResponseBody
}

type UpdateHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateHanaBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *UpdateHanaBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateHanaBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateHanaBackupPlanResponse) GetBody() *UpdateHanaBackupPlanResponseBody {
	return s.Body
}

func (s *UpdateHanaBackupPlanResponse) SetHeaders(v map[string]*string) *UpdateHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *UpdateHanaBackupPlanResponse) SetStatusCode(v int32) *UpdateHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateHanaBackupPlanResponse) SetBody(v *UpdateHanaBackupPlanResponseBody) *UpdateHanaBackupPlanResponse {
	s.Body = v
	return s
}

func (s *UpdateHanaBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

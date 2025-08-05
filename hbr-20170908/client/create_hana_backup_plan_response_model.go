// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHanaBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHanaBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHanaBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateHanaBackupPlanResponseBody) *CreateHanaBackupPlanResponse
	GetBody() *CreateHanaBackupPlanResponseBody
}

type CreateHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHanaBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateHanaBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHanaBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHanaBackupPlanResponse) GetBody() *CreateHanaBackupPlanResponseBody {
	return s.Body
}

func (s *CreateHanaBackupPlanResponse) SetHeaders(v map[string]*string) *CreateHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateHanaBackupPlanResponse) SetStatusCode(v int32) *CreateHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHanaBackupPlanResponse) SetBody(v *CreateHanaBackupPlanResponseBody) *CreateHanaBackupPlanResponse {
	s.Body = v
	return s
}

func (s *CreateHanaBackupPlanResponse) Validate() error {
	return dara.Validate(s)
}

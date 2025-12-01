// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAndStartBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAndStartBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAndStartBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateAndStartBackupPlanResponseBody) *CreateAndStartBackupPlanResponse
	GetBody() *CreateAndStartBackupPlanResponseBody
}

type CreateAndStartBackupPlanResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAndStartBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAndStartBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAndStartBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateAndStartBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAndStartBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAndStartBackupPlanResponse) GetBody() *CreateAndStartBackupPlanResponseBody {
	return s.Body
}

func (s *CreateAndStartBackupPlanResponse) SetHeaders(v map[string]*string) *CreateAndStartBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateAndStartBackupPlanResponse) SetStatusCode(v int32) *CreateAndStartBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAndStartBackupPlanResponse) SetBody(v *CreateAndStartBackupPlanResponseBody) *CreateAndStartBackupPlanResponse {
	s.Body = v
	return s
}

func (s *CreateAndStartBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

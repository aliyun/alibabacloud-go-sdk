// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse
	GetBody() *CreateBackupPlanResponseBody
}

type CreateBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBackupPlanResponse) GetBody() *CreateBackupPlanResponseBody {
	return s.Body
}

func (s *CreateBackupPlanResponse) SetHeaders(v map[string]*string) *CreateBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *CreateBackupPlanResponse) SetStatusCode(v int32) *CreateBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBackupPlanResponse) SetBody(v *CreateBackupPlanResponseBody) *CreateBackupPlanResponse {
	s.Body = v
	return s
}

func (s *CreateBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

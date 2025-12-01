// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPlanNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupPlanNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupPlanNameResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupPlanNameResponseBody) *ModifyBackupPlanNameResponse
	GetBody() *ModifyBackupPlanNameResponseBody
}

type ModifyBackupPlanNameResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPlanNameResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPlanNameResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPlanNameResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupPlanNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupPlanNameResponse) GetBody() *ModifyBackupPlanNameResponseBody {
	return s.Body
}

func (s *ModifyBackupPlanNameResponse) SetHeaders(v map[string]*string) *ModifyBackupPlanNameResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPlanNameResponse) SetStatusCode(v int32) *ModifyBackupPlanNameResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPlanNameResponse) SetBody(v *ModifyBackupPlanNameResponseBody) *ModifyBackupPlanNameResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupPlanNameResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

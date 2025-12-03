// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPlanConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyBackupPlanConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyBackupPlanConfigResponse
	GetStatusCode() *int32
	SetBody(v *ModifyBackupPlanConfigResponseBody) *ModifyBackupPlanConfigResponse
	GetBody() *ModifyBackupPlanConfigResponseBody
}

type ModifyBackupPlanConfigResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyBackupPlanConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyBackupPlanConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPlanConfigResponse) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyBackupPlanConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyBackupPlanConfigResponse) GetBody() *ModifyBackupPlanConfigResponseBody {
	return s.Body
}

func (s *ModifyBackupPlanConfigResponse) SetHeaders(v map[string]*string) *ModifyBackupPlanConfigResponse {
	s.Headers = v
	return s
}

func (s *ModifyBackupPlanConfigResponse) SetStatusCode(v int32) *ModifyBackupPlanConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyBackupPlanConfigResponse) SetBody(v *ModifyBackupPlanConfigResponseBody) *ModifyBackupPlanConfigResponse {
	s.Body = v
	return s
}

func (s *ModifyBackupPlanConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

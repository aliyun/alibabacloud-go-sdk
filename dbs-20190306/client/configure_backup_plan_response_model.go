// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ConfigureBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ConfigureBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *ConfigureBackupPlanResponseBody) *ConfigureBackupPlanResponse
	GetBody() *ConfigureBackupPlanResponseBody
}

type ConfigureBackupPlanResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ConfigureBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ConfigureBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *ConfigureBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ConfigureBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ConfigureBackupPlanResponse) GetBody() *ConfigureBackupPlanResponseBody {
	return s.Body
}

func (s *ConfigureBackupPlanResponse) SetHeaders(v map[string]*string) *ConfigureBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *ConfigureBackupPlanResponse) SetStatusCode(v int32) *ConfigureBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ConfigureBackupPlanResponse) SetBody(v *ConfigureBackupPlanResponseBody) *ConfigureBackupPlanResponse {
	s.Body = v
	return s
}

func (s *ConfigureBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

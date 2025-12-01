// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *StartBackupPlanResponseBody) *StartBackupPlanResponse
	GetBody() *StartBackupPlanResponseBody
}

type StartBackupPlanResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s StartBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *StartBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartBackupPlanResponse) GetBody() *StartBackupPlanResponseBody {
	return s.Body
}

func (s *StartBackupPlanResponse) SetHeaders(v map[string]*string) *StartBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *StartBackupPlanResponse) SetStatusCode(v int32) *StartBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBackupPlanResponse) SetBody(v *StartBackupPlanResponseBody) *StartBackupPlanResponse {
	s.Body = v
	return s
}

func (s *StartBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

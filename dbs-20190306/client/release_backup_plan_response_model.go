// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReleaseBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReleaseBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReleaseBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *ReleaseBackupPlanResponseBody) *ReleaseBackupPlanResponse
	GetBody() *ReleaseBackupPlanResponseBody
}

type ReleaseBackupPlanResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReleaseBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReleaseBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s ReleaseBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *ReleaseBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReleaseBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReleaseBackupPlanResponse) GetBody() *ReleaseBackupPlanResponseBody {
	return s.Body
}

func (s *ReleaseBackupPlanResponse) SetHeaders(v map[string]*string) *ReleaseBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *ReleaseBackupPlanResponse) SetStatusCode(v int32) *ReleaseBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *ReleaseBackupPlanResponse) SetBody(v *ReleaseBackupPlanResponseBody) *ReleaseBackupPlanResponse {
	s.Body = v
	return s
}

func (s *ReleaseBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

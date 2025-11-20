// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteBackupPlanResponseBody) *DeleteBackupPlanResponse
	GetBody() *DeleteBackupPlanResponseBody
}

type DeleteBackupPlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteBackupPlanResponse) GetBody() *DeleteBackupPlanResponseBody {
	return s.Body
}

func (s *DeleteBackupPlanResponse) SetHeaders(v map[string]*string) *DeleteBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteBackupPlanResponse) SetStatusCode(v int32) *DeleteBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteBackupPlanResponse) SetBody(v *DeleteBackupPlanResponseBody) *DeleteBackupPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

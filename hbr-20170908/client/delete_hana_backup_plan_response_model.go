// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHanaBackupPlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHanaBackupPlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHanaBackupPlanResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHanaBackupPlanResponseBody) *DeleteHanaBackupPlanResponse
	GetBody() *DeleteHanaBackupPlanResponseBody
}

type DeleteHanaBackupPlanResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHanaBackupPlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHanaBackupPlanResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHanaBackupPlanResponse) GoString() string {
	return s.String()
}

func (s *DeleteHanaBackupPlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHanaBackupPlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHanaBackupPlanResponse) GetBody() *DeleteHanaBackupPlanResponseBody {
	return s.Body
}

func (s *DeleteHanaBackupPlanResponse) SetHeaders(v map[string]*string) *DeleteHanaBackupPlanResponse {
	s.Headers = v
	return s
}

func (s *DeleteHanaBackupPlanResponse) SetStatusCode(v int32) *DeleteHanaBackupPlanResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHanaBackupPlanResponse) SetBody(v *DeleteHanaBackupPlanResponseBody) *DeleteHanaBackupPlanResponse {
	s.Body = v
	return s
}

func (s *DeleteHanaBackupPlanResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

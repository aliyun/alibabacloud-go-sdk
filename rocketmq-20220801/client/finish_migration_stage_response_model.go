// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFinishMigrationStageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *FinishMigrationStageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *FinishMigrationStageResponse
	GetStatusCode() *int32
	SetBody(v *FinishMigrationStageResponseBody) *FinishMigrationStageResponse
	GetBody() *FinishMigrationStageResponseBody
}

type FinishMigrationStageResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FinishMigrationStageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FinishMigrationStageResponse) String() string {
	return dara.Prettify(s)
}

func (s FinishMigrationStageResponse) GoString() string {
	return s.String()
}

func (s *FinishMigrationStageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *FinishMigrationStageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *FinishMigrationStageResponse) GetBody() *FinishMigrationStageResponseBody {
	return s.Body
}

func (s *FinishMigrationStageResponse) SetHeaders(v map[string]*string) *FinishMigrationStageResponse {
	s.Headers = v
	return s
}

func (s *FinishMigrationStageResponse) SetStatusCode(v int32) *FinishMigrationStageResponse {
	s.StatusCode = &v
	return s
}

func (s *FinishMigrationStageResponse) SetBody(v *FinishMigrationStageResponseBody) *FinishMigrationStageResponse {
	s.Body = v
	return s
}

func (s *FinishMigrationStageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

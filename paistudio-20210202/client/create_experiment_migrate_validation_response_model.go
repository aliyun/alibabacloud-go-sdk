// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentMigrateValidationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperimentMigrateValidationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperimentMigrateValidationResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperimentMigrateValidationResponseBody) *CreateExperimentMigrateValidationResponse
	GetBody() *CreateExperimentMigrateValidationResponseBody
}

type CreateExperimentMigrateValidationResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentMigrateValidationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentMigrateValidationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentMigrateValidationResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentMigrateValidationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperimentMigrateValidationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperimentMigrateValidationResponse) GetBody() *CreateExperimentMigrateValidationResponseBody {
	return s.Body
}

func (s *CreateExperimentMigrateValidationResponse) SetHeaders(v map[string]*string) *CreateExperimentMigrateValidationResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentMigrateValidationResponse) SetStatusCode(v int32) *CreateExperimentMigrateValidationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentMigrateValidationResponse) SetBody(v *CreateExperimentMigrateValidationResponseBody) *CreateExperimentMigrateValidationResponse {
	s.Body = v
	return s
}

func (s *CreateExperimentMigrateValidationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

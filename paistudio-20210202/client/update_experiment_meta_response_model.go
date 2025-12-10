// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentMetaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentMetaResponseBody) *UpdateExperimentMetaResponse
	GetBody() *UpdateExperimentMetaResponseBody
}

type UpdateExperimentMetaResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentMetaResponse) GetBody() *UpdateExperimentMetaResponseBody {
	return s.Body
}

func (s *UpdateExperimentMetaResponse) SetHeaders(v map[string]*string) *UpdateExperimentMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentMetaResponse) SetStatusCode(v int32) *UpdateExperimentMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentMetaResponse) SetBody(v *UpdateExperimentMetaResponseBody) *UpdateExperimentMetaResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentMetaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

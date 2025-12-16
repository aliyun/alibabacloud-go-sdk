// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateABTestExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateABTestExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateABTestExperimentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateABTestExperimentResponseBody) *UpdateABTestExperimentResponse
	GetBody() *UpdateABTestExperimentResponseBody
}

type UpdateABTestExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateABTestExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *UpdateABTestExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateABTestExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateABTestExperimentResponse) GetBody() *UpdateABTestExperimentResponseBody {
	return s.Body
}

func (s *UpdateABTestExperimentResponse) SetHeaders(v map[string]*string) *UpdateABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *UpdateABTestExperimentResponse) SetStatusCode(v int32) *UpdateABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateABTestExperimentResponse) SetBody(v *UpdateABTestExperimentResponseBody) *UpdateABTestExperimentResponse {
	s.Body = v
	return s
}

func (s *UpdateABTestExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

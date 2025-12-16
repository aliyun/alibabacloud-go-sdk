// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateABTestExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateABTestExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateABTestExperimentResponse
	GetStatusCode() *int32
	SetBody(v *CreateABTestExperimentResponseBody) *CreateABTestExperimentResponse
	GetBody() *CreateABTestExperimentResponseBody
}

type CreateABTestExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateABTestExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *CreateABTestExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateABTestExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateABTestExperimentResponse) GetBody() *CreateABTestExperimentResponseBody {
	return s.Body
}

func (s *CreateABTestExperimentResponse) SetHeaders(v map[string]*string) *CreateABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *CreateABTestExperimentResponse) SetStatusCode(v int32) *CreateABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateABTestExperimentResponse) SetBody(v *CreateABTestExperimentResponseBody) *CreateABTestExperimentResponse {
	s.Body = v
	return s
}

func (s *CreateABTestExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteABTestExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteABTestExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteABTestExperimentResponse
	GetStatusCode() *int32
	SetBody(v *DeleteABTestExperimentResponseBody) *DeleteABTestExperimentResponse
	GetBody() *DeleteABTestExperimentResponseBody
}

type DeleteABTestExperimentResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteABTestExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteABTestExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteABTestExperimentResponse) GoString() string {
	return s.String()
}

func (s *DeleteABTestExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteABTestExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteABTestExperimentResponse) GetBody() *DeleteABTestExperimentResponseBody {
	return s.Body
}

func (s *DeleteABTestExperimentResponse) SetHeaders(v map[string]*string) *DeleteABTestExperimentResponse {
	s.Headers = v
	return s
}

func (s *DeleteABTestExperimentResponse) SetStatusCode(v int32) *DeleteABTestExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteABTestExperimentResponse) SetBody(v *DeleteABTestExperimentResponseBody) *DeleteABTestExperimentResponse {
	s.Body = v
	return s
}

func (s *DeleteABTestExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentContentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentContentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentContentResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentContentResponseBody) *UpdateExperimentContentResponse
	GetBody() *UpdateExperimentContentResponseBody
}

type UpdateExperimentContentResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentContentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentContentResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentContentResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentContentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentContentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentContentResponse) GetBody() *UpdateExperimentContentResponseBody {
	return s.Body
}

func (s *UpdateExperimentContentResponse) SetHeaders(v map[string]*string) *UpdateExperimentContentResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentContentResponse) SetStatusCode(v int32) *UpdateExperimentContentResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentContentResponse) SetBody(v *UpdateExperimentContentResponseBody) *UpdateExperimentContentResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentContentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

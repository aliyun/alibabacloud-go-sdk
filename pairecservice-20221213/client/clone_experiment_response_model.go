// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneExperimentResponse
	GetStatusCode() *int32
	SetBody(v *CloneExperimentResponseBody) *CloneExperimentResponse
	GetBody() *CloneExperimentResponseBody
}

type CloneExperimentResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneExperimentResponse) GoString() string {
	return s.String()
}

func (s *CloneExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneExperimentResponse) GetBody() *CloneExperimentResponseBody {
	return s.Body
}

func (s *CloneExperimentResponse) SetHeaders(v map[string]*string) *CloneExperimentResponse {
	s.Headers = v
	return s
}

func (s *CloneExperimentResponse) SetStatusCode(v int32) *CloneExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneExperimentResponse) SetBody(v *CloneExperimentResponseBody) *CloneExperimentResponse {
	s.Body = v
	return s
}

func (s *CloneExperimentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

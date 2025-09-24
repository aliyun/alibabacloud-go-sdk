// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListExperimentResponse
	GetStatusCode() *int32
	SetBody(v *ListExperimentResponseBody) *ListExperimentResponse
	GetBody() *ListExperimentResponseBody
}

type ListExperimentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s ListExperimentResponse) GoString() string {
	return s.String()
}

func (s *ListExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListExperimentResponse) GetBody() *ListExperimentResponseBody {
	return s.Body
}

func (s *ListExperimentResponse) SetHeaders(v map[string]*string) *ListExperimentResponse {
	s.Headers = v
	return s
}

func (s *ListExperimentResponse) SetStatusCode(v int32) *ListExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListExperimentResponse) SetBody(v *ListExperimentResponseBody) *ListExperimentResponse {
	s.Body = v
	return s
}

func (s *ListExperimentResponse) Validate() error {
	return dara.Validate(s)
}

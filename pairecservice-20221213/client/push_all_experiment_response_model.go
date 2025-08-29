// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushAllExperimentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PushAllExperimentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PushAllExperimentResponse
	GetStatusCode() *int32
	SetBody(v *PushAllExperimentResponseBody) *PushAllExperimentResponse
	GetBody() *PushAllExperimentResponseBody
}

type PushAllExperimentResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PushAllExperimentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PushAllExperimentResponse) String() string {
	return dara.Prettify(s)
}

func (s PushAllExperimentResponse) GoString() string {
	return s.String()
}

func (s *PushAllExperimentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PushAllExperimentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PushAllExperimentResponse) GetBody() *PushAllExperimentResponseBody {
	return s.Body
}

func (s *PushAllExperimentResponse) SetHeaders(v map[string]*string) *PushAllExperimentResponse {
	s.Headers = v
	return s
}

func (s *PushAllExperimentResponse) SetStatusCode(v int32) *PushAllExperimentResponse {
	s.StatusCode = &v
	return s
}

func (s *PushAllExperimentResponse) SetBody(v *PushAllExperimentResponseBody) *PushAllExperimentResponse {
	s.Body = v
	return s
}

func (s *PushAllExperimentResponse) Validate() error {
	return dara.Validate(s)
}

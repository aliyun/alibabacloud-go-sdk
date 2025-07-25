// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrainingJobResponseBody) *CreateTrainingJobResponse
	GetBody() *CreateTrainingJobResponseBody
}

type CreateTrainingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *CreateTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrainingJobResponse) GetBody() *CreateTrainingJobResponseBody {
	return s.Body
}

func (s *CreateTrainingJobResponse) SetHeaders(v map[string]*string) *CreateTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *CreateTrainingJobResponse) SetStatusCode(v int32) *CreateTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrainingJobResponse) SetBody(v *CreateTrainingJobResponseBody) *CreateTrainingJobResponse {
	s.Body = v
	return s
}

func (s *CreateTrainingJobResponse) Validate() error {
	return dara.Validate(s)
}

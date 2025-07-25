// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateTrainingJobLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateTrainingJobLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateTrainingJobLabelsResponse
	GetStatusCode() *int32
	SetBody(v *UpdateTrainingJobLabelsResponseBody) *UpdateTrainingJobLabelsResponse
	GetBody() *UpdateTrainingJobLabelsResponseBody
}

type UpdateTrainingJobLabelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateTrainingJobLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateTrainingJobLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateTrainingJobLabelsResponse) GoString() string {
	return s.String()
}

func (s *UpdateTrainingJobLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateTrainingJobLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateTrainingJobLabelsResponse) GetBody() *UpdateTrainingJobLabelsResponseBody {
	return s.Body
}

func (s *UpdateTrainingJobLabelsResponse) SetHeaders(v map[string]*string) *UpdateTrainingJobLabelsResponse {
	s.Headers = v
	return s
}

func (s *UpdateTrainingJobLabelsResponse) SetStatusCode(v int32) *UpdateTrainingJobLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateTrainingJobLabelsResponse) SetBody(v *UpdateTrainingJobLabelsResponseBody) *UpdateTrainingJobLabelsResponse {
	s.Body = v
	return s
}

func (s *UpdateTrainingJobLabelsResponse) Validate() error {
	return dara.Validate(s)
}

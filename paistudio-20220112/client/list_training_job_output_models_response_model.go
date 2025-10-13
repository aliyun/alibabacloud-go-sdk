// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTrainingJobOutputModelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTrainingJobOutputModelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTrainingJobOutputModelsResponse
	GetStatusCode() *int32
	SetBody(v *ListTrainingJobOutputModelsResponseBody) *ListTrainingJobOutputModelsResponse
	GetBody() *ListTrainingJobOutputModelsResponseBody
}

type ListTrainingJobOutputModelsResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTrainingJobOutputModelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTrainingJobOutputModelsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTrainingJobOutputModelsResponse) GoString() string {
	return s.String()
}

func (s *ListTrainingJobOutputModelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTrainingJobOutputModelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTrainingJobOutputModelsResponse) GetBody() *ListTrainingJobOutputModelsResponseBody {
	return s.Body
}

func (s *ListTrainingJobOutputModelsResponse) SetHeaders(v map[string]*string) *ListTrainingJobOutputModelsResponse {
	s.Headers = v
	return s
}

func (s *ListTrainingJobOutputModelsResponse) SetStatusCode(v int32) *ListTrainingJobOutputModelsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTrainingJobOutputModelsResponse) SetBody(v *ListTrainingJobOutputModelsResponseBody) *ListTrainingJobOutputModelsResponse {
	s.Body = v
	return s
}

func (s *ListTrainingJobOutputModelsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

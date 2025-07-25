// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrainingJobLabelsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrainingJobLabelsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrainingJobLabelsResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrainingJobLabelsResponseBody) *DeleteTrainingJobLabelsResponse
	GetBody() *DeleteTrainingJobLabelsResponseBody
}

type DeleteTrainingJobLabelsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrainingJobLabelsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrainingJobLabelsResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrainingJobLabelsResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobLabelsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrainingJobLabelsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrainingJobLabelsResponse) GetBody() *DeleteTrainingJobLabelsResponseBody {
	return s.Body
}

func (s *DeleteTrainingJobLabelsResponse) SetHeaders(v map[string]*string) *DeleteTrainingJobLabelsResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrainingJobLabelsResponse) SetStatusCode(v int32) *DeleteTrainingJobLabelsResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrainingJobLabelsResponse) SetBody(v *DeleteTrainingJobLabelsResponseBody) *DeleteTrainingJobLabelsResponse {
	s.Body = v
	return s
}

func (s *DeleteTrainingJobLabelsResponse) Validate() error {
	return dara.Validate(s)
}

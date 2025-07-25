// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTrainingJobResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTrainingJobResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTrainingJobResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTrainingJobResponseBody) *DeleteTrainingJobResponse
	GetBody() *DeleteTrainingJobResponseBody
}

type DeleteTrainingJobResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTrainingJobResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTrainingJobResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTrainingJobResponse) GoString() string {
	return s.String()
}

func (s *DeleteTrainingJobResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTrainingJobResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTrainingJobResponse) GetBody() *DeleteTrainingJobResponseBody {
	return s.Body
}

func (s *DeleteTrainingJobResponse) SetHeaders(v map[string]*string) *DeleteTrainingJobResponse {
	s.Headers = v
	return s
}

func (s *DeleteTrainingJobResponse) SetStatusCode(v int32) *DeleteTrainingJobResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTrainingJobResponse) SetBody(v *DeleteTrainingJobResponseBody) *DeleteTrainingJobResponse {
	s.Body = v
	return s
}

func (s *DeleteTrainingJobResponse) Validate() error {
	return dara.Validate(s)
}

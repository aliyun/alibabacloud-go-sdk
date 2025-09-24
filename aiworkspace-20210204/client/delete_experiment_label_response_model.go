// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentLabelResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperimentLabelResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperimentLabelResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperimentLabelResponseBody) *DeleteExperimentLabelResponse
	GetBody() *DeleteExperimentLabelResponseBody
}

type DeleteExperimentLabelResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentLabelResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentLabelResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentLabelResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentLabelResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperimentLabelResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperimentLabelResponse) GetBody() *DeleteExperimentLabelResponseBody {
	return s.Body
}

func (s *DeleteExperimentLabelResponse) SetHeaders(v map[string]*string) *DeleteExperimentLabelResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentLabelResponse) SetStatusCode(v int32) *DeleteExperimentLabelResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentLabelResponse) SetBody(v *DeleteExperimentLabelResponseBody) *DeleteExperimentLabelResponse {
	s.Body = v
	return s
}

func (s *DeleteExperimentLabelResponse) Validate() error {
	return dara.Validate(s)
}

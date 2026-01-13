// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteExperimentGroupResponseBody) *DeleteExperimentGroupResponse
	GetBody() *DeleteExperimentGroupResponseBody
}

type DeleteExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteExperimentGroupResponse) GetBody() *DeleteExperimentGroupResponseBody {
	return s.Body
}

func (s *DeleteExperimentGroupResponse) SetHeaders(v map[string]*string) *DeleteExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteExperimentGroupResponse) SetStatusCode(v int32) *DeleteExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteExperimentGroupResponse) SetBody(v *DeleteExperimentGroupResponseBody) *DeleteExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteExperimentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

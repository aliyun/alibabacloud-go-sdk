// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateExperimentGroupResponseBody) *CreateExperimentGroupResponse
	GetBody() *CreateExperimentGroupResponseBody
}

type CreateExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateExperimentGroupResponse) GetBody() *CreateExperimentGroupResponseBody {
	return s.Body
}

func (s *CreateExperimentGroupResponse) SetHeaders(v map[string]*string) *CreateExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateExperimentGroupResponse) SetStatusCode(v int32) *CreateExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateExperimentGroupResponse) SetBody(v *CreateExperimentGroupResponseBody) *CreateExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *CreateExperimentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

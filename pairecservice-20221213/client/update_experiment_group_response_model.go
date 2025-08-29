// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateExperimentGroupResponseBody) *UpdateExperimentGroupResponse
	GetBody() *UpdateExperimentGroupResponseBody
}

type UpdateExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateExperimentGroupResponse) GetBody() *UpdateExperimentGroupResponseBody {
	return s.Body
}

func (s *UpdateExperimentGroupResponse) SetHeaders(v map[string]*string) *UpdateExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateExperimentGroupResponse) SetStatusCode(v int32) *UpdateExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateExperimentGroupResponse) SetBody(v *UpdateExperimentGroupResponseBody) *UpdateExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *UpdateExperimentGroupResponse) Validate() error {
	return dara.Validate(s)
}

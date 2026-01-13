// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloneExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloneExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloneExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *CloneExperimentGroupResponseBody) *CloneExperimentGroupResponse
	GetBody() *CloneExperimentGroupResponseBody
}

type CloneExperimentGroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloneExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloneExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CloneExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *CloneExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloneExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloneExperimentGroupResponse) GetBody() *CloneExperimentGroupResponseBody {
	return s.Body
}

func (s *CloneExperimentGroupResponse) SetHeaders(v map[string]*string) *CloneExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *CloneExperimentGroupResponse) SetStatusCode(v int32) *CloneExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CloneExperimentGroupResponse) SetBody(v *CloneExperimentGroupResponseBody) *CloneExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *CloneExperimentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

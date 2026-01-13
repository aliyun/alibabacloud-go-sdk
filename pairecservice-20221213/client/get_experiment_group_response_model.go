// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *GetExperimentGroupResponseBody) *GetExperimentGroupResponse
	GetBody() *GetExperimentGroupResponseBody
}

type GetExperimentGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s GetExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *GetExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetExperimentGroupResponse) GetBody() *GetExperimentGroupResponseBody {
	return s.Body
}

func (s *GetExperimentGroupResponse) SetHeaders(v map[string]*string) *GetExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *GetExperimentGroupResponse) SetStatusCode(v int32) *GetExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetExperimentGroupResponse) SetBody(v *GetExperimentGroupResponseBody) *GetExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *GetExperimentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

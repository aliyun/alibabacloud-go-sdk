// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineExperimentGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *OnlineExperimentGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *OnlineExperimentGroupResponse
	GetStatusCode() *int32
	SetBody(v *OnlineExperimentGroupResponseBody) *OnlineExperimentGroupResponse
	GetBody() *OnlineExperimentGroupResponseBody
}

type OnlineExperimentGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *OnlineExperimentGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s OnlineExperimentGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s OnlineExperimentGroupResponse) GoString() string {
	return s.String()
}

func (s *OnlineExperimentGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *OnlineExperimentGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *OnlineExperimentGroupResponse) GetBody() *OnlineExperimentGroupResponseBody {
	return s.Body
}

func (s *OnlineExperimentGroupResponse) SetHeaders(v map[string]*string) *OnlineExperimentGroupResponse {
	s.Headers = v
	return s
}

func (s *OnlineExperimentGroupResponse) SetStatusCode(v int32) *OnlineExperimentGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *OnlineExperimentGroupResponse) SetBody(v *OnlineExperimentGroupResponseBody) *OnlineExperimentGroupResponse {
	s.Body = v
	return s
}

func (s *OnlineExperimentGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinPipelineGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinPipelineGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinPipelineGroupResponse
	GetStatusCode() *int32
	SetBody(v *JoinPipelineGroupResponseBody) *JoinPipelineGroupResponse
	GetBody() *JoinPipelineGroupResponseBody
}

type JoinPipelineGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinPipelineGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinPipelineGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinPipelineGroupResponse) GoString() string {
	return s.String()
}

func (s *JoinPipelineGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinPipelineGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinPipelineGroupResponse) GetBody() *JoinPipelineGroupResponseBody {
	return s.Body
}

func (s *JoinPipelineGroupResponse) SetHeaders(v map[string]*string) *JoinPipelineGroupResponse {
	s.Headers = v
	return s
}

func (s *JoinPipelineGroupResponse) SetStatusCode(v int32) *JoinPipelineGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinPipelineGroupResponse) SetBody(v *JoinPipelineGroupResponseBody) *JoinPipelineGroupResponse {
	s.Body = v
	return s
}

func (s *JoinPipelineGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RestartNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RestartNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *RestartNodeGroupResponseBody) *RestartNodeGroupResponse
	GetBody() *RestartNodeGroupResponseBody
}

type RestartNodeGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestartNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestartNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s RestartNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *RestartNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RestartNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RestartNodeGroupResponse) GetBody() *RestartNodeGroupResponseBody {
	return s.Body
}

func (s *RestartNodeGroupResponse) SetHeaders(v map[string]*string) *RestartNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *RestartNodeGroupResponse) SetStatusCode(v int32) *RestartNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *RestartNodeGroupResponse) SetBody(v *RestartNodeGroupResponseBody) *RestartNodeGroupResponse {
	s.Body = v
	return s
}

func (s *RestartNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

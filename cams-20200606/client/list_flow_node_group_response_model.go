// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFlowNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFlowNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFlowNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListFlowNodeGroupResponseBody) *ListFlowNodeGroupResponse
	GetBody() *ListFlowNodeGroupResponseBody
}

type ListFlowNodeGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFlowNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFlowNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFlowNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *ListFlowNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFlowNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFlowNodeGroupResponse) GetBody() *ListFlowNodeGroupResponseBody {
	return s.Body
}

func (s *ListFlowNodeGroupResponse) SetHeaders(v map[string]*string) *ListFlowNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *ListFlowNodeGroupResponse) SetStatusCode(v int32) *ListFlowNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFlowNodeGroupResponse) SetBody(v *ListFlowNodeGroupResponseBody) *ListFlowNodeGroupResponse {
	s.Body = v
	return s
}

func (s *ListFlowNodeGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

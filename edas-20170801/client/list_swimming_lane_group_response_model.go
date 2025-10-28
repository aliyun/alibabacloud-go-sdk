// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSwimmingLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSwimmingLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *ListSwimmingLaneGroupResponseBody) *ListSwimmingLaneGroupResponse
	GetBody() *ListSwimmingLaneGroupResponseBody
}

type ListSwimmingLaneGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSwimmingLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSwimmingLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSwimmingLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSwimmingLaneGroupResponse) GetBody() *ListSwimmingLaneGroupResponseBody {
	return s.Body
}

func (s *ListSwimmingLaneGroupResponse) SetHeaders(v map[string]*string) *ListSwimmingLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *ListSwimmingLaneGroupResponse) SetStatusCode(v int32) *ListSwimmingLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSwimmingLaneGroupResponse) SetBody(v *ListSwimmingLaneGroupResponseBody) *ListSwimmingLaneGroupResponse {
	s.Body = v
	return s
}

func (s *ListSwimmingLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

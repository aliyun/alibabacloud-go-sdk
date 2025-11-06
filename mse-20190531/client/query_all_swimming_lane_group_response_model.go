// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAllSwimmingLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAllSwimmingLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAllSwimmingLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryAllSwimmingLaneGroupResponseBody) *QueryAllSwimmingLaneGroupResponse
	GetBody() *QueryAllSwimmingLaneGroupResponseBody
}

type QueryAllSwimmingLaneGroupResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAllSwimmingLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAllSwimmingLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAllSwimmingLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryAllSwimmingLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAllSwimmingLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAllSwimmingLaneGroupResponse) GetBody() *QueryAllSwimmingLaneGroupResponseBody {
	return s.Body
}

func (s *QueryAllSwimmingLaneGroupResponse) SetHeaders(v map[string]*string) *QueryAllSwimmingLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponse) SetStatusCode(v int32) *QueryAllSwimmingLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponse) SetBody(v *QueryAllSwimmingLaneGroupResponseBody) *QueryAllSwimmingLaneGroupResponse {
	s.Body = v
	return s
}

func (s *QueryAllSwimmingLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

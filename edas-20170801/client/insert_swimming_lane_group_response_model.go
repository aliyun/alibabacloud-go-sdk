// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInsertSwimmingLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InsertSwimmingLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InsertSwimmingLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *InsertSwimmingLaneGroupResponseBody) *InsertSwimmingLaneGroupResponse
	GetBody() *InsertSwimmingLaneGroupResponseBody
}

type InsertSwimmingLaneGroupResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InsertSwimmingLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InsertSwimmingLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s InsertSwimmingLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *InsertSwimmingLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InsertSwimmingLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InsertSwimmingLaneGroupResponse) GetBody() *InsertSwimmingLaneGroupResponseBody {
	return s.Body
}

func (s *InsertSwimmingLaneGroupResponse) SetHeaders(v map[string]*string) *InsertSwimmingLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *InsertSwimmingLaneGroupResponse) SetStatusCode(v int32) *InsertSwimmingLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *InsertSwimmingLaneGroupResponse) SetBody(v *InsertSwimmingLaneGroupResponseBody) *InsertSwimmingLaneGroupResponse {
	s.Body = v
	return s
}

func (s *InsertSwimmingLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

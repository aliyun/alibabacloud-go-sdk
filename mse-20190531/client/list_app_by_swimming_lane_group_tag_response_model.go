// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppBySwimmingLaneGroupTagResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppBySwimmingLaneGroupTagResponse
	GetStatusCode() *int32
	SetBody(v *ListAppBySwimmingLaneGroupTagResponseBody) *ListAppBySwimmingLaneGroupTagResponse
	GetBody() *ListAppBySwimmingLaneGroupTagResponseBody
}

type ListAppBySwimmingLaneGroupTagResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppBySwimmingLaneGroupTagResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppBySwimmingLaneGroupTagResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagResponse) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppBySwimmingLaneGroupTagResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppBySwimmingLaneGroupTagResponse) GetBody() *ListAppBySwimmingLaneGroupTagResponseBody {
	return s.Body
}

func (s *ListAppBySwimmingLaneGroupTagResponse) SetHeaders(v map[string]*string) *ListAppBySwimmingLaneGroupTagResponse {
	s.Headers = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponse) SetStatusCode(v int32) *ListAppBySwimmingLaneGroupTagResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponse) SetBody(v *ListAppBySwimmingLaneGroupTagResponseBody) *ListAppBySwimmingLaneGroupTagResponse {
	s.Body = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagResponse) Validate() error {
	return dara.Validate(s)
}

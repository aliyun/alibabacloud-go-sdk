// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListApplicationsForSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListApplicationsForSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListApplicationsForSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *ListApplicationsForSwimmingLaneResponseBody) *ListApplicationsForSwimmingLaneResponse
	GetBody() *ListApplicationsForSwimmingLaneResponseBody
}

type ListApplicationsForSwimmingLaneResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListApplicationsForSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListApplicationsForSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListApplicationsForSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *ListApplicationsForSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListApplicationsForSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListApplicationsForSwimmingLaneResponse) GetBody() *ListApplicationsForSwimmingLaneResponseBody {
	return s.Body
}

func (s *ListApplicationsForSwimmingLaneResponse) SetHeaders(v map[string]*string) *ListApplicationsForSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponse) SetStatusCode(v int32) *ListApplicationsForSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponse) SetBody(v *ListApplicationsForSwimmingLaneResponseBody) *ListApplicationsForSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *ListApplicationsForSwimmingLaneResponse) Validate() error {
	return dara.Validate(s)
}

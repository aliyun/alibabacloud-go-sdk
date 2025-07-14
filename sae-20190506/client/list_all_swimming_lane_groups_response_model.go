// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAllSwimmingLaneGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAllSwimmingLaneGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAllSwimmingLaneGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListAllSwimmingLaneGroupsResponseBody) *ListAllSwimmingLaneGroupsResponse
	GetBody() *ListAllSwimmingLaneGroupsResponseBody
}

type ListAllSwimmingLaneGroupsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAllSwimmingLaneGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAllSwimmingLaneGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAllSwimmingLaneGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAllSwimmingLaneGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAllSwimmingLaneGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAllSwimmingLaneGroupsResponse) GetBody() *ListAllSwimmingLaneGroupsResponseBody {
	return s.Body
}

func (s *ListAllSwimmingLaneGroupsResponse) SetHeaders(v map[string]*string) *ListAllSwimmingLaneGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponse) SetStatusCode(v int32) *ListAllSwimmingLaneGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponse) SetBody(v *ListAllSwimmingLaneGroupsResponseBody) *ListAllSwimmingLaneGroupsResponse {
	s.Body = v
	return s
}

func (s *ListAllSwimmingLaneGroupsResponse) Validate() error {
	return dara.Validate(s)
}

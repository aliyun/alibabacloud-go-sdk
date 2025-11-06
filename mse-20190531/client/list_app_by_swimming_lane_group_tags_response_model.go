// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppBySwimmingLaneGroupTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAppBySwimmingLaneGroupTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAppBySwimmingLaneGroupTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListAppBySwimmingLaneGroupTagsResponseBody) *ListAppBySwimmingLaneGroupTagsResponse
	GetBody() *ListAppBySwimmingLaneGroupTagsResponseBody
}

type ListAppBySwimmingLaneGroupTagsResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAppBySwimmingLaneGroupTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAppBySwimmingLaneGroupTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAppBySwimmingLaneGroupTagsResponse) GoString() string {
	return s.String()
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) GetBody() *ListAppBySwimmingLaneGroupTagsResponseBody {
	return s.Body
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) SetHeaders(v map[string]*string) *ListAppBySwimmingLaneGroupTagsResponse {
	s.Headers = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) SetStatusCode(v int32) *ListAppBySwimmingLaneGroupTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) SetBody(v *ListAppBySwimmingLaneGroupTagsResponseBody) *ListAppBySwimmingLaneGroupTagsResponse {
	s.Body = v
	return s
}

func (s *ListAppBySwimmingLaneGroupTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

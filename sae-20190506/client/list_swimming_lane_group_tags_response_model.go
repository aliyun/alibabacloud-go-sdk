// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneGroupTagsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSwimmingLaneGroupTagsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSwimmingLaneGroupTagsResponse
	GetStatusCode() *int32
	SetBody(v *ListSwimmingLaneGroupTagsResponseBody) *ListSwimmingLaneGroupTagsResponse
	GetBody() *ListSwimmingLaneGroupTagsResponseBody
}

type ListSwimmingLaneGroupTagsResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSwimmingLaneGroupTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSwimmingLaneGroupTagsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneGroupTagsResponse) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneGroupTagsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSwimmingLaneGroupTagsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSwimmingLaneGroupTagsResponse) GetBody() *ListSwimmingLaneGroupTagsResponseBody {
	return s.Body
}

func (s *ListSwimmingLaneGroupTagsResponse) SetHeaders(v map[string]*string) *ListSwimmingLaneGroupTagsResponse {
	s.Headers = v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponse) SetStatusCode(v int32) *ListSwimmingLaneGroupTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponse) SetBody(v *ListSwimmingLaneGroupTagsResponseBody) *ListSwimmingLaneGroupTagsResponse {
	s.Body = v
	return s
}

func (s *ListSwimmingLaneGroupTagsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

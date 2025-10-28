// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *ListSwimmingLaneResponseBody) *ListSwimmingLaneResponse
	GetBody() *ListSwimmingLaneResponseBody
}

type ListSwimmingLaneResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *ListSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSwimmingLaneResponse) GetBody() *ListSwimmingLaneResponseBody {
	return s.Body
}

func (s *ListSwimmingLaneResponse) SetHeaders(v map[string]*string) *ListSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *ListSwimmingLaneResponse) SetStatusCode(v int32) *ListSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSwimmingLaneResponse) SetBody(v *ListSwimmingLaneResponseBody) *ListSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *ListSwimmingLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

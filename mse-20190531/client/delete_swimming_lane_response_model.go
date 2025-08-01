// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSwimmingLaneResponseBody) *DeleteSwimmingLaneResponse
	GetBody() *DeleteSwimmingLaneResponseBody
}

type DeleteSwimmingLaneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *DeleteSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSwimmingLaneResponse) GetBody() *DeleteSwimmingLaneResponseBody {
	return s.Body
}

func (s *DeleteSwimmingLaneResponse) SetHeaders(v map[string]*string) *DeleteSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *DeleteSwimmingLaneResponse) SetStatusCode(v int32) *DeleteSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSwimmingLaneResponse) SetBody(v *DeleteSwimmingLaneResponseBody) *DeleteSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *DeleteSwimmingLaneResponse) Validate() error {
	return dara.Validate(s)
}

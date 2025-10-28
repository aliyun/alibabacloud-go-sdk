// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimmingLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSwimmingLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSwimmingLaneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSwimmingLaneResponseBody) *UpdateSwimmingLaneResponse
	GetBody() *UpdateSwimmingLaneResponseBody
}

type UpdateSwimmingLaneResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSwimmingLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSwimmingLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimmingLaneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimmingLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSwimmingLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSwimmingLaneResponse) GetBody() *UpdateSwimmingLaneResponseBody {
	return s.Body
}

func (s *UpdateSwimmingLaneResponse) SetHeaders(v map[string]*string) *UpdateSwimmingLaneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimmingLaneResponse) SetStatusCode(v int32) *UpdateSwimmingLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimmingLaneResponse) SetBody(v *UpdateSwimmingLaneResponseBody) *UpdateSwimmingLaneResponse {
	s.Body = v
	return s
}

func (s *UpdateSwimmingLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

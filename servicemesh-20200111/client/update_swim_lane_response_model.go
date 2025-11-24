// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateSwimLaneResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateSwimLaneResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateSwimLaneResponse
	GetStatusCode() *int32
	SetBody(v *UpdateSwimLaneResponseBody) *UpdateSwimLaneResponse
	GetBody() *UpdateSwimLaneResponseBody
}

type UpdateSwimLaneResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateSwimLaneResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateSwimLaneResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateSwimLaneResponse) GoString() string {
	return s.String()
}

func (s *UpdateSwimLaneResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateSwimLaneResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateSwimLaneResponse) GetBody() *UpdateSwimLaneResponseBody {
	return s.Body
}

func (s *UpdateSwimLaneResponse) SetHeaders(v map[string]*string) *UpdateSwimLaneResponse {
	s.Headers = v
	return s
}

func (s *UpdateSwimLaneResponse) SetStatusCode(v int32) *UpdateSwimLaneResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateSwimLaneResponse) SetBody(v *UpdateSwimLaneResponseBody) *UpdateSwimLaneResponse {
	s.Body = v
	return s
}

func (s *UpdateSwimLaneResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

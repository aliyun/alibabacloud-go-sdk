// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSwimLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSwimLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSwimLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateSwimLaneGroupResponseBody) *CreateSwimLaneGroupResponse
	GetBody() *CreateSwimLaneGroupResponseBody
}

type CreateSwimLaneGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSwimLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSwimLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSwimLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateSwimLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSwimLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSwimLaneGroupResponse) GetBody() *CreateSwimLaneGroupResponseBody {
	return s.Body
}

func (s *CreateSwimLaneGroupResponse) SetHeaders(v map[string]*string) *CreateSwimLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateSwimLaneGroupResponse) SetStatusCode(v int32) *CreateSwimLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSwimLaneGroupResponse) SetBody(v *CreateSwimLaneGroupResponseBody) *CreateSwimLaneGroupResponse {
	s.Body = v
	return s
}

func (s *CreateSwimLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

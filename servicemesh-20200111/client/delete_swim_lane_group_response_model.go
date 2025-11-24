// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSwimLaneGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteSwimLaneGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteSwimLaneGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteSwimLaneGroupResponseBody) *DeleteSwimLaneGroupResponse
	GetBody() *DeleteSwimLaneGroupResponseBody
}

type DeleteSwimLaneGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteSwimLaneGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteSwimLaneGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteSwimLaneGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteSwimLaneGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteSwimLaneGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteSwimLaneGroupResponse) GetBody() *DeleteSwimLaneGroupResponseBody {
	return s.Body
}

func (s *DeleteSwimLaneGroupResponse) SetHeaders(v map[string]*string) *DeleteSwimLaneGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteSwimLaneGroupResponse) SetStatusCode(v int32) *DeleteSwimLaneGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteSwimLaneGroupResponse) SetBody(v *DeleteSwimLaneGroupResponseBody) *DeleteSwimLaneGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteSwimLaneGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

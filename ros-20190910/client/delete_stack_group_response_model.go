// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStackGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStackGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStackGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStackGroupResponseBody) *DeleteStackGroupResponse
	GetBody() *DeleteStackGroupResponseBody
}

type DeleteStackGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStackGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStackGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStackGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteStackGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStackGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStackGroupResponse) GetBody() *DeleteStackGroupResponseBody {
	return s.Body
}

func (s *DeleteStackGroupResponse) SetHeaders(v map[string]*string) *DeleteStackGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteStackGroupResponse) SetStatusCode(v int32) *DeleteStackGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStackGroupResponse) SetBody(v *DeleteStackGroupResponseBody) *DeleteStackGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteStackGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

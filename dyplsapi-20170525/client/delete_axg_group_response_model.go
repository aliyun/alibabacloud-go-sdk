// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAxgGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAxgGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAxgGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAxgGroupResponseBody) *DeleteAxgGroupResponse
	GetBody() *DeleteAxgGroupResponseBody
}

type DeleteAxgGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAxgGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAxgGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAxgGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteAxgGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAxgGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAxgGroupResponse) GetBody() *DeleteAxgGroupResponseBody {
	return s.Body
}

func (s *DeleteAxgGroupResponse) SetHeaders(v map[string]*string) *DeleteAxgGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteAxgGroupResponse) SetStatusCode(v int32) *DeleteAxgGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAxgGroupResponse) SetBody(v *DeleteAxgGroupResponseBody) *DeleteAxgGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteAxgGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

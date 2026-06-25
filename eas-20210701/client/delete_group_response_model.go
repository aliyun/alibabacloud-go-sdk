// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse
	GetBody() *DeleteGroupResponseBody
}

type DeleteGroupResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGroupResponse) GetBody() *DeleteGroupResponseBody {
	return s.Body
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupResponse) SetBody(v *DeleteGroupResponseBody) *DeleteGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

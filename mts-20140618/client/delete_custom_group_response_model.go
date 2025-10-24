// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCustomGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCustomGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCustomGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCustomGroupResponseBody) *DeleteCustomGroupResponse
	GetBody() *DeleteCustomGroupResponseBody
}

type DeleteCustomGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCustomGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCustomGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCustomGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteCustomGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCustomGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCustomGroupResponse) GetBody() *DeleteCustomGroupResponseBody {
	return s.Body
}

func (s *DeleteCustomGroupResponse) SetHeaders(v map[string]*string) *DeleteCustomGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteCustomGroupResponse) SetStatusCode(v int32) *DeleteCustomGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCustomGroupResponse) SetBody(v *DeleteCustomGroupResponseBody) *DeleteCustomGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteCustomGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

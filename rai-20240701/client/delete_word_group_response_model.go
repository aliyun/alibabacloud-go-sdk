// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWordGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteWordGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteWordGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteWordGroupResponseBody) *DeleteWordGroupResponse
	GetBody() *DeleteWordGroupResponseBody
}

type DeleteWordGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteWordGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteWordGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteWordGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteWordGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteWordGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteWordGroupResponse) GetBody() *DeleteWordGroupResponseBody {
	return s.Body
}

func (s *DeleteWordGroupResponse) SetHeaders(v map[string]*string) *DeleteWordGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteWordGroupResponse) SetStatusCode(v int32) *DeleteWordGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteWordGroupResponse) SetBody(v *DeleteWordGroupResponseBody) *DeleteWordGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteWordGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

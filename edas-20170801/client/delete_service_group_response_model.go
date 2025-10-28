// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteServiceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteServiceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteServiceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteServiceGroupResponseBody) *DeleteServiceGroupResponse
	GetBody() *DeleteServiceGroupResponseBody
}

type DeleteServiceGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteServiceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteServiceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteServiceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteServiceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteServiceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteServiceGroupResponse) GetBody() *DeleteServiceGroupResponseBody {
	return s.Body
}

func (s *DeleteServiceGroupResponse) SetHeaders(v map[string]*string) *DeleteServiceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteServiceGroupResponse) SetStatusCode(v int32) *DeleteServiceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteServiceGroupResponse) SetBody(v *DeleteServiceGroupResponseBody) *DeleteServiceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteServiceGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

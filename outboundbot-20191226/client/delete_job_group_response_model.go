// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteJobGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteJobGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteJobGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteJobGroupResponseBody) *DeleteJobGroupResponse
	GetBody() *DeleteJobGroupResponseBody
}

type DeleteJobGroupResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteJobGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteJobGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteJobGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteJobGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteJobGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteJobGroupResponse) GetBody() *DeleteJobGroupResponseBody {
	return s.Body
}

func (s *DeleteJobGroupResponse) SetHeaders(v map[string]*string) *DeleteJobGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteJobGroupResponse) SetStatusCode(v int32) *DeleteJobGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteJobGroupResponse) SetBody(v *DeleteJobGroupResponseBody) *DeleteJobGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteJobGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

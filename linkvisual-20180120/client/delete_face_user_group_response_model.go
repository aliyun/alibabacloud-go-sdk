// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceUserGroupResponseBody) *DeleteFaceUserGroupResponse
	GetBody() *DeleteFaceUserGroupResponseBody
}

type DeleteFaceUserGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceUserGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceUserGroupResponse) GetBody() *DeleteFaceUserGroupResponseBody {
	return s.Body
}

func (s *DeleteFaceUserGroupResponse) SetHeaders(v map[string]*string) *DeleteFaceUserGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceUserGroupResponse) SetStatusCode(v int32) *DeleteFaceUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceUserGroupResponse) SetBody(v *DeleteFaceUserGroupResponseBody) *DeleteFaceUserGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

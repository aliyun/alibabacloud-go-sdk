// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFaceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteFaceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteFaceGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteFaceGroupResponseBody) *DeleteFaceGroupResponse
	GetBody() *DeleteFaceGroupResponseBody
}

type DeleteFaceGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFaceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFaceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteFaceGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteFaceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteFaceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteFaceGroupResponse) GetBody() *DeleteFaceGroupResponseBody {
	return s.Body
}

func (s *DeleteFaceGroupResponse) SetHeaders(v map[string]*string) *DeleteFaceGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteFaceGroupResponse) SetStatusCode(v int32) *DeleteFaceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFaceGroupResponse) SetBody(v *DeleteFaceGroupResponseBody) *DeleteFaceGroupResponse {
	s.Body = v
	return s
}

func (s *DeleteFaceGroupResponse) Validate() error {
	return dara.Validate(s)
}

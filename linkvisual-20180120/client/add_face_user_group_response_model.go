// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddFaceUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddFaceUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddFaceUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *AddFaceUserGroupResponseBody) *AddFaceUserGroupResponse
	GetBody() *AddFaceUserGroupResponseBody
}

type AddFaceUserGroupResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddFaceUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddFaceUserGroupResponse) GoString() string {
	return s.String()
}

func (s *AddFaceUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddFaceUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddFaceUserGroupResponse) GetBody() *AddFaceUserGroupResponseBody {
	return s.Body
}

func (s *AddFaceUserGroupResponse) SetHeaders(v map[string]*string) *AddFaceUserGroupResponse {
	s.Headers = v
	return s
}

func (s *AddFaceUserGroupResponse) SetStatusCode(v int32) *AddFaceUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddFaceUserGroupResponse) SetBody(v *AddFaceUserGroupResponseBody) *AddFaceUserGroupResponse {
	s.Body = v
	return s
}

func (s *AddFaceUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

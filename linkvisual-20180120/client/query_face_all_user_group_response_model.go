// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceAllUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceAllUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceAllUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceAllUserGroupResponseBody) *QueryFaceAllUserGroupResponse
	GetBody() *QueryFaceAllUserGroupResponseBody
}

type QueryFaceAllUserGroupResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceAllUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceAllUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceAllUserGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceAllUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceAllUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceAllUserGroupResponse) GetBody() *QueryFaceAllUserGroupResponseBody {
	return s.Body
}

func (s *QueryFaceAllUserGroupResponse) SetHeaders(v map[string]*string) *QueryFaceAllUserGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceAllUserGroupResponse) SetStatusCode(v int32) *QueryFaceAllUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceAllUserGroupResponse) SetBody(v *QueryFaceAllUserGroupResponseBody) *QueryFaceAllUserGroupResponse {
	s.Body = v
	return s
}

func (s *QueryFaceAllUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

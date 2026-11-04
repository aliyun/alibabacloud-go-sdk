// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryFaceUserGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryFaceUserGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryFaceUserGroupResponse
	GetStatusCode() *int32
	SetBody(v *QueryFaceUserGroupResponseBody) *QueryFaceUserGroupResponse
	GetBody() *QueryFaceUserGroupResponseBody
}

type QueryFaceUserGroupResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryFaceUserGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryFaceUserGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryFaceUserGroupResponse) GoString() string {
	return s.String()
}

func (s *QueryFaceUserGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryFaceUserGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryFaceUserGroupResponse) GetBody() *QueryFaceUserGroupResponseBody {
	return s.Body
}

func (s *QueryFaceUserGroupResponse) SetHeaders(v map[string]*string) *QueryFaceUserGroupResponse {
	s.Headers = v
	return s
}

func (s *QueryFaceUserGroupResponse) SetStatusCode(v int32) *QueryFaceUserGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryFaceUserGroupResponse) SetBody(v *QueryFaceUserGroupResponseBody) *QueryFaceUserGroupResponse {
	s.Body = v
	return s
}

func (s *QueryFaceUserGroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

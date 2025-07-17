// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProjectMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteProjectMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteProjectMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteProjectMemberResponseBody) *DeleteProjectMemberResponse
	GetBody() *DeleteProjectMemberResponseBody
}

type DeleteProjectMemberResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteProjectMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteProjectMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteProjectMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteProjectMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteProjectMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteProjectMemberResponse) GetBody() *DeleteProjectMemberResponseBody {
	return s.Body
}

func (s *DeleteProjectMemberResponse) SetHeaders(v map[string]*string) *DeleteProjectMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteProjectMemberResponse) SetStatusCode(v int32) *DeleteProjectMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteProjectMemberResponse) SetBody(v *DeleteProjectMemberResponseBody) *DeleteProjectMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteProjectMemberResponse) Validate() error {
	return dara.Validate(s)
}

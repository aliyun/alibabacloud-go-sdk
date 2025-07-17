// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLhMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLhMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLhMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLhMembersResponseBody) *DeleteLhMembersResponse
	GetBody() *DeleteLhMembersResponseBody
}

type DeleteLhMembersResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLhMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLhMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLhMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteLhMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLhMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLhMembersResponse) GetBody() *DeleteLhMembersResponseBody {
	return s.Body
}

func (s *DeleteLhMembersResponse) SetHeaders(v map[string]*string) *DeleteLhMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteLhMembersResponse) SetStatusCode(v int32) *DeleteLhMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLhMembersResponse) SetBody(v *DeleteLhMembersResponseBody) *DeleteLhMembersResponse {
	s.Body = v
	return s
}

func (s *DeleteLhMembersResponse) Validate() error {
	return dara.Validate(s)
}

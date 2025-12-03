// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteGroupMemberResponseBody) *DeleteGroupMemberResponse
	GetBody() *DeleteGroupMemberResponseBody
}

type DeleteGroupMemberResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteGroupMemberResponse) GetBody() *DeleteGroupMemberResponseBody {
	return s.Body
}

func (s *DeleteGroupMemberResponse) SetHeaders(v map[string]*string) *DeleteGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupMemberResponse) SetStatusCode(v int32) *DeleteGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteGroupMemberResponse) SetBody(v *DeleteGroupMemberResponseBody) *DeleteGroupMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

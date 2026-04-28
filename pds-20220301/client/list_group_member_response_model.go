// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGroupMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListGroupMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListGroupMemberResponse
	GetStatusCode() *int32
	SetBody(v *ListGroupMemberResponseBody) *ListGroupMemberResponse
	GetBody() *ListGroupMemberResponseBody
}

type ListGroupMemberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s ListGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListGroupMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListGroupMemberResponse) GetBody() *ListGroupMemberResponseBody {
	return s.Body
}

func (s *ListGroupMemberResponse) SetHeaders(v map[string]*string) *ListGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *ListGroupMemberResponse) SetStatusCode(v int32) *ListGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupMemberResponse) SetBody(v *ListGroupMemberResponseBody) *ListGroupMemberResponse {
	s.Body = v
	return s
}

func (s *ListGroupMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

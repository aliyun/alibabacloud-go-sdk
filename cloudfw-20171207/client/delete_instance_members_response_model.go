// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteInstanceMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteInstanceMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteInstanceMembersResponseBody) *DeleteInstanceMembersResponse
	GetBody() *DeleteInstanceMembersResponseBody
}

type DeleteInstanceMembersResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteInstanceMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteInstanceMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteInstanceMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteInstanceMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteInstanceMembersResponse) GetBody() *DeleteInstanceMembersResponseBody {
	return s.Body
}

func (s *DeleteInstanceMembersResponse) SetHeaders(v map[string]*string) *DeleteInstanceMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteInstanceMembersResponse) SetStatusCode(v int32) *DeleteInstanceMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteInstanceMembersResponse) SetBody(v *DeleteInstanceMembersResponseBody) *DeleteInstanceMembersResponse {
	s.Body = v
	return s
}

func (s *DeleteInstanceMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

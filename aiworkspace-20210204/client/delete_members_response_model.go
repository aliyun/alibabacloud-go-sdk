// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMembersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMembersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMembersResponse
	GetStatusCode() *int32
	SetBody(v *DeleteMembersResponseBody) *DeleteMembersResponse
	GetBody() *DeleteMembersResponseBody
}

type DeleteMembersResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteMembersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteMembersResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMembersResponse) GoString() string {
	return s.String()
}

func (s *DeleteMembersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMembersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMembersResponse) GetBody() *DeleteMembersResponseBody {
	return s.Body
}

func (s *DeleteMembersResponse) SetHeaders(v map[string]*string) *DeleteMembersResponse {
	s.Headers = v
	return s
}

func (s *DeleteMembersResponse) SetStatusCode(v int32) *DeleteMembersResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMembersResponse) SetBody(v *DeleteMembersResponseBody) *DeleteMembersResponse {
	s.Body = v
	return s
}

func (s *DeleteMembersResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

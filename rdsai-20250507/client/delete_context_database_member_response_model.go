// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContextDatabaseMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteContextDatabaseMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteContextDatabaseMemberResponse
	GetStatusCode() *int32
	SetBody(v *DeleteContextDatabaseMemberResponseBody) *DeleteContextDatabaseMemberResponse
	GetBody() *DeleteContextDatabaseMemberResponseBody
}

type DeleteContextDatabaseMemberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteContextDatabaseMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteContextDatabaseMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteContextDatabaseMemberResponse) GoString() string {
	return s.String()
}

func (s *DeleteContextDatabaseMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteContextDatabaseMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteContextDatabaseMemberResponse) GetBody() *DeleteContextDatabaseMemberResponseBody {
	return s.Body
}

func (s *DeleteContextDatabaseMemberResponse) SetHeaders(v map[string]*string) *DeleteContextDatabaseMemberResponse {
	s.Headers = v
	return s
}

func (s *DeleteContextDatabaseMemberResponse) SetStatusCode(v int32) *DeleteContextDatabaseMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteContextDatabaseMemberResponse) SetBody(v *DeleteContextDatabaseMemberResponseBody) *DeleteContextDatabaseMemberResponse {
	s.Body = v
	return s
}

func (s *DeleteContextDatabaseMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

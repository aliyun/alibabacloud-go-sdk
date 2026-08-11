// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateContextDatabaseMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateContextDatabaseMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateContextDatabaseMemberResponse
	GetStatusCode() *int32
	SetBody(v *CreateContextDatabaseMemberResponseBody) *CreateContextDatabaseMemberResponse
	GetBody() *CreateContextDatabaseMemberResponseBody
}

type CreateContextDatabaseMemberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateContextDatabaseMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateContextDatabaseMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateContextDatabaseMemberResponse) GoString() string {
	return s.String()
}

func (s *CreateContextDatabaseMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateContextDatabaseMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateContextDatabaseMemberResponse) GetBody() *CreateContextDatabaseMemberResponseBody {
	return s.Body
}

func (s *CreateContextDatabaseMemberResponse) SetHeaders(v map[string]*string) *CreateContextDatabaseMemberResponse {
	s.Headers = v
	return s
}

func (s *CreateContextDatabaseMemberResponse) SetStatusCode(v int32) *CreateContextDatabaseMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateContextDatabaseMemberResponse) SetBody(v *CreateContextDatabaseMemberResponseBody) *CreateContextDatabaseMemberResponse {
	s.Body = v
	return s
}

func (s *CreateContextDatabaseMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

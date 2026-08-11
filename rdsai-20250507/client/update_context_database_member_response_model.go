// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateContextDatabaseMemberResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateContextDatabaseMemberResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateContextDatabaseMemberResponse
	GetStatusCode() *int32
	SetBody(v *UpdateContextDatabaseMemberResponseBody) *UpdateContextDatabaseMemberResponse
	GetBody() *UpdateContextDatabaseMemberResponseBody
}

type UpdateContextDatabaseMemberResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateContextDatabaseMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateContextDatabaseMemberResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateContextDatabaseMemberResponse) GoString() string {
	return s.String()
}

func (s *UpdateContextDatabaseMemberResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateContextDatabaseMemberResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateContextDatabaseMemberResponse) GetBody() *UpdateContextDatabaseMemberResponseBody {
	return s.Body
}

func (s *UpdateContextDatabaseMemberResponse) SetHeaders(v map[string]*string) *UpdateContextDatabaseMemberResponse {
	s.Headers = v
	return s
}

func (s *UpdateContextDatabaseMemberResponse) SetStatusCode(v int32) *UpdateContextDatabaseMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateContextDatabaseMemberResponse) SetBody(v *UpdateContextDatabaseMemberResponseBody) *UpdateContextDatabaseMemberResponse {
	s.Body = v
	return s
}

func (s *UpdateContextDatabaseMemberResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

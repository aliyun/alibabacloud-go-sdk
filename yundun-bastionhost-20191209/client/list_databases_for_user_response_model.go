// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabasesForUserResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabasesForUserResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabasesForUserResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabasesForUserResponseBody) *ListDatabasesForUserResponse
	GetBody() *ListDatabasesForUserResponseBody
}

type ListDatabasesForUserResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabasesForUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabasesForUserResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabasesForUserResponse) GoString() string {
	return s.String()
}

func (s *ListDatabasesForUserResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabasesForUserResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabasesForUserResponse) GetBody() *ListDatabasesForUserResponseBody {
	return s.Body
}

func (s *ListDatabasesForUserResponse) SetHeaders(v map[string]*string) *ListDatabasesForUserResponse {
	s.Headers = v
	return s
}

func (s *ListDatabasesForUserResponse) SetStatusCode(v int32) *ListDatabasesForUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabasesForUserResponse) SetBody(v *ListDatabasesForUserResponseBody) *ListDatabasesForUserResponse {
	s.Body = v
	return s
}

func (s *ListDatabasesForUserResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

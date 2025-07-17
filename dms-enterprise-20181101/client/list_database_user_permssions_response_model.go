// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseUserPermssionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabaseUserPermssionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabaseUserPermssionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabaseUserPermssionsResponseBody) *ListDatabaseUserPermssionsResponse
	GetBody() *ListDatabaseUserPermssionsResponseBody
}

type ListDatabaseUserPermssionsResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseUserPermssionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseUserPermssionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseUserPermssionsResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseUserPermssionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabaseUserPermssionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabaseUserPermssionsResponse) GetBody() *ListDatabaseUserPermssionsResponseBody {
	return s.Body
}

func (s *ListDatabaseUserPermssionsResponse) SetHeaders(v map[string]*string) *ListDatabaseUserPermssionsResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseUserPermssionsResponse) SetStatusCode(v int32) *ListDatabaseUserPermssionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseUserPermssionsResponse) SetBody(v *ListDatabaseUserPermssionsResponseBody) *ListDatabaseUserPermssionsResponse {
	s.Body = v
	return s
}

func (s *ListDatabaseUserPermssionsResponse) Validate() error {
	return dara.Validate(s)
}

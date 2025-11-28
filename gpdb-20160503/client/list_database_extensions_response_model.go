// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseExtensionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDatabaseExtensionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDatabaseExtensionsResponse
	GetStatusCode() *int32
	SetBody(v *ListDatabaseExtensionsResponseBody) *ListDatabaseExtensionsResponse
	GetBody() *ListDatabaseExtensionsResponseBody
}

type ListDatabaseExtensionsResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDatabaseExtensionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDatabaseExtensionsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseExtensionsResponse) GoString() string {
	return s.String()
}

func (s *ListDatabaseExtensionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDatabaseExtensionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDatabaseExtensionsResponse) GetBody() *ListDatabaseExtensionsResponseBody {
	return s.Body
}

func (s *ListDatabaseExtensionsResponse) SetHeaders(v map[string]*string) *ListDatabaseExtensionsResponse {
	s.Headers = v
	return s
}

func (s *ListDatabaseExtensionsResponse) SetStatusCode(v int32) *ListDatabaseExtensionsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDatabaseExtensionsResponse) SetBody(v *ListDatabaseExtensionsResponseBody) *ListDatabaseExtensionsResponse {
	s.Body = v
	return s
}

func (s *ListDatabaseExtensionsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

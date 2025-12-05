// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListOperationDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListOperationDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListOperationDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *ListOperationDatabasesResponseBody) *ListOperationDatabasesResponse
	GetBody() *ListOperationDatabasesResponseBody
}

type ListOperationDatabasesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListOperationDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListOperationDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListOperationDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListOperationDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListOperationDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListOperationDatabasesResponse) GetBody() *ListOperationDatabasesResponseBody {
	return s.Body
}

func (s *ListOperationDatabasesResponse) SetHeaders(v map[string]*string) *ListOperationDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListOperationDatabasesResponse) SetStatusCode(v int32) *ListOperationDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListOperationDatabasesResponse) SetBody(v *ListOperationDatabasesResponseBody) *ListOperationDatabasesResponse {
	s.Body = v
	return s
}

func (s *ListOperationDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstanceDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstanceDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *ListInstanceDatabasesResponseBody) *ListInstanceDatabasesResponse
	GetBody() *ListInstanceDatabasesResponseBody
}

type ListInstanceDatabasesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstanceDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstanceDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListInstanceDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstanceDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstanceDatabasesResponse) GetBody() *ListInstanceDatabasesResponseBody {
	return s.Body
}

func (s *ListInstanceDatabasesResponse) SetHeaders(v map[string]*string) *ListInstanceDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListInstanceDatabasesResponse) SetStatusCode(v int32) *ListInstanceDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstanceDatabasesResponse) SetBody(v *ListInstanceDatabasesResponseBody) *ListInstanceDatabasesResponse {
	s.Body = v
	return s
}

func (s *ListInstanceDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

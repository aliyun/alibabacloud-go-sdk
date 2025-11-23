// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLogicDatabasesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListLogicDatabasesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListLogicDatabasesResponse
	GetStatusCode() *int32
	SetBody(v *ListLogicDatabasesResponseBody) *ListLogicDatabasesResponse
	GetBody() *ListLogicDatabasesResponseBody
}

type ListLogicDatabasesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListLogicDatabasesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListLogicDatabasesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListLogicDatabasesResponse) GoString() string {
	return s.String()
}

func (s *ListLogicDatabasesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListLogicDatabasesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListLogicDatabasesResponse) GetBody() *ListLogicDatabasesResponseBody {
	return s.Body
}

func (s *ListLogicDatabasesResponse) SetHeaders(v map[string]*string) *ListLogicDatabasesResponse {
	s.Headers = v
	return s
}

func (s *ListLogicDatabasesResponse) SetStatusCode(v int32) *ListLogicDatabasesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListLogicDatabasesResponse) SetBody(v *ListLogicDatabasesResponseBody) *ListLogicDatabasesResponse {
	s.Body = v
	return s
}

func (s *ListLogicDatabasesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

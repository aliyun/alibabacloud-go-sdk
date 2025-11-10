// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationOperationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMigrationOperationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMigrationOperationsResponse
	GetStatusCode() *int32
	SetBody(v *ListMigrationOperationsResponseBody) *ListMigrationOperationsResponse
	GetBody() *ListMigrationOperationsResponseBody
}

type ListMigrationOperationsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMigrationOperationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMigrationOperationsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationOperationsResponse) GoString() string {
	return s.String()
}

func (s *ListMigrationOperationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMigrationOperationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMigrationOperationsResponse) GetBody() *ListMigrationOperationsResponseBody {
	return s.Body
}

func (s *ListMigrationOperationsResponse) SetHeaders(v map[string]*string) *ListMigrationOperationsResponse {
	s.Headers = v
	return s
}

func (s *ListMigrationOperationsResponse) SetStatusCode(v int32) *ListMigrationOperationsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMigrationOperationsResponse) SetBody(v *ListMigrationOperationsResponseBody) *ListMigrationOperationsResponse {
	s.Body = v
	return s
}

func (s *ListMigrationOperationsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

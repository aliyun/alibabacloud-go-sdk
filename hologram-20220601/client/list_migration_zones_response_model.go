// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMigrationZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMigrationZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMigrationZonesResponse
	GetStatusCode() *int32
	SetBody(v *ListMigrationZonesResponseBody) *ListMigrationZonesResponse
	GetBody() *ListMigrationZonesResponseBody
}

type ListMigrationZonesResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMigrationZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMigrationZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMigrationZonesResponse) GoString() string {
	return s.String()
}

func (s *ListMigrationZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMigrationZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMigrationZonesResponse) GetBody() *ListMigrationZonesResponseBody {
	return s.Body
}

func (s *ListMigrationZonesResponse) SetHeaders(v map[string]*string) *ListMigrationZonesResponse {
	s.Headers = v
	return s
}

func (s *ListMigrationZonesResponse) SetStatusCode(v int32) *ListMigrationZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMigrationZonesResponse) SetBody(v *ListMigrationZonesResponseBody) *ListMigrationZonesResponse {
	s.Body = v
	return s
}

func (s *ListMigrationZonesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

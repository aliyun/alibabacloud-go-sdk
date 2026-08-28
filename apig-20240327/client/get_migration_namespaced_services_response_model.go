// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMigrationNamespacedServicesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetMigrationNamespacedServicesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetMigrationNamespacedServicesResponse
	GetStatusCode() *int32
	SetBody(v *GetMigrationNamespacedServicesResponseBody) *GetMigrationNamespacedServicesResponse
	GetBody() *GetMigrationNamespacedServicesResponseBody
}

type GetMigrationNamespacedServicesResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetMigrationNamespacedServicesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetMigrationNamespacedServicesResponse) String() string {
	return dara.Prettify(s)
}

func (s GetMigrationNamespacedServicesResponse) GoString() string {
	return s.String()
}

func (s *GetMigrationNamespacedServicesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetMigrationNamespacedServicesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetMigrationNamespacedServicesResponse) GetBody() *GetMigrationNamespacedServicesResponseBody {
	return s.Body
}

func (s *GetMigrationNamespacedServicesResponse) SetHeaders(v map[string]*string) *GetMigrationNamespacedServicesResponse {
	s.Headers = v
	return s
}

func (s *GetMigrationNamespacedServicesResponse) SetStatusCode(v int32) *GetMigrationNamespacedServicesResponse {
	s.StatusCode = &v
	return s
}

func (s *GetMigrationNamespacedServicesResponse) SetBody(v *GetMigrationNamespacedServicesResponseBody) *GetMigrationNamespacedServicesResponse {
	s.Body = v
	return s
}

func (s *GetMigrationNamespacedServicesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

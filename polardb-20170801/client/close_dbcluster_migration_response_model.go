// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDBClusterMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseDBClusterMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseDBClusterMigrationResponse
	GetStatusCode() *int32
	SetBody(v *CloseDBClusterMigrationResponseBody) *CloseDBClusterMigrationResponse
	GetBody() *CloseDBClusterMigrationResponseBody
}

type CloseDBClusterMigrationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseDBClusterMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseDBClusterMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseDBClusterMigrationResponse) GetBody() *CloseDBClusterMigrationResponseBody {
	return s.Body
}

func (s *CloseDBClusterMigrationResponse) SetHeaders(v map[string]*string) *CloseDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *CloseDBClusterMigrationResponse) SetStatusCode(v int32) *CloseDBClusterMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseDBClusterMigrationResponse) SetBody(v *CloseDBClusterMigrationResponseBody) *CloseDBClusterMigrationResponse {
	s.Body = v
	return s
}

func (s *CloseDBClusterMigrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

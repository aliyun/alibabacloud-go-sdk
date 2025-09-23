// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDBClusterMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ContinueDBClusterMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ContinueDBClusterMigrationResponse
	GetStatusCode() *int32
	SetBody(v *ContinueDBClusterMigrationResponseBody) *ContinueDBClusterMigrationResponse
	GetBody() *ContinueDBClusterMigrationResponseBody
}

type ContinueDBClusterMigrationResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ContinueDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ContinueDBClusterMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s ContinueDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *ContinueDBClusterMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ContinueDBClusterMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ContinueDBClusterMigrationResponse) GetBody() *ContinueDBClusterMigrationResponseBody {
	return s.Body
}

func (s *ContinueDBClusterMigrationResponse) SetHeaders(v map[string]*string) *ContinueDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *ContinueDBClusterMigrationResponse) SetStatusCode(v int32) *ContinueDBClusterMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ContinueDBClusterMigrationResponse) SetBody(v *ContinueDBClusterMigrationResponseBody) *ContinueDBClusterMigrationResponse {
	s.Body = v
	return s
}

func (s *ContinueDBClusterMigrationResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAbortDBClusterMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AbortDBClusterMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AbortDBClusterMigrationResponse
	GetStatusCode() *int32
	SetBody(v *AbortDBClusterMigrationResponseBody) *AbortDBClusterMigrationResponse
	GetBody() *AbortDBClusterMigrationResponseBody
}

type AbortDBClusterMigrationResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AbortDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AbortDBClusterMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s AbortDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *AbortDBClusterMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AbortDBClusterMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AbortDBClusterMigrationResponse) GetBody() *AbortDBClusterMigrationResponseBody {
	return s.Body
}

func (s *AbortDBClusterMigrationResponse) SetHeaders(v map[string]*string) *AbortDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *AbortDBClusterMigrationResponse) SetStatusCode(v int32) *AbortDBClusterMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *AbortDBClusterMigrationResponse) SetBody(v *AbortDBClusterMigrationResponseBody) *AbortDBClusterMigrationResponse {
	s.Body = v
	return s
}

func (s *AbortDBClusterMigrationResponse) Validate() error {
	return dara.Validate(s)
}

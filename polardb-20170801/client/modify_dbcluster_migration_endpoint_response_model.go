// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMigrationEndpointResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterMigrationEndpointResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterMigrationEndpointResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterMigrationEndpointResponseBody) *ModifyDBClusterMigrationEndpointResponse
	GetBody() *ModifyDBClusterMigrationEndpointResponseBody
}

type ModifyDBClusterMigrationEndpointResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterMigrationEndpointResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterMigrationEndpointResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMigrationEndpointResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationEndpointResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterMigrationEndpointResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterMigrationEndpointResponse) GetBody() *ModifyDBClusterMigrationEndpointResponseBody {
	return s.Body
}

func (s *ModifyDBClusterMigrationEndpointResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMigrationEndpointResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMigrationEndpointResponse) SetStatusCode(v int32) *ModifyDBClusterMigrationEndpointResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterMigrationEndpointResponse) SetBody(v *ModifyDBClusterMigrationEndpointResponseBody) *ModifyDBClusterMigrationEndpointResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterMigrationEndpointResponse) Validate() error {
	return dara.Validate(s)
}

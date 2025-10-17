// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBClusterMigrationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyDBClusterMigrationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyDBClusterMigrationResponse
	GetStatusCode() *int32
	SetBody(v *ModifyDBClusterMigrationResponseBody) *ModifyDBClusterMigrationResponse
	GetBody() *ModifyDBClusterMigrationResponseBody
}

type ModifyDBClusterMigrationResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyDBClusterMigrationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyDBClusterMigrationResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBClusterMigrationResponse) GoString() string {
	return s.String()
}

func (s *ModifyDBClusterMigrationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyDBClusterMigrationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyDBClusterMigrationResponse) GetBody() *ModifyDBClusterMigrationResponseBody {
	return s.Body
}

func (s *ModifyDBClusterMigrationResponse) SetHeaders(v map[string]*string) *ModifyDBClusterMigrationResponse {
	s.Headers = v
	return s
}

func (s *ModifyDBClusterMigrationResponse) SetStatusCode(v int32) *ModifyDBClusterMigrationResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyDBClusterMigrationResponse) SetBody(v *ModifyDBClusterMigrationResponseBody) *ModifyDBClusterMigrationResponse {
	s.Body = v
	return s
}

func (s *ModifyDBClusterMigrationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateDBClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateDBClusterResponse
	GetStatusCode() *int32
	SetBody(v *MigrateDBClusterResponseBody) *MigrateDBClusterResponse
	GetBody() *MigrateDBClusterResponseBody
}

type MigrateDBClusterResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateDBClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateDBClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBClusterResponse) GoString() string {
	return s.String()
}

func (s *MigrateDBClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateDBClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateDBClusterResponse) GetBody() *MigrateDBClusterResponseBody {
	return s.Body
}

func (s *MigrateDBClusterResponse) SetHeaders(v map[string]*string) *MigrateDBClusterResponse {
	s.Headers = v
	return s
}

func (s *MigrateDBClusterResponse) SetStatusCode(v int32) *MigrateDBClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateDBClusterResponse) SetBody(v *MigrateDBClusterResponseBody) *MigrateDBClusterResponse {
	s.Body = v
	return s
}

func (s *MigrateDBClusterResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

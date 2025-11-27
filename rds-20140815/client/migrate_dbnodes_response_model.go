// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateDBNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateDBNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateDBNodesResponse
	GetStatusCode() *int32
	SetBody(v *MigrateDBNodesResponseBody) *MigrateDBNodesResponse
	GetBody() *MigrateDBNodesResponseBody
}

type MigrateDBNodesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateDBNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateDBNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateDBNodesResponse) GoString() string {
	return s.String()
}

func (s *MigrateDBNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateDBNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateDBNodesResponse) GetBody() *MigrateDBNodesResponseBody {
	return s.Body
}

func (s *MigrateDBNodesResponse) SetHeaders(v map[string]*string) *MigrateDBNodesResponse {
	s.Headers = v
	return s
}

func (s *MigrateDBNodesResponse) SetStatusCode(v int32) *MigrateDBNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateDBNodesResponse) SetBody(v *MigrateDBNodesResponseBody) *MigrateDBNodesResponse {
	s.Body = v
	return s
}

func (s *MigrateDBNodesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

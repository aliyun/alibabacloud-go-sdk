// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMigrateClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MigrateClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MigrateClusterResponse
	GetStatusCode() *int32
	SetBody(v *MigrateClusterResponseBody) *MigrateClusterResponse
	GetBody() *MigrateClusterResponseBody
}

type MigrateClusterResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MigrateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MigrateClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s MigrateClusterResponse) GoString() string {
	return s.String()
}

func (s *MigrateClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MigrateClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MigrateClusterResponse) GetBody() *MigrateClusterResponseBody {
	return s.Body
}

func (s *MigrateClusterResponse) SetHeaders(v map[string]*string) *MigrateClusterResponse {
	s.Headers = v
	return s
}

func (s *MigrateClusterResponse) SetStatusCode(v int32) *MigrateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *MigrateClusterResponse) SetBody(v *MigrateClusterResponseBody) *MigrateClusterResponse {
	s.Body = v
	return s
}

func (s *MigrateClusterResponse) Validate() error {
	return dara.Validate(s)
}

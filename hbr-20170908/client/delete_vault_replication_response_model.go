// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVaultReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVaultReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVaultReplicationResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVaultReplicationResponseBody) *DeleteVaultReplicationResponse
	GetBody() *DeleteVaultReplicationResponseBody
}

type DeleteVaultReplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVaultReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVaultReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVaultReplicationResponse) GoString() string {
	return s.String()
}

func (s *DeleteVaultReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVaultReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVaultReplicationResponse) GetBody() *DeleteVaultReplicationResponseBody {
	return s.Body
}

func (s *DeleteVaultReplicationResponse) SetHeaders(v map[string]*string) *DeleteVaultReplicationResponse {
	s.Headers = v
	return s
}

func (s *DeleteVaultReplicationResponse) SetStatusCode(v int32) *DeleteVaultReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVaultReplicationResponse) SetBody(v *DeleteVaultReplicationResponseBody) *DeleteVaultReplicationResponse {
	s.Body = v
	return s
}

func (s *DeleteVaultReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

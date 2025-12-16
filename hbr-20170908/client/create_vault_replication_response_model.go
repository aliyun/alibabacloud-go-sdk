// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateVaultReplicationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVaultReplicationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVaultReplicationResponse
	GetStatusCode() *int32
	SetBody(v *CreateVaultReplicationResponseBody) *CreateVaultReplicationResponse
	GetBody() *CreateVaultReplicationResponseBody
}

type CreateVaultReplicationResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVaultReplicationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVaultReplicationResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVaultReplicationResponse) GoString() string {
	return s.String()
}

func (s *CreateVaultReplicationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVaultReplicationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVaultReplicationResponse) GetBody() *CreateVaultReplicationResponseBody {
	return s.Body
}

func (s *CreateVaultReplicationResponse) SetHeaders(v map[string]*string) *CreateVaultReplicationResponse {
	s.Headers = v
	return s
}

func (s *CreateVaultReplicationResponse) SetStatusCode(v int32) *CreateVaultReplicationResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVaultReplicationResponse) SetBody(v *CreateVaultReplicationResponseBody) *CreateVaultReplicationResponse {
	s.Body = v
	return s
}

func (s *CreateVaultReplicationResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

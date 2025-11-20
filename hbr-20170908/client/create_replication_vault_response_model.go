// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateReplicationVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateReplicationVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateReplicationVaultResponse
	GetStatusCode() *int32
	SetBody(v *CreateReplicationVaultResponseBody) *CreateReplicationVaultResponse
	GetBody() *CreateReplicationVaultResponseBody
}

type CreateReplicationVaultResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateReplicationVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateReplicationVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateReplicationVaultResponse) GoString() string {
	return s.String()
}

func (s *CreateReplicationVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateReplicationVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateReplicationVaultResponse) GetBody() *CreateReplicationVaultResponseBody {
	return s.Body
}

func (s *CreateReplicationVaultResponse) SetHeaders(v map[string]*string) *CreateReplicationVaultResponse {
	s.Headers = v
	return s
}

func (s *CreateReplicationVaultResponse) SetStatusCode(v int32) *CreateReplicationVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateReplicationVaultResponse) SetBody(v *CreateReplicationVaultResponseBody) *CreateReplicationVaultResponse {
	s.Body = v
	return s
}

func (s *CreateReplicationVaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

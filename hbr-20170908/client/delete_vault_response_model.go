// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVaultResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVaultResponseBody) *DeleteVaultResponse
	GetBody() *DeleteVaultResponseBody
}

type DeleteVaultResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVaultResponse) GoString() string {
	return s.String()
}

func (s *DeleteVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVaultResponse) GetBody() *DeleteVaultResponseBody {
	return s.Body
}

func (s *DeleteVaultResponse) SetHeaders(v map[string]*string) *DeleteVaultResponse {
	s.Headers = v
	return s
}

func (s *DeleteVaultResponse) SetStatusCode(v int32) *DeleteVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVaultResponse) SetBody(v *DeleteVaultResponseBody) *DeleteVaultResponse {
	s.Body = v
	return s
}

func (s *DeleteVaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

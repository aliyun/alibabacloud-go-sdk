// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTokenVaultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteTokenVaultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteTokenVaultResponse
	GetStatusCode() *int32
	SetBody(v *DeleteTokenVaultResponseBody) *DeleteTokenVaultResponse
	GetBody() *DeleteTokenVaultResponseBody
}

type DeleteTokenVaultResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteTokenVaultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteTokenVaultResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteTokenVaultResponse) GoString() string {
	return s.String()
}

func (s *DeleteTokenVaultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteTokenVaultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteTokenVaultResponse) GetBody() *DeleteTokenVaultResponseBody {
	return s.Body
}

func (s *DeleteTokenVaultResponse) SetHeaders(v map[string]*string) *DeleteTokenVaultResponse {
	s.Headers = v
	return s
}

func (s *DeleteTokenVaultResponse) SetStatusCode(v int32) *DeleteTokenVaultResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteTokenVaultResponse) SetBody(v *DeleteTokenVaultResponseBody) *DeleteTokenVaultResponse {
	s.Body = v
	return s
}

func (s *DeleteTokenVaultResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

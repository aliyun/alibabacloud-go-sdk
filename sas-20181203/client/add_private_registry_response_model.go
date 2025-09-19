// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddPrivateRegistryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddPrivateRegistryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddPrivateRegistryResponse
	GetStatusCode() *int32
	SetBody(v *AddPrivateRegistryResponseBody) *AddPrivateRegistryResponse
	GetBody() *AddPrivateRegistryResponseBody
}

type AddPrivateRegistryResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddPrivateRegistryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddPrivateRegistryResponse) String() string {
	return dara.Prettify(s)
}

func (s AddPrivateRegistryResponse) GoString() string {
	return s.String()
}

func (s *AddPrivateRegistryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddPrivateRegistryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddPrivateRegistryResponse) GetBody() *AddPrivateRegistryResponseBody {
	return s.Body
}

func (s *AddPrivateRegistryResponse) SetHeaders(v map[string]*string) *AddPrivateRegistryResponse {
	s.Headers = v
	return s
}

func (s *AddPrivateRegistryResponse) SetStatusCode(v int32) *AddPrivateRegistryResponse {
	s.StatusCode = &v
	return s
}

func (s *AddPrivateRegistryResponse) SetBody(v *AddPrivateRegistryResponseBody) *AddPrivateRegistryResponse {
	s.Body = v
	return s
}

func (s *AddPrivateRegistryResponse) Validate() error {
	return dara.Validate(s)
}

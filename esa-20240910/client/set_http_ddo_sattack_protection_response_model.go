// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetHttpDDoSAttackProtectionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetHttpDDoSAttackProtectionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetHttpDDoSAttackProtectionResponse
	GetStatusCode() *int32
	SetBody(v *SetHttpDDoSAttackProtectionResponseBody) *SetHttpDDoSAttackProtectionResponse
	GetBody() *SetHttpDDoSAttackProtectionResponseBody
}

type SetHttpDDoSAttackProtectionResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetHttpDDoSAttackProtectionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetHttpDDoSAttackProtectionResponse) String() string {
	return dara.Prettify(s)
}

func (s SetHttpDDoSAttackProtectionResponse) GoString() string {
	return s.String()
}

func (s *SetHttpDDoSAttackProtectionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetHttpDDoSAttackProtectionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetHttpDDoSAttackProtectionResponse) GetBody() *SetHttpDDoSAttackProtectionResponseBody {
	return s.Body
}

func (s *SetHttpDDoSAttackProtectionResponse) SetHeaders(v map[string]*string) *SetHttpDDoSAttackProtectionResponse {
	s.Headers = v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponse) SetStatusCode(v int32) *SetHttpDDoSAttackProtectionResponse {
	s.StatusCode = &v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponse) SetBody(v *SetHttpDDoSAttackProtectionResponseBody) *SetHttpDDoSAttackProtectionResponse {
	s.Body = v
	return s
}

func (s *SetHttpDDoSAttackProtectionResponse) Validate() error {
	return dara.Validate(s)
}

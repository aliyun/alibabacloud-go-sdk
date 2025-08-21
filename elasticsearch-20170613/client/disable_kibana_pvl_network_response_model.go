// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableKibanaPvlNetworkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisableKibanaPvlNetworkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisableKibanaPvlNetworkResponse
	GetStatusCode() *int32
	SetBody(v *DisableKibanaPvlNetworkResponseBody) *DisableKibanaPvlNetworkResponse
	GetBody() *DisableKibanaPvlNetworkResponseBody
}

type DisableKibanaPvlNetworkResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisableKibanaPvlNetworkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisableKibanaPvlNetworkResponse) String() string {
	return dara.Prettify(s)
}

func (s DisableKibanaPvlNetworkResponse) GoString() string {
	return s.String()
}

func (s *DisableKibanaPvlNetworkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisableKibanaPvlNetworkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisableKibanaPvlNetworkResponse) GetBody() *DisableKibanaPvlNetworkResponseBody {
	return s.Body
}

func (s *DisableKibanaPvlNetworkResponse) SetHeaders(v map[string]*string) *DisableKibanaPvlNetworkResponse {
	s.Headers = v
	return s
}

func (s *DisableKibanaPvlNetworkResponse) SetStatusCode(v int32) *DisableKibanaPvlNetworkResponse {
	s.StatusCode = &v
	return s
}

func (s *DisableKibanaPvlNetworkResponse) SetBody(v *DisableKibanaPvlNetworkResponseBody) *DisableKibanaPvlNetworkResponse {
	s.Body = v
	return s
}

func (s *DisableKibanaPvlNetworkResponse) Validate() error {
	return dara.Validate(s)
}

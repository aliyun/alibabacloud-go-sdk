// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisablePolicyTypeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DisablePolicyTypeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DisablePolicyTypeResponse
	GetStatusCode() *int32
	SetBody(v *DisablePolicyTypeResponseBody) *DisablePolicyTypeResponse
	GetBody() *DisablePolicyTypeResponseBody
}

type DisablePolicyTypeResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DisablePolicyTypeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DisablePolicyTypeResponse) String() string {
	return dara.Prettify(s)
}

func (s DisablePolicyTypeResponse) GoString() string {
	return s.String()
}

func (s *DisablePolicyTypeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DisablePolicyTypeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DisablePolicyTypeResponse) GetBody() *DisablePolicyTypeResponseBody {
	return s.Body
}

func (s *DisablePolicyTypeResponse) SetHeaders(v map[string]*string) *DisablePolicyTypeResponse {
	s.Headers = v
	return s
}

func (s *DisablePolicyTypeResponse) SetStatusCode(v int32) *DisablePolicyTypeResponse {
	s.StatusCode = &v
	return s
}

func (s *DisablePolicyTypeResponse) SetBody(v *DisablePolicyTypeResponseBody) *DisablePolicyTypeResponse {
	s.Body = v
	return s
}

func (s *DisablePolicyTypeResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateEipAddressProResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AllocateEipAddressProResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AllocateEipAddressProResponse
	GetStatusCode() *int32
	SetBody(v *AllocateEipAddressProResponseBody) *AllocateEipAddressProResponse
	GetBody() *AllocateEipAddressProResponseBody
}

type AllocateEipAddressProResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AllocateEipAddressProResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AllocateEipAddressProResponse) String() string {
	return dara.Prettify(s)
}

func (s AllocateEipAddressProResponse) GoString() string {
	return s.String()
}

func (s *AllocateEipAddressProResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AllocateEipAddressProResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AllocateEipAddressProResponse) GetBody() *AllocateEipAddressProResponseBody {
	return s.Body
}

func (s *AllocateEipAddressProResponse) SetHeaders(v map[string]*string) *AllocateEipAddressProResponse {
	s.Headers = v
	return s
}

func (s *AllocateEipAddressProResponse) SetStatusCode(v int32) *AllocateEipAddressProResponse {
	s.StatusCode = &v
	return s
}

func (s *AllocateEipAddressProResponse) SetBody(v *AllocateEipAddressProResponseBody) *AllocateEipAddressProResponse {
	s.Body = v
	return s
}

func (s *AllocateEipAddressProResponse) Validate() error {
	return dara.Validate(s)
}

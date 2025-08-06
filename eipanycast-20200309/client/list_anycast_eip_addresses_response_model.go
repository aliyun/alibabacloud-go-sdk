// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAnycastEipAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListAnycastEipAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListAnycastEipAddressesResponse
	GetStatusCode() *int32
	SetBody(v *ListAnycastEipAddressesResponseBody) *ListAnycastEipAddressesResponse
	GetBody() *ListAnycastEipAddressesResponseBody
}

type ListAnycastEipAddressesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAnycastEipAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAnycastEipAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListAnycastEipAddressesResponse) GoString() string {
	return s.String()
}

func (s *ListAnycastEipAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListAnycastEipAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListAnycastEipAddressesResponse) GetBody() *ListAnycastEipAddressesResponseBody {
	return s.Body
}

func (s *ListAnycastEipAddressesResponse) SetHeaders(v map[string]*string) *ListAnycastEipAddressesResponse {
	s.Headers = v
	return s
}

func (s *ListAnycastEipAddressesResponse) SetStatusCode(v int32) *ListAnycastEipAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAnycastEipAddressesResponse) SetBody(v *ListAnycastEipAddressesResponseBody) *ListAnycastEipAddressesResponse {
	s.Body = v
	return s
}

func (s *ListAnycastEipAddressesResponse) Validate() error {
	return dara.Validate(s)
}

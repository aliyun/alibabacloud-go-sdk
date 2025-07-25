// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListCloudGtmAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListCloudGtmAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListCloudGtmAddressesResponse
	GetStatusCode() *int32
	SetBody(v *ListCloudGtmAddressesResponseBody) *ListCloudGtmAddressesResponse
	GetBody() *ListCloudGtmAddressesResponseBody
}

type ListCloudGtmAddressesResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListCloudGtmAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListCloudGtmAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListCloudGtmAddressesResponse) GoString() string {
	return s.String()
}

func (s *ListCloudGtmAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListCloudGtmAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListCloudGtmAddressesResponse) GetBody() *ListCloudGtmAddressesResponseBody {
	return s.Body
}

func (s *ListCloudGtmAddressesResponse) SetHeaders(v map[string]*string) *ListCloudGtmAddressesResponse {
	s.Headers = v
	return s
}

func (s *ListCloudGtmAddressesResponse) SetStatusCode(v int32) *ListCloudGtmAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListCloudGtmAddressesResponse) SetBody(v *ListCloudGtmAddressesResponseBody) *ListCloudGtmAddressesResponse {
	s.Body = v
	return s
}

func (s *ListCloudGtmAddressesResponse) Validate() error {
	return dara.Validate(s)
}

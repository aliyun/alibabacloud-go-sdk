// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSearchCloudGtmAddressesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SearchCloudGtmAddressesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SearchCloudGtmAddressesResponse
	GetStatusCode() *int32
	SetBody(v *SearchCloudGtmAddressesResponseBody) *SearchCloudGtmAddressesResponse
	GetBody() *SearchCloudGtmAddressesResponseBody
}

type SearchCloudGtmAddressesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchCloudGtmAddressesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchCloudGtmAddressesResponse) String() string {
	return dara.Prettify(s)
}

func (s SearchCloudGtmAddressesResponse) GoString() string {
	return s.String()
}

func (s *SearchCloudGtmAddressesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SearchCloudGtmAddressesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SearchCloudGtmAddressesResponse) GetBody() *SearchCloudGtmAddressesResponseBody {
	return s.Body
}

func (s *SearchCloudGtmAddressesResponse) SetHeaders(v map[string]*string) *SearchCloudGtmAddressesResponse {
	s.Headers = v
	return s
}

func (s *SearchCloudGtmAddressesResponse) SetStatusCode(v int32) *SearchCloudGtmAddressesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchCloudGtmAddressesResponse) SetBody(v *SearchCloudGtmAddressesResponseBody) *SearchCloudGtmAddressesResponse {
	s.Body = v
	return s
}

func (s *SearchCloudGtmAddressesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

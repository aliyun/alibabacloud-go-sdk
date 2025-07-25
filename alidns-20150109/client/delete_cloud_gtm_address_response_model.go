// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCloudGtmAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteCloudGtmAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteCloudGtmAddressResponse
	GetStatusCode() *int32
	SetBody(v *DeleteCloudGtmAddressResponseBody) *DeleteCloudGtmAddressResponse
	GetBody() *DeleteCloudGtmAddressResponseBody
}

type DeleteCloudGtmAddressResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteCloudGtmAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteCloudGtmAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteCloudGtmAddressResponse) GoString() string {
	return s.String()
}

func (s *DeleteCloudGtmAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteCloudGtmAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteCloudGtmAddressResponse) GetBody() *DeleteCloudGtmAddressResponseBody {
	return s.Body
}

func (s *DeleteCloudGtmAddressResponse) SetHeaders(v map[string]*string) *DeleteCloudGtmAddressResponse {
	s.Headers = v
	return s
}

func (s *DeleteCloudGtmAddressResponse) SetStatusCode(v int32) *DeleteCloudGtmAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteCloudGtmAddressResponse) SetBody(v *DeleteCloudGtmAddressResponseBody) *DeleteCloudGtmAddressResponse {
	s.Body = v
	return s
}

func (s *DeleteCloudGtmAddressResponse) Validate() error {
	return dara.Validate(s)
}

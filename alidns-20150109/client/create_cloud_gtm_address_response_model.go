// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudGtmAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudGtmAddressResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudGtmAddressResponseBody) *CreateCloudGtmAddressResponse
	GetBody() *CreateCloudGtmAddressResponseBody
}

type CreateCloudGtmAddressResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudGtmAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudGtmAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudGtmAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudGtmAddressResponse) GetBody() *CreateCloudGtmAddressResponseBody {
	return s.Body
}

func (s *CreateCloudGtmAddressResponse) SetHeaders(v map[string]*string) *CreateCloudGtmAddressResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudGtmAddressResponse) SetStatusCode(v int32) *CreateCloudGtmAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudGtmAddressResponse) SetBody(v *CreateCloudGtmAddressResponseBody) *CreateCloudGtmAddressResponse {
	s.Body = v
	return s
}

func (s *CreateCloudGtmAddressResponse) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressResponseBody) *UpdateCloudGtmAddressResponse
	GetBody() *UpdateCloudGtmAddressResponseBody
}

type UpdateCloudGtmAddressResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressResponse) GetBody() *UpdateCloudGtmAddressResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressResponse) SetBody(v *UpdateCloudGtmAddressResponseBody) *UpdateCloudGtmAddressResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

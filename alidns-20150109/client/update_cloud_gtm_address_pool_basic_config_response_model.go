// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolBasicConfigResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolBasicConfigResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressPoolBasicConfigResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressPoolBasicConfigResponseBody) *UpdateCloudGtmAddressPoolBasicConfigResponse
	GetBody() *UpdateCloudGtmAddressPoolBasicConfigResponseBody
}

type UpdateCloudGtmAddressPoolBasicConfigResponse struct {
	Headers    map[string]*string                                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressPoolBasicConfigResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressPoolBasicConfigResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolBasicConfigResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) GetBody() *UpdateCloudGtmAddressPoolBasicConfigResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolBasicConfigResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressPoolBasicConfigResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) SetBody(v *UpdateCloudGtmAddressPoolBasicConfigResponseBody) *UpdateCloudGtmAddressPoolBasicConfigResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressPoolBasicConfigResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

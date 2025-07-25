// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressPoolEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressPoolEnableStatusResponseBody) *UpdateCloudGtmAddressPoolEnableStatusResponse
	GetBody() *UpdateCloudGtmAddressPoolEnableStatusResponseBody
}

type UpdateCloudGtmAddressPoolEnableStatusResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressPoolEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressPoolEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) GetBody() *UpdateCloudGtmAddressPoolEnableStatusResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressPoolEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) SetBody(v *UpdateCloudGtmAddressPoolEnableStatusResponseBody) *UpdateCloudGtmAddressPoolEnableStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressPoolEnableStatusResponse) Validate() error {
	return dara.Validate(s)
}

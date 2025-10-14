// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressEnableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressEnableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressEnableStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressEnableStatusResponseBody) *UpdateCloudGtmAddressEnableStatusResponse
	GetBody() *UpdateCloudGtmAddressEnableStatusResponseBody
}

type UpdateCloudGtmAddressEnableStatusResponse struct {
	Headers    map[string]*string                             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressEnableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressEnableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressEnableStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) GetBody() *UpdateCloudGtmAddressEnableStatusResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressEnableStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressEnableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) SetBody(v *UpdateCloudGtmAddressEnableStatusResponseBody) *UpdateCloudGtmAddressEnableStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressEnableStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

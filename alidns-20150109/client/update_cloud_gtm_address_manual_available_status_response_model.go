// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressManualAvailableStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressManualAvailableStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressManualAvailableStatusResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressManualAvailableStatusResponseBody) *UpdateCloudGtmAddressManualAvailableStatusResponse
	GetBody() *UpdateCloudGtmAddressManualAvailableStatusResponseBody
}

type UpdateCloudGtmAddressManualAvailableStatusResponse struct {
	Headers    map[string]*string                                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressManualAvailableStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressManualAvailableStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressManualAvailableStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) GetBody() *UpdateCloudGtmAddressManualAvailableStatusResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressManualAvailableStatusResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressManualAvailableStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) SetBody(v *UpdateCloudGtmAddressManualAvailableStatusResponseBody) *UpdateCloudGtmAddressManualAvailableStatusResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressManualAvailableStatusResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

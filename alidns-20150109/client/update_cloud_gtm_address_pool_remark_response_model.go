// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressPoolRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressPoolRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressPoolRemarkResponseBody) *UpdateCloudGtmAddressPoolRemarkResponse
	GetBody() *UpdateCloudGtmAddressPoolRemarkResponseBody
}

type UpdateCloudGtmAddressPoolRemarkResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressPoolRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressPoolRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressPoolRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) GetBody() *UpdateCloudGtmAddressPoolRemarkResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressPoolRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressPoolRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) SetBody(v *UpdateCloudGtmAddressPoolRemarkResponseBody) *UpdateCloudGtmAddressPoolRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressPoolRemarkResponse) Validate() error {
	return dara.Validate(s)
}

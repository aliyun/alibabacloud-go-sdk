// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCloudGtmAddressRemarkResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateCloudGtmAddressRemarkResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateCloudGtmAddressRemarkResponse
	GetStatusCode() *int32
	SetBody(v *UpdateCloudGtmAddressRemarkResponseBody) *UpdateCloudGtmAddressRemarkResponse
	GetBody() *UpdateCloudGtmAddressRemarkResponseBody
}

type UpdateCloudGtmAddressRemarkResponse struct {
	Headers    map[string]*string                       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateCloudGtmAddressRemarkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateCloudGtmAddressRemarkResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateCloudGtmAddressRemarkResponse) GoString() string {
	return s.String()
}

func (s *UpdateCloudGtmAddressRemarkResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateCloudGtmAddressRemarkResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateCloudGtmAddressRemarkResponse) GetBody() *UpdateCloudGtmAddressRemarkResponseBody {
	return s.Body
}

func (s *UpdateCloudGtmAddressRemarkResponse) SetHeaders(v map[string]*string) *UpdateCloudGtmAddressRemarkResponse {
	s.Headers = v
	return s
}

func (s *UpdateCloudGtmAddressRemarkResponse) SetStatusCode(v int32) *UpdateCloudGtmAddressRemarkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateCloudGtmAddressRemarkResponse) SetBody(v *UpdateCloudGtmAddressRemarkResponseBody) *UpdateCloudGtmAddressRemarkResponse {
	s.Body = v
	return s
}

func (s *UpdateCloudGtmAddressRemarkResponse) Validate() error {
	return dara.Validate(s)
}

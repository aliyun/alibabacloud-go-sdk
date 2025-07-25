// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateCloudGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateCloudGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *CreateCloudGtmAddressPoolResponseBody) *CreateCloudGtmAddressPoolResponse
	GetBody() *CreateCloudGtmAddressPoolResponseBody
}

type CreateCloudGtmAddressPoolResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCloudGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCloudGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *CreateCloudGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateCloudGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateCloudGtmAddressPoolResponse) GetBody() *CreateCloudGtmAddressPoolResponseBody {
	return s.Body
}

func (s *CreateCloudGtmAddressPoolResponse) SetHeaders(v map[string]*string) *CreateCloudGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *CreateCloudGtmAddressPoolResponse) SetStatusCode(v int32) *CreateCloudGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCloudGtmAddressPoolResponse) SetBody(v *CreateCloudGtmAddressPoolResponseBody) *CreateCloudGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *CreateCloudGtmAddressPoolResponse) Validate() error {
	return dara.Validate(s)
}

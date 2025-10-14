// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *UpdateGtmAddressPoolResponseBody) *UpdateGtmAddressPoolResponse
	GetBody() *UpdateGtmAddressPoolResponseBody
}

type UpdateGtmAddressPoolResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *UpdateGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateGtmAddressPoolResponse) GetBody() *UpdateGtmAddressPoolResponseBody {
	return s.Body
}

func (s *UpdateGtmAddressPoolResponse) SetHeaders(v map[string]*string) *UpdateGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *UpdateGtmAddressPoolResponse) SetStatusCode(v int32) *UpdateGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGtmAddressPoolResponse) SetBody(v *UpdateGtmAddressPoolResponseBody) *UpdateGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *UpdateGtmAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

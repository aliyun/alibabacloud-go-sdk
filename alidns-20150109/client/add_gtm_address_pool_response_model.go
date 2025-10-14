// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddGtmAddressPoolResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddGtmAddressPoolResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddGtmAddressPoolResponse
	GetStatusCode() *int32
	SetBody(v *AddGtmAddressPoolResponseBody) *AddGtmAddressPoolResponse
	GetBody() *AddGtmAddressPoolResponseBody
}

type AddGtmAddressPoolResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddGtmAddressPoolResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddGtmAddressPoolResponse) String() string {
	return dara.Prettify(s)
}

func (s AddGtmAddressPoolResponse) GoString() string {
	return s.String()
}

func (s *AddGtmAddressPoolResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddGtmAddressPoolResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddGtmAddressPoolResponse) GetBody() *AddGtmAddressPoolResponseBody {
	return s.Body
}

func (s *AddGtmAddressPoolResponse) SetHeaders(v map[string]*string) *AddGtmAddressPoolResponse {
	s.Headers = v
	return s
}

func (s *AddGtmAddressPoolResponse) SetStatusCode(v int32) *AddGtmAddressPoolResponse {
	s.StatusCode = &v
	return s
}

func (s *AddGtmAddressPoolResponse) SetBody(v *AddGtmAddressPoolResponseBody) *AddGtmAddressPoolResponse {
	s.Body = v
	return s
}

func (s *AddGtmAddressPoolResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

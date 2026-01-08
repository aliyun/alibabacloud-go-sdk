// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressRecoverSuspendResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAddressRecoverSuspendResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAddressRecoverSuspendResponse
	GetStatusCode() *int32
	SetBody(v *AddAddressRecoverSuspendResponseBody) *AddAddressRecoverSuspendResponse
	GetBody() *AddAddressRecoverSuspendResponseBody
}

type AddAddressRecoverSuspendResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAddressRecoverSuspendResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAddressRecoverSuspendResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAddressRecoverSuspendResponse) GoString() string {
	return s.String()
}

func (s *AddAddressRecoverSuspendResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAddressRecoverSuspendResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAddressRecoverSuspendResponse) GetBody() *AddAddressRecoverSuspendResponseBody {
	return s.Body
}

func (s *AddAddressRecoverSuspendResponse) SetHeaders(v map[string]*string) *AddAddressRecoverSuspendResponse {
	s.Headers = v
	return s
}

func (s *AddAddressRecoverSuspendResponse) SetStatusCode(v int32) *AddAddressRecoverSuspendResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAddressRecoverSuspendResponse) SetBody(v *AddAddressRecoverSuspendResponseBody) *AddAddressRecoverSuspendResponse {
	s.Body = v
	return s
}

func (s *AddAddressRecoverSuspendResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

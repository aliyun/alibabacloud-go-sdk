// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressBookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddAddressBookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddAddressBookResponse
	GetStatusCode() *int32
	SetBody(v *AddAddressBookResponseBody) *AddAddressBookResponse
	GetBody() *AddAddressBookResponseBody
}

type AddAddressBookResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddAddressBookResponse) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookResponse) GoString() string {
	return s.String()
}

func (s *AddAddressBookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddAddressBookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddAddressBookResponse) GetBody() *AddAddressBookResponseBody {
	return s.Body
}

func (s *AddAddressBookResponse) SetHeaders(v map[string]*string) *AddAddressBookResponse {
	s.Headers = v
	return s
}

func (s *AddAddressBookResponse) SetStatusCode(v int32) *AddAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *AddAddressBookResponse) SetBody(v *AddAddressBookResponseBody) *AddAddressBookResponse {
	s.Body = v
	return s
}

func (s *AddAddressBookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

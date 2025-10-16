// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddressBookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteAddressBookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteAddressBookResponse
	GetStatusCode() *int32
	SetBody(v *DeleteAddressBookResponseBody) *DeleteAddressBookResponse
	GetBody() *DeleteAddressBookResponseBody
}

type DeleteAddressBookResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteAddressBookResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddressBookResponse) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteAddressBookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteAddressBookResponse) GetBody() *DeleteAddressBookResponseBody {
	return s.Body
}

func (s *DeleteAddressBookResponse) SetHeaders(v map[string]*string) *DeleteAddressBookResponse {
	s.Headers = v
	return s
}

func (s *DeleteAddressBookResponse) SetStatusCode(v int32) *DeleteAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteAddressBookResponse) SetBody(v *DeleteAddressBookResponseBody) *DeleteAddressBookResponse {
	s.Body = v
	return s
}

func (s *DeleteAddressBookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

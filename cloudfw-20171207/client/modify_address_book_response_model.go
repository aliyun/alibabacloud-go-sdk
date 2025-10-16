// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAddressBookResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyAddressBookResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyAddressBookResponse
	GetStatusCode() *int32
	SetBody(v *ModifyAddressBookResponseBody) *ModifyAddressBookResponse
	GetBody() *ModifyAddressBookResponseBody
}

type ModifyAddressBookResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyAddressBookResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyAddressBookResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookResponse) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyAddressBookResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyAddressBookResponse) GetBody() *ModifyAddressBookResponseBody {
	return s.Body
}

func (s *ModifyAddressBookResponse) SetHeaders(v map[string]*string) *ModifyAddressBookResponse {
	s.Headers = v
	return s
}

func (s *ModifyAddressBookResponse) SetStatusCode(v int32) *ModifyAddressBookResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyAddressBookResponse) SetBody(v *ModifyAddressBookResponseBody) *ModifyAddressBookResponse {
	s.Body = v
	return s
}

func (s *ModifyAddressBookResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}

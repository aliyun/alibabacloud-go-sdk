// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAddressBookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupUuid(v string) *AddAddressBookResponseBody
	GetGroupUuid() *string
	SetRequestId(v string) *AddAddressBookResponseBody
	GetRequestId() *string
}

type AddAddressBookResponseBody struct {
	// The UUID of the returned address book.
	//
	// example:
	//
	// f04ac7ce-628b-4cb7-be61-310222b7****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The request ID.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAddressBookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *AddAddressBookResponseBody) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *AddAddressBookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAddressBookResponseBody) SetGroupUuid(v string) *AddAddressBookResponseBody {
	s.GroupUuid = &v
	return s
}

func (s *AddAddressBookResponseBody) SetRequestId(v string) *AddAddressBookResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAddressBookResponseBody) Validate() error {
	return dara.Validate(s)
}

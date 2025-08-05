// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAddressBookRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGroupUuid(v string) *DeleteAddressBookRequest
	GetGroupUuid() *string
	SetLang(v string) *DeleteAddressBookRequest
	GetLang() *string
	SetSourceIp(v string) *DeleteAddressBookRequest
	GetSourceIp() *string
}

type DeleteAddressBookRequest struct {
	// The ID of the address book.
	//
	// To delete the address book, you must provide the ID of the address book. You can call the DescribeAddressBook operation to query the ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0657ab9d-fe8b-4174-b2a6-6baf358e****
	GroupUuid *string `json:"GroupUuid,omitempty" xml:"GroupUuid,omitempty"`
	// The natural language of the request and response. Valid values:
	//
	// 	- **zh**: Chinese (default)
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// Deprecated
	//
	// The source IP address of the request.
	//
	// example:
	//
	// 192.0.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s DeleteAddressBookRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAddressBookRequest) GoString() string {
	return s.String()
}

func (s *DeleteAddressBookRequest) GetGroupUuid() *string {
	return s.GroupUuid
}

func (s *DeleteAddressBookRequest) GetLang() *string {
	return s.Lang
}

func (s *DeleteAddressBookRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DeleteAddressBookRequest) SetGroupUuid(v string) *DeleteAddressBookRequest {
	s.GroupUuid = &v
	return s
}

func (s *DeleteAddressBookRequest) SetLang(v string) *DeleteAddressBookRequest {
	s.Lang = &v
	return s
}

func (s *DeleteAddressBookRequest) SetSourceIp(v string) *DeleteAddressBookRequest {
	s.SourceIp = &v
	return s
}

func (s *DeleteAddressBookRequest) Validate() error {
	return dara.Validate(s)
}

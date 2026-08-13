// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAddressBookResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *ModifyAddressBookResponseBody
	GetDryRun() *bool
	SetRequestId(v string) *ModifyAddressBookResponseBody
	GetRequestId() *string
}

type ModifyAddressBookResponseBody struct {
	// Indicates that this is a successful dry run response. A value of true indicates that only the dry run was completed and no actual modification was performed.
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// CBF1E9B7-D6A0-4E9E-AD3E-2B47E6C2837D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAddressBookResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAddressBookResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAddressBookResponseBody) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAddressBookResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAddressBookResponseBody) SetDryRun(v bool) *ModifyAddressBookResponseBody {
	s.DryRun = &v
	return s
}

func (s *ModifyAddressBookResponseBody) SetRequestId(v string) *ModifyAddressBookResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAddressBookResponseBody) Validate() error {
	return dara.Validate(s)
}

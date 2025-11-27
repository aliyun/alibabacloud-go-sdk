// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckRdsCustomInitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetHasServiceLinkedRole(v string) *CheckRdsCustomInitResponseBody
	GetHasServiceLinkedRole() *string
	SetRegisterUidSuccess(v bool) *CheckRdsCustomInitResponseBody
	GetRegisterUidSuccess() *bool
	SetRequestId(v string) *CheckRdsCustomInitResponseBody
	GetRequestId() *string
	SetRequireServiceLinkedRole(v string) *CheckRdsCustomInitResponseBody
	GetRequireServiceLinkedRole() *string
}

type CheckRdsCustomInitResponseBody struct {
	// example:
	//
	// true
	HasServiceLinkedRole *string `json:"HasServiceLinkedRole,omitempty" xml:"HasServiceLinkedRole,omitempty"`
	// example:
	//
	// true
	RegisterUidSuccess *bool `json:"RegisterUidSuccess,omitempty" xml:"RegisterUidSuccess,omitempty"`
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	RequireServiceLinkedRole *string `json:"RequireServiceLinkedRole,omitempty" xml:"RequireServiceLinkedRole,omitempty"`
}

func (s CheckRdsCustomInitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckRdsCustomInitResponseBody) GoString() string {
	return s.String()
}

func (s *CheckRdsCustomInitResponseBody) GetHasServiceLinkedRole() *string {
	return s.HasServiceLinkedRole
}

func (s *CheckRdsCustomInitResponseBody) GetRegisterUidSuccess() *bool {
	return s.RegisterUidSuccess
}

func (s *CheckRdsCustomInitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckRdsCustomInitResponseBody) GetRequireServiceLinkedRole() *string {
	return s.RequireServiceLinkedRole
}

func (s *CheckRdsCustomInitResponseBody) SetHasServiceLinkedRole(v string) *CheckRdsCustomInitResponseBody {
	s.HasServiceLinkedRole = &v
	return s
}

func (s *CheckRdsCustomInitResponseBody) SetRegisterUidSuccess(v bool) *CheckRdsCustomInitResponseBody {
	s.RegisterUidSuccess = &v
	return s
}

func (s *CheckRdsCustomInitResponseBody) SetRequestId(v string) *CheckRdsCustomInitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckRdsCustomInitResponseBody) SetRequireServiceLinkedRole(v string) *CheckRdsCustomInitResponseBody {
	s.RequireServiceLinkedRole = &v
	return s
}

func (s *CheckRdsCustomInitResponseBody) Validate() error {
	return dara.Validate(s)
}

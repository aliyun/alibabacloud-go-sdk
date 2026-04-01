// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountMaskingPrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAccountMaskingPrivilegeResponseBodyData) *DescribeAccountMaskingPrivilegeResponseBody
	GetData() *DescribeAccountMaskingPrivilegeResponseBodyData
	SetRequestId(v string) *DescribeAccountMaskingPrivilegeResponseBody
	GetRequestId() *string
}

type DescribeAccountMaskingPrivilegeResponseBody struct {
	Data *DescribeAccountMaskingPrivilegeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// D0073A98-52F1-3075-8256-394**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountMaskingPrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountMaskingPrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountMaskingPrivilegeResponseBody) GetData() *DescribeAccountMaskingPrivilegeResponseBodyData {
	return s.Data
}

func (s *DescribeAccountMaskingPrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountMaskingPrivilegeResponseBody) SetData(v *DescribeAccountMaskingPrivilegeResponseBodyData) *DescribeAccountMaskingPrivilegeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponseBody) SetRequestId(v string) *DescribeAccountMaskingPrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountMaskingPrivilegeResponseBodyData struct {
	UserPrivilege []*DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege `json:"UserPrivilege,omitempty" xml:"UserPrivilege,omitempty" type:"Repeated"`
}

func (s DescribeAccountMaskingPrivilegeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountMaskingPrivilegeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyData) GetUserPrivilege() []*DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege {
	return s.UserPrivilege
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyData) SetUserPrivilege(v []*DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) *DescribeAccountMaskingPrivilegeResponseBodyData {
	s.UserPrivilege = v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyData) Validate() error {
	if s.UserPrivilege != nil {
		for _, item := range s.UserPrivilege {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege struct {
	// example:
	//
	// 2026-01-22T02:01:20Z
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// example:
	//
	// restrictedAccess
	Privilege *string `json:"Privilege,omitempty" xml:"Privilege,omitempty"`
	// example:
	//
	// rds
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) GoString() string {
	return s.String()
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) GetPrivilege() *string {
	return s.Privilege
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) GetUserName() *string {
	return s.UserName
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) SetExpireTime(v string) *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege {
	s.ExpireTime = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) SetPrivilege(v string) *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege {
	s.Privilege = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) SetUserName(v string) *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege {
	s.UserName = &v
	return s
}

func (s *DescribeAccountMaskingPrivilegeResponseBodyDataUserPrivilege) Validate() error {
	return dara.Validate(s)
}

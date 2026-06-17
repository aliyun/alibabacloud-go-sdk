// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEncryptionDBRolePrivilegeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeEncryptionDBRolePrivilegeResponseBody
	GetDBClusterId() *string
	SetData(v *DescribeEncryptionDBRolePrivilegeResponseBodyData) *DescribeEncryptionDBRolePrivilegeResponseBody
	GetData() *DescribeEncryptionDBRolePrivilegeResponseBodyData
	SetMessage(v string) *DescribeEncryptionDBRolePrivilegeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeEncryptionDBRolePrivilegeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeEncryptionDBRolePrivilegeResponseBody
	GetSuccess() *bool
}

type DescribeEncryptionDBRolePrivilegeResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// pc-****************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The result set.
	Data *DescribeEncryptionDBRolePrivilegeResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response message.
	//
	// > If the request is successful, \\`Successful\\` is returned. If the request fails, an error message, such as an error code, is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// D0CEC6AC-7760-409A-A0D5-E6CD86******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeEncryptionDBRolePrivilegeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBRolePrivilegeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) GetData() *DescribeEncryptionDBRolePrivilegeResponseBodyData {
	return s.Data
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) SetDBClusterId(v string) *DescribeEncryptionDBRolePrivilegeResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) SetData(v *DescribeEncryptionDBRolePrivilegeResponseBodyData) *DescribeEncryptionDBRolePrivilegeResponseBody {
	s.Data = v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) SetMessage(v string) *DescribeEncryptionDBRolePrivilegeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) SetRequestId(v string) *DescribeEncryptionDBRolePrivilegeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) SetSuccess(v bool) *DescribeEncryptionDBRolePrivilegeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeEncryptionDBRolePrivilegeResponseBodyData struct {
	// A list of role access policies.
	RolePrivilegeList []*DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList `json:"RolePrivilegeList,omitempty" xml:"RolePrivilegeList,omitempty" type:"Repeated"`
}

func (s DescribeEncryptionDBRolePrivilegeResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBRolePrivilegeResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyData) GetRolePrivilegeList() []*DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList {
	return s.RolePrivilegeList
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyData) SetRolePrivilegeList(v []*DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) *DescribeEncryptionDBRolePrivilegeResponseBodyData {
	s.RolePrivilegeList = v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyData) Validate() error {
	if s.RolePrivilegeList != nil {
		for _, item := range s.RolePrivilegeList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList struct {
	// The regular users.
	//
	// example:
	//
	// [alton01]
	Encryption *string `json:"Encryption,omitempty" xml:"Encryption,omitempty"`
	// Other users.
	//
	// example:
	//
	// test
	Negation *string `json:"Negation,omitempty" xml:"Negation,omitempty"`
	// The privileged users.
	//
	// example:
	//
	// [alton]
	NotEncryption *string `json:"NotEncryption,omitempty" xml:"NotEncryption,omitempty"`
	// The name of the role permission.
	//
	// example:
	//
	// test
	RolePrivilegeName *string `json:"RolePrivilegeName,omitempty" xml:"RolePrivilegeName,omitempty"`
}

func (s DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) String() string {
	return dara.Prettify(s)
}

func (s DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) GoString() string {
	return s.String()
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) GetEncryption() *string {
	return s.Encryption
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) GetNegation() *string {
	return s.Negation
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) GetNotEncryption() *string {
	return s.NotEncryption
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) GetRolePrivilegeName() *string {
	return s.RolePrivilegeName
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) SetEncryption(v string) *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList {
	s.Encryption = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) SetNegation(v string) *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList {
	s.Negation = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) SetNotEncryption(v string) *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList {
	s.NotEncryption = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) SetRolePrivilegeName(v string) *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList {
	s.RolePrivilegeName = &v
	return s
}

func (s *DescribeEncryptionDBRolePrivilegeResponseBodyDataRolePrivilegeList) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetResourceDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetResourceDirectoryResponseBody
	GetRequestId() *string
	SetResourceDirectory(v *GetResourceDirectoryResponseBodyResourceDirectory) *GetResourceDirectoryResponseBody
	GetResourceDirectory() *GetResourceDirectoryResponseBodyResourceDirectory
}

type GetResourceDirectoryResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// CD76D376-2517-4924-92C5-DBC52262F93A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the resource directory.
	ResourceDirectory *GetResourceDirectoryResponseBodyResourceDirectory `json:"ResourceDirectory,omitempty" xml:"ResourceDirectory,omitempty" type:"Struct"`
}

func (s GetResourceDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetResourceDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetResourceDirectoryResponseBody) GetResourceDirectory() *GetResourceDirectoryResponseBodyResourceDirectory {
	return s.ResourceDirectory
}

func (s *GetResourceDirectoryResponseBody) SetRequestId(v string) *GetResourceDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetResourceDirectoryResponseBody) SetResourceDirectory(v *GetResourceDirectoryResponseBodyResourceDirectory) *GetResourceDirectoryResponseBody {
	s.ResourceDirectory = v
	return s
}

func (s *GetResourceDirectoryResponseBody) Validate() error {
	if s.ResourceDirectory != nil {
		if err := s.ResourceDirectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetResourceDirectoryResponseBodyResourceDirectory struct {
	// The status of the Control Policy feature. Valid values:
	//
	// - Enabled: The feature is enabled.
	//
	// - PendingEnable: The feature is being enabled.
	//
	// - Disabled: The feature is disabled.
	//
	// - PendingDisable: The feature is being disabled.
	//
	// example:
	//
	// Enabled
	ControlPolicyStatus *string `json:"ControlPolicyStatus,omitempty" xml:"ControlPolicyStatus,omitempty"`
	// The time when the resource directory was enabled.
	//
	// example:
	//
	// 2019-02-18T15:32:10.473Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The real-name verification information.
	//
	// example:
	//
	// **	- Co., Ltd.
	IdentityInformation *string `json:"IdentityInformation,omitempty" xml:"IdentityInformation,omitempty"`
	// The ID of the management account.
	//
	// example:
	//
	// 172845045600****
	MasterAccountId *string `json:"MasterAccountId,omitempty" xml:"MasterAccountId,omitempty"`
	// The name of the management account.
	//
	// example:
	//
	// aliyun-admin
	MasterAccountName *string `json:"MasterAccountName,omitempty" xml:"MasterAccountName,omitempty"`
	// The status of the Member Display Name Synchronization feature. Valid values:
	//
	// 	- Enabled
	//
	// 	- Disabled
	//
	// example:
	//
	// Enabled
	MemberAccountDisplayNameSyncStatus *string `json:"MemberAccountDisplayNameSyncStatus,omitempty" xml:"MemberAccountDisplayNameSyncStatus,omitempty"`
	// The status of the member deletion feature. Valid values:
	//
	// - Enabled: The feature is enabled. You can call the [DeleteAccount](~~DeleteAccount~~) operation to delete members of the resource account type.
	//
	// - Disabled: The feature is disabled. You cannot delete members of the resource account type.
	//
	// example:
	//
	// Enabled
	MemberDeletionStatus *string `json:"MemberDeletionStatus,omitempty" xml:"MemberDeletionStatus,omitempty"`
	// The ID of the resource directory.
	//
	// example:
	//
	// rd-St****
	ResourceDirectoryId *string `json:"ResourceDirectoryId,omitempty" xml:"ResourceDirectoryId,omitempty"`
	// The ID of the Root folder.
	//
	// example:
	//
	// r-Zo****
	RootFolderId *string `json:"RootFolderId,omitempty" xml:"RootFolderId,omitempty"`
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) String() string {
	return dara.Prettify(s)
}

func (s GetResourceDirectoryResponseBodyResourceDirectory) GoString() string {
	return s.String()
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetControlPolicyStatus() *string {
	return s.ControlPolicyStatus
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetIdentityInformation() *string {
	return s.IdentityInformation
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetMasterAccountId() *string {
	return s.MasterAccountId
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetMasterAccountName() *string {
	return s.MasterAccountName
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetMemberAccountDisplayNameSyncStatus() *string {
	return s.MemberAccountDisplayNameSyncStatus
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetMemberDeletionStatus() *string {
	return s.MemberDeletionStatus
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetResourceDirectoryId() *string {
	return s.ResourceDirectoryId
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) GetRootFolderId() *string {
	return s.RootFolderId
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetControlPolicyStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.ControlPolicyStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetCreateTime(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.CreateTime = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetIdentityInformation(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.IdentityInformation = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMasterAccountName(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MasterAccountName = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMemberAccountDisplayNameSyncStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MemberAccountDisplayNameSyncStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetMemberDeletionStatus(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.MemberDeletionStatus = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetResourceDirectoryId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.ResourceDirectoryId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) SetRootFolderId(v string) *GetResourceDirectoryResponseBodyResourceDirectory {
	s.RootFolderId = &v
	return s
}

func (s *GetResourceDirectoryResponseBodyResourceDirectory) Validate() error {
	return dara.Validate(s)
}

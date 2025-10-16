// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUsersInGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEndUsers(v []*DescribeUsersInGroupResponseBodyEndUsers) *DescribeUsersInGroupResponseBody
	GetEndUsers() []*DescribeUsersInGroupResponseBodyEndUsers
	SetNextToken(v string) *DescribeUsersInGroupResponseBody
	GetNextToken() *string
	SetOnlineUsersCount(v int32) *DescribeUsersInGroupResponseBody
	GetOnlineUsersCount() *int32
	SetRequestId(v string) *DescribeUsersInGroupResponseBody
	GetRequestId() *string
	SetUserGroupName(v string) *DescribeUsersInGroupResponseBody
	GetUserGroupName() *string
	SetUserOuPath(v string) *DescribeUsersInGroupResponseBody
	GetUserOuPath() *string
	SetUsersCount(v int32) *DescribeUsersInGroupResponseBody
	GetUsersCount() *int32
}

type DescribeUsersInGroupResponseBody struct {
	// The authorized users.
	EndUsers []*DescribeUsersInGroupResponseBodyEndUsers `json:"EndUsers,omitempty" xml:"EndUsers,omitempty" type:"Repeated"`
	// The token that is used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The total number of authorized users that are connected to cloud computers of the cloud computer share.
	//
	// example:
	//
	// 0
	OnlineUsersCount *int32 `json:"OnlineUsersCount,omitempty" xml:"OnlineUsersCount,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UserGroupName *string `json:"UserGroupName,omitempty" xml:"UserGroupName,omitempty"`
	UserOuPath    *string `json:"UserOuPath,omitempty" xml:"UserOuPath,omitempty"`
	// The total number of authorized users of the cloud computer share.
	//
	// example:
	//
	// 1
	UsersCount *int32 `json:"UsersCount,omitempty" xml:"UsersCount,omitempty"`
}

func (s DescribeUsersInGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBody) GetEndUsers() []*DescribeUsersInGroupResponseBodyEndUsers {
	return s.EndUsers
}

func (s *DescribeUsersInGroupResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeUsersInGroupResponseBody) GetOnlineUsersCount() *int32 {
	return s.OnlineUsersCount
}

func (s *DescribeUsersInGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUsersInGroupResponseBody) GetUserGroupName() *string {
	return s.UserGroupName
}

func (s *DescribeUsersInGroupResponseBody) GetUserOuPath() *string {
	return s.UserOuPath
}

func (s *DescribeUsersInGroupResponseBody) GetUsersCount() *int32 {
	return s.UsersCount
}

func (s *DescribeUsersInGroupResponseBody) SetEndUsers(v []*DescribeUsersInGroupResponseBodyEndUsers) *DescribeUsersInGroupResponseBody {
	s.EndUsers = v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetNextToken(v string) *DescribeUsersInGroupResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetOnlineUsersCount(v int32) *DescribeUsersInGroupResponseBody {
	s.OnlineUsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetRequestId(v string) *DescribeUsersInGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetUserGroupName(v string) *DescribeUsersInGroupResponseBody {
	s.UserGroupName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetUserOuPath(v string) *DescribeUsersInGroupResponseBody {
	s.UserOuPath = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) SetUsersCount(v int32) *DescribeUsersInGroupResponseBody {
	s.UsersCount = &v
	return s
}

func (s *DescribeUsersInGroupResponseBody) Validate() error {
	if s.EndUsers != nil {
		for _, item := range s.EndUsers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsersInGroupResponseBodyEndUsers struct {
	// The connection status.
	//
	// Valid values:
	//
	// 	- 0: disconnected
	//
	// 	- 1: connecting
	//
	// example:
	//
	// 1
	ConnectionStatus *string `json:"ConnectionStatus,omitempty" xml:"ConnectionStatus,omitempty"`
	// The ID of the cloud computer.
	//
	// example:
	//
	// ecd-gx2x1dhsmucyy****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The name of the cloud computer.
	//
	// example:
	//
	// testName
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The display name of the enterprise AD account.
	//
	// example:
	//
	// alice
	DisplayName    *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	DisplayNameNew *string `json:"DisplayNameNew,omitempty" xml:"DisplayNameNew,omitempty"`
	// The email address of the authorized user.
	//
	// example:
	//
	// alice@example.com
	EndUserEmail *string `json:"EndUserEmail,omitempty" xml:"EndUserEmail,omitempty"`
	// The ID of the authorized user.
	//
	// example:
	//
	// alice
	EndUserId *string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty"`
	// The username of the authorized user.
	//
	// example:
	//
	// alice
	EndUserName *string `json:"EndUserName,omitempty" xml:"EndUserName,omitempty"`
	// The mobile number of the authorized user.
	//
	// example:
	//
	// 1381111****
	EndUserPhone *string `json:"EndUserPhone,omitempty" xml:"EndUserPhone,omitempty"`
	// The remarks.
	//
	// example:
	//
	// Note
	EndUserRemark *string `json:"EndUserRemark,omitempty" xml:"EndUserRemark,omitempty"`
	// The user account type.
	//
	// Valid values:
	//
	// 	- SIMPLE: convenience account
	//
	// 	- AD_CONNECTOR: enterprise Active Directory (AD) account
	//
	// example:
	//
	// SIMPLE
	EndUserType *string `json:"EndUserType,omitempty" xml:"EndUserType,omitempty"`
	// The appended information.
	ExternalInfo *DescribeUsersInGroupResponseBodyEndUsersExternalInfo `json:"ExternalInfo,omitempty" xml:"ExternalInfo,omitempty" type:"Struct"`
	// The ID of the cloud computer that is used by the user.
	//
	// example:
	//
	// ud-i896ze8hazpvl****
	UserDesktopId     *string `json:"UserDesktopId,omitempty" xml:"UserDesktopId,omitempty"`
	UserPrincipalName *string `json:"UserPrincipalName,omitempty" xml:"UserPrincipalName,omitempty"`
	// Details about the seats of users.
	UserSetPropertiesModels []*DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels `json:"UserSetPropertiesModels,omitempty" xml:"UserSetPropertiesModels,omitempty" type:"Repeated"`
}

func (s DescribeUsersInGroupResponseBodyEndUsers) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupResponseBodyEndUsers) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetConnectionStatus() *string {
	return s.ConnectionStatus
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetDisplayName() *string {
	return s.DisplayName
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetDisplayNameNew() *string {
	return s.DisplayNameNew
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetEndUserEmail() *string {
	return s.EndUserEmail
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetEndUserId() *string {
	return s.EndUserId
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetEndUserName() *string {
	return s.EndUserName
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetEndUserPhone() *string {
	return s.EndUserPhone
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetEndUserRemark() *string {
	return s.EndUserRemark
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetEndUserType() *string {
	return s.EndUserType
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetExternalInfo() *DescribeUsersInGroupResponseBodyEndUsersExternalInfo {
	return s.ExternalInfo
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetUserDesktopId() *string {
	return s.UserDesktopId
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetUserPrincipalName() *string {
	return s.UserPrincipalName
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) GetUserSetPropertiesModels() []*DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	return s.UserSetPropertiesModels
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetConnectionStatus(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.ConnectionStatus = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDesktopId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DesktopId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDesktopName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DesktopName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDisplayName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DisplayName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetDisplayNameNew(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.DisplayNameNew = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserEmail(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserEmail = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserPhone(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserPhone = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserRemark(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserRemark = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetEndUserType(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.EndUserType = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetExternalInfo(v *DescribeUsersInGroupResponseBodyEndUsersExternalInfo) *DescribeUsersInGroupResponseBodyEndUsers {
	s.ExternalInfo = v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetUserDesktopId(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.UserDesktopId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetUserPrincipalName(v string) *DescribeUsersInGroupResponseBodyEndUsers {
	s.UserPrincipalName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) SetUserSetPropertiesModels(v []*DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) *DescribeUsersInGroupResponseBodyEndUsers {
	s.UserSetPropertiesModels = v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsers) Validate() error {
	if s.ExternalInfo != nil {
		if err := s.ExternalInfo.Validate(); err != nil {
			return err
		}
	}
	if s.UserSetPropertiesModels != nil {
		for _, item := range s.UserSetPropertiesModels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsersInGroupResponseBodyEndUsersExternalInfo struct {
	// The external name.
	//
	// example:
	//
	// nameDemo
	ExternalName *string `json:"ExternalName,omitempty" xml:"ExternalName,omitempty"`
	// The employee ID.
	//
	// example:
	//
	// 123
	JobNumber *string `json:"JobNumber,omitempty" xml:"JobNumber,omitempty"`
}

func (s DescribeUsersInGroupResponseBodyEndUsersExternalInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupResponseBodyEndUsersExternalInfo) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBodyEndUsersExternalInfo) GetExternalName() *string {
	return s.ExternalName
}

func (s *DescribeUsersInGroupResponseBodyEndUsersExternalInfo) GetJobNumber() *string {
	return s.JobNumber
}

func (s *DescribeUsersInGroupResponseBodyEndUsersExternalInfo) SetExternalName(v string) *DescribeUsersInGroupResponseBodyEndUsersExternalInfo {
	s.ExternalName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersExternalInfo) SetJobNumber(v string) *DescribeUsersInGroupResponseBodyEndUsersExternalInfo {
	s.JobNumber = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersExternalInfo) Validate() error {
	return dara.Validate(s)
}

type DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels struct {
	// The property ID.
	//
	// example:
	//
	// 123
	PropertyId *int64 `json:"PropertyId,omitempty" xml:"PropertyId,omitempty"`
	// The property name.
	//
	// example:
	//
	// key
	PropertyKey *string `json:"PropertyKey,omitempty" xml:"PropertyKey,omitempty"`
	// The property type.
	//
	// Valid values:
	//
	// 	- 1: system property
	//
	// 	- 2: custom property
	//
	// example:
	//
	// 1
	PropertyType *int32 `json:"PropertyType,omitempty" xml:"PropertyType,omitempty"`
	// Details about property values.
	PropertyValues []*DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues `json:"PropertyValues,omitempty" xml:"PropertyValues,omitempty" type:"Repeated"`
	// The user ID.
	//
	// example:
	//
	// 123
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The username.
	//
	// example:
	//
	// nameDemo
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
}

func (s DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GetPropertyId() *int64 {
	return s.PropertyId
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GetPropertyKey() *string {
	return s.PropertyKey
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GetPropertyType() *int32 {
	return s.PropertyType
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GetPropertyValues() []*DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues {
	return s.PropertyValues
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) GetUserName() *string {
	return s.UserName
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) SetPropertyId(v int64) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	s.PropertyId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) SetPropertyKey(v string) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	s.PropertyKey = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) SetPropertyType(v int32) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	s.PropertyType = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) SetPropertyValues(v []*DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	s.PropertyValues = v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) SetUserId(v int64) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	s.UserId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) SetUserName(v string) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels {
	s.UserName = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModels) Validate() error {
	if s.PropertyValues != nil {
		for _, item := range s.PropertyValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues struct {
	// The property value.
	//
	// example:
	//
	// value
	PropertyValue *string `json:"PropertyValue,omitempty" xml:"PropertyValue,omitempty"`
	// The ID of the property value.
	//
	// example:
	//
	// 123
	PropertyValueId *int64 `json:"PropertyValueId,omitempty" xml:"PropertyValueId,omitempty"`
}

func (s DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) GoString() string {
	return s.String()
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) GetPropertyValue() *string {
	return s.PropertyValue
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) GetPropertyValueId() *int64 {
	return s.PropertyValueId
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) SetPropertyValue(v string) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValue = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) SetPropertyValueId(v int64) *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues {
	s.PropertyValueId = &v
	return s
}

func (s *DescribeUsersInGroupResponseBodyEndUsersUserSetPropertiesModelsPropertyValues) Validate() error {
	return dara.Validate(s)
}

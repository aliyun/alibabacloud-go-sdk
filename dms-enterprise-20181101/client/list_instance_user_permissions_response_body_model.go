// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstanceUserPermissionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListInstanceUserPermissionsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListInstanceUserPermissionsResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ListInstanceUserPermissionsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListInstanceUserPermissionsResponseBody
	GetSuccess() *bool
	SetTotalCount(v int64) *ListInstanceUserPermissionsResponseBody
	GetTotalCount() *int64
	SetUserPermissions(v *ListInstanceUserPermissionsResponseBodyUserPermissions) *ListInstanceUserPermissionsResponseBody
	GetUserPermissions() *ListInstanceUserPermissionsResponseBodyUserPermissions
}

type ListInstanceUserPermissionsResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UnknownError
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7D162AAE-6501-5691-BF14-D7018F662895
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - true: The request is successful.
	//
	// - false: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of returned entries.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	// The permissions of the user on the instance.
	UserPermissions *ListInstanceUserPermissionsResponseBodyUserPermissions `json:"UserPermissions,omitempty" xml:"UserPermissions,omitempty" type:"Struct"`
}

func (s ListInstanceUserPermissionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListInstanceUserPermissionsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListInstanceUserPermissionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListInstanceUserPermissionsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListInstanceUserPermissionsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListInstanceUserPermissionsResponseBody) GetUserPermissions() *ListInstanceUserPermissionsResponseBodyUserPermissions {
	return s.UserPermissions
}

func (s *ListInstanceUserPermissionsResponseBody) SetErrorCode(v string) *ListInstanceUserPermissionsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBody) SetErrorMessage(v string) *ListInstanceUserPermissionsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBody) SetRequestId(v string) *ListInstanceUserPermissionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBody) SetSuccess(v bool) *ListInstanceUserPermissionsResponseBody {
	s.Success = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBody) SetTotalCount(v int64) *ListInstanceUserPermissionsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBody) SetUserPermissions(v *ListInstanceUserPermissionsResponseBodyUserPermissions) *ListInstanceUserPermissionsResponseBody {
	s.UserPermissions = v
	return s
}

func (s *ListInstanceUserPermissionsResponseBody) Validate() error {
	if s.UserPermissions != nil {
		if err := s.UserPermissions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceUserPermissionsResponseBodyUserPermissions struct {
	UserPermission []*ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission `json:"UserPermission,omitempty" xml:"UserPermission,omitempty" type:"Repeated"`
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissions) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissions) GetUserPermission() []*ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission {
	return s.UserPermission
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissions) SetUserPermission(v []*ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) *ListInstanceUserPermissionsResponseBodyUserPermissions {
	s.UserPermission = v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissions) Validate() error {
	if s.UserPermission != nil {
		for _, item := range s.UserPermission {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission struct {
	// The ID of the instance.
	//
	// example:
	//
	// 174****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The details of permissions.
	PermDetails *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails `json:"PermDetails,omitempty" xml:"PermDetails,omitempty" type:"Struct"`
	// The ID of the user.
	//
	// example:
	//
	// 51****
	UserId *string `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// test_nick_name
	UserNickName *string `json:"UserNickName,omitempty" xml:"UserNickName,omitempty"`
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) GetPermDetails() *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails {
	return s.PermDetails
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) GetUserId() *string {
	return s.UserId
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) GetUserNickName() *string {
	return s.UserNickName
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) SetInstanceId(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.InstanceId = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) SetPermDetails(v *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.PermDetails = v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) SetUserId(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.UserId = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) SetUserNickName(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission {
	s.UserNickName = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermission) Validate() error {
	if s.PermDetails != nil {
		if err := s.PermDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails struct {
	PermDetail []*ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail `json:"PermDetail,omitempty" xml:"PermDetail,omitempty" type:"Repeated"`
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) GetPermDetail() []*ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	return s.PermDetail
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) SetPermDetail(v []*ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails {
	s.PermDetail = v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetails) Validate() error {
	if s.PermDetail != nil {
		for _, item := range s.PermDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail struct {
	// The time when the permissions were granted.
	//
	// example:
	//
	// 2019-12-12 00:00:00
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The time when the permissions expire.
	//
	// example:
	//
	// 2020-12-12 00:00:00
	ExpireDate *string `json:"ExpireDate,omitempty" xml:"ExpireDate,omitempty"`
	// This parameter is reserved.
	//
	// example:
	//
	// XXX
	ExtraData *string `json:"ExtraData,omitempty" xml:"ExtraData,omitempty"`
	// The user who grants the permissions.
	//
	// example:
	//
	// xxx authorization
	OriginFrom *string `json:"OriginFrom,omitempty" xml:"OriginFrom,omitempty"`
	// The type of the permissions. Valid values:
	//
	// 	- LOGIN: the logon permissions
	//
	// 	- PERF: the query permissions on the instance
	//
	// example:
	//
	// LOGIN
	PermType *string `json:"PermType,omitempty" xml:"PermType,omitempty"`
	// The ID of the authorization record.
	//
	// example:
	//
	// 773****
	UserAccessId *string `json:"UserAccessId,omitempty" xml:"UserAccessId,omitempty"`
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) String() string {
	return dara.Prettify(s)
}

func (s ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GoString() string {
	return s.String()
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetCreateDate() *string {
	return s.CreateDate
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetExpireDate() *string {
	return s.ExpireDate
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetExtraData() *string {
	return s.ExtraData
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetOriginFrom() *string {
	return s.OriginFrom
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetPermType() *string {
	return s.PermType
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) GetUserAccessId() *string {
	return s.UserAccessId
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetCreateDate(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.CreateDate = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExpireDate(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExpireDate = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetExtraData(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.ExtraData = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetOriginFrom(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.OriginFrom = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetPermType(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.PermType = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) SetUserAccessId(v string) *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail {
	s.UserAccessId = &v
	return s
}

func (s *ListInstanceUserPermissionsResponseBodyUserPermissionsUserPermissionPermDetailsPermDetail) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListProxyAccessesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *ListProxyAccessesResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *ListProxyAccessesResponseBody
	GetErrorMessage() *string
	SetProxyAccessList(v []*ListProxyAccessesResponseBodyProxyAccessList) *ListProxyAccessesResponseBody
	GetProxyAccessList() []*ListProxyAccessesResponseBodyProxyAccessList
	SetRequestId(v string) *ListProxyAccessesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListProxyAccessesResponseBody
	GetSuccess() *bool
}

type ListProxyAccessesResponseBody struct {
	// The error code returned.
	//
	// example:
	//
	// MissingProxyId
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned.
	//
	// example:
	//
	// ProxyId is mandatory for this action.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The information about the users that are authorized to access the database instance by using the secure access proxy feature.
	ProxyAccessList []*ListProxyAccessesResponseBodyProxyAccessList `json:"ProxyAccessList,omitempty" xml:"ProxyAccessList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// E53D178A-85E9-5E1F-88B6-3CB1FCF2****
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

func (s ListProxyAccessesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListProxyAccessesResponseBody) GoString() string {
	return s.String()
}

func (s *ListProxyAccessesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListProxyAccessesResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ListProxyAccessesResponseBody) GetProxyAccessList() []*ListProxyAccessesResponseBodyProxyAccessList {
	return s.ProxyAccessList
}

func (s *ListProxyAccessesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListProxyAccessesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListProxyAccessesResponseBody) SetErrorCode(v string) *ListProxyAccessesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListProxyAccessesResponseBody) SetErrorMessage(v string) *ListProxyAccessesResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ListProxyAccessesResponseBody) SetProxyAccessList(v []*ListProxyAccessesResponseBodyProxyAccessList) *ListProxyAccessesResponseBody {
	s.ProxyAccessList = v
	return s
}

func (s *ListProxyAccessesResponseBody) SetRequestId(v string) *ListProxyAccessesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListProxyAccessesResponseBody) SetSuccess(v bool) *ListProxyAccessesResponseBody {
	s.Success = &v
	return s
}

func (s *ListProxyAccessesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListProxyAccessesResponseBodyProxyAccessList struct {
	// The username of the database account that is authorized to access the database instance by using the secure access proxy feature.
	//
	// example:
	//
	// MXPL8HalI22m****
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The time when the user is authorized to access the database instance by using the secure access proxy feature.
	//
	// example:
	//
	// 2021-03-31 10:34:18
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The username of the independent database account.
	//
	// example:
	//
	// ****
	IndepAccount *string `json:"IndepAccount,omitempty" xml:"IndepAccount,omitempty"`
	// The ID of the database instance.
	//
	// example:
	//
	// 164****
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The method that is used to authorize the user to access the database instance by using the secure access proxy feature. Valid values:
	//
	// - **Authorization by the Alibaba Cloud Account ()**: The information in the parentheses () indicates the user ID (UID) of the Alibaba Cloud account.
	//
	// - **Authorization by submitting the ticket ()**:The information in the parentheses () indicates the number of the ticket that the user submits to apply for permissions.
	//
	// example:
	//
	// Authorization by the Alibaba Cloud Account(29490401597700****)
	OriginInfo *string `json:"OriginInfo,omitempty" xml:"OriginInfo,omitempty"`
	// The ID that DMS generates after the user is authorized to access the database instance by using the secure access proxy feature. The ID is unique in DMS.
	//
	// example:
	//
	// ****
	ProxyAccessId *int64 `json:"ProxyAccessId,omitempty" xml:"ProxyAccessId,omitempty"`
	// The ID of the secure access proxy.
	//
	// example:
	//
	// 47
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 26****
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The UID of the Alibaba Cloud account.
	//
	// example:
	//
	// 25936669186260****
	UserUid *string `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s ListProxyAccessesResponseBodyProxyAccessList) String() string {
	return dara.Prettify(s)
}

func (s ListProxyAccessesResponseBodyProxyAccessList) GoString() string {
	return s.String()
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetAccessId() *string {
	return s.AccessId
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetIndepAccount() *string {
	return s.IndepAccount
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetOriginInfo() *string {
	return s.OriginInfo
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetProxyAccessId() *int64 {
	return s.ProxyAccessId
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetUserId() *int64 {
	return s.UserId
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetUserName() *string {
	return s.UserName
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) GetUserUid() *string {
	return s.UserUid
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetAccessId(v string) *ListProxyAccessesResponseBodyProxyAccessList {
	s.AccessId = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetGmtCreate(v string) *ListProxyAccessesResponseBodyProxyAccessList {
	s.GmtCreate = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetIndepAccount(v string) *ListProxyAccessesResponseBodyProxyAccessList {
	s.IndepAccount = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetInstanceId(v int64) *ListProxyAccessesResponseBodyProxyAccessList {
	s.InstanceId = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetOriginInfo(v string) *ListProxyAccessesResponseBodyProxyAccessList {
	s.OriginInfo = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetProxyAccessId(v int64) *ListProxyAccessesResponseBodyProxyAccessList {
	s.ProxyAccessId = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetProxyId(v int64) *ListProxyAccessesResponseBodyProxyAccessList {
	s.ProxyId = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetUserId(v int64) *ListProxyAccessesResponseBodyProxyAccessList {
	s.UserId = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetUserName(v string) *ListProxyAccessesResponseBodyProxyAccessList {
	s.UserName = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) SetUserUid(v string) *ListProxyAccessesResponseBodyProxyAccessList {
	s.UserUid = &v
	return s
}

func (s *ListProxyAccessesResponseBodyProxyAccessList) Validate() error {
	return dara.Validate(s)
}

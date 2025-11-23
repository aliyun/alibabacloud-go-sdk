// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProxyAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *GetProxyAccessResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *GetProxyAccessResponseBody
	GetErrorMessage() *string
	SetProxyAccess(v *GetProxyAccessResponseBodyProxyAccess) *GetProxyAccessResponseBody
	GetProxyAccess() *GetProxyAccessResponseBodyProxyAccess
	SetRequestId(v string) *GetProxyAccessResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetProxyAccessResponseBody
	GetSuccess() *bool
}

type GetProxyAccessResponseBody struct {
	// The error code.
	//
	// example:
	//
	// UserNotExist
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The specified user not exists.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The authorization information about the secure access proxy feature.
	ProxyAccess *GetProxyAccessResponseBodyProxyAccess `json:"ProxyAccess,omitempty" xml:"ProxyAccess,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 3CDB8601-AD74-4A47-8114-08E08CD6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetProxyAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProxyAccessResponseBody) GoString() string {
	return s.String()
}

func (s *GetProxyAccessResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *GetProxyAccessResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *GetProxyAccessResponseBody) GetProxyAccess() *GetProxyAccessResponseBodyProxyAccess {
	return s.ProxyAccess
}

func (s *GetProxyAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProxyAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetProxyAccessResponseBody) SetErrorCode(v string) *GetProxyAccessResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetProxyAccessResponseBody) SetErrorMessage(v string) *GetProxyAccessResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetProxyAccessResponseBody) SetProxyAccess(v *GetProxyAccessResponseBodyProxyAccess) *GetProxyAccessResponseBody {
	s.ProxyAccess = v
	return s
}

func (s *GetProxyAccessResponseBody) SetRequestId(v string) *GetProxyAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProxyAccessResponseBody) SetSuccess(v bool) *GetProxyAccessResponseBody {
	s.Success = &v
	return s
}

func (s *GetProxyAccessResponseBody) Validate() error {
	if s.ProxyAccess != nil {
		if err := s.ProxyAccess.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProxyAccessResponseBodyProxyAccess struct {
	// The username of the database account that is authorized to enable the secure access proxy feature for an instance.
	//
	// example:
	//
	// hObpgEXoca42q***
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The time when the user is authorized to enable the secure access proxy feature for an instance.
	//
	// example:
	//
	// 1643034647
	GmtCreate *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The username of the independent database account.
	//
	// example:
	//
	// ***
	IndepAccount *string `json:"IndepAccount,omitempty" xml:"IndepAccount,omitempty"`
	// The ID of the instance for which the secure access proxy feature is enabled.
	//
	// example:
	//
	// 1922545
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The method that is used to authorize the user to enable the secure access proxy feature for an instance. Valid values:
	//
	// 	- **Authorization by the Alibaba Cloud Account ()**: The information in the parentheses () indicates the ID of the Alibaba Cloud account.
	//
	// 	- **Authorization by submitting the ticket ()**:The information in the parentheses () indicates the number of the ticket that the user submits to apply for permissions.
	//
	// example:
	//
	// Authorization by the Alibaba Cloud account (29490401597700\\*\\*\\*\\*)
	OriginInfo *string `json:"OriginInfo,omitempty" xml:"OriginInfo,omitempty"`
	// The ID that DMS generates after the user is authorized to enable the secure access proxy feature for an instance. The ID is unique in DMS. You can call the [ListProxyAccesses](https://help.aliyun.com/document_detail/295386.html) operation to query the ID.
	//
	// example:
	//
	// 2002
	ProxyAccessId *int64 `json:"ProxyAccessId,omitempty" xml:"ProxyAccessId,omitempty"`
	// The ID of the secure access proxy.
	//
	// >  You can call the [ListProxies](https://help.aliyun.com/document_detail/295371.html) operation to query the ID of the secure access proxy.
	//
	// example:
	//
	// 1905
	ProxyId *int64 `json:"ProxyId,omitempty" xml:"ProxyId,omitempty"`
	// The ID of the user.
	//
	// example:
	//
	// 12***
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// user
	UserName *string `json:"UserName,omitempty" xml:"UserName,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 25936669186260****
	UserUid *string `json:"UserUid,omitempty" xml:"UserUid,omitempty"`
}

func (s GetProxyAccessResponseBodyProxyAccess) String() string {
	return dara.Prettify(s)
}

func (s GetProxyAccessResponseBodyProxyAccess) GoString() string {
	return s.String()
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetAccessId() *string {
	return s.AccessId
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetIndepAccount() *string {
	return s.IndepAccount
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetOriginInfo() *string {
	return s.OriginInfo
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetProxyAccessId() *int64 {
	return s.ProxyAccessId
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetProxyId() *int64 {
	return s.ProxyId
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetUserId() *int64 {
	return s.UserId
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetUserName() *string {
	return s.UserName
}

func (s *GetProxyAccessResponseBodyProxyAccess) GetUserUid() *string {
	return s.UserUid
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetAccessId(v string) *GetProxyAccessResponseBodyProxyAccess {
	s.AccessId = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetGmtCreate(v string) *GetProxyAccessResponseBodyProxyAccess {
	s.GmtCreate = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetIndepAccount(v string) *GetProxyAccessResponseBodyProxyAccess {
	s.IndepAccount = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetInstanceId(v int64) *GetProxyAccessResponseBodyProxyAccess {
	s.InstanceId = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetOriginInfo(v string) *GetProxyAccessResponseBodyProxyAccess {
	s.OriginInfo = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetProxyAccessId(v int64) *GetProxyAccessResponseBodyProxyAccess {
	s.ProxyAccessId = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetProxyId(v int64) *GetProxyAccessResponseBodyProxyAccess {
	s.ProxyId = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetUserId(v int64) *GetProxyAccessResponseBodyProxyAccess {
	s.UserId = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetUserName(v string) *GetProxyAccessResponseBodyProxyAccess {
	s.UserName = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) SetUserUid(v string) *GetProxyAccessResponseBodyProxyAccess {
	s.UserUid = &v
	return s
}

func (s *GetProxyAccessResponseBodyProxyAccess) Validate() error {
	return dara.Validate(s)
}

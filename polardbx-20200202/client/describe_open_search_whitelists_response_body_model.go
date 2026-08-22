// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeOpenSearchWhitelistsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) *DescribeOpenSearchWhitelistsResponseBody
	GetAccessDeniedDetail() *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail
	SetData(v *DescribeOpenSearchWhitelistsResponseBodyData) *DescribeOpenSearchWhitelistsResponseBody
	GetData() *DescribeOpenSearchWhitelistsResponseBodyData
	SetRequestId(v string) *DescribeOpenSearchWhitelistsResponseBody
	GetRequestId() *string
}

type DescribeOpenSearchWhitelistsResponseBody struct {
	// The details of the access denial.
	AccessDeniedDetail *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty" type:"Struct"`
	// The monitoring data.
	Data *DescribeOpenSearchWhitelistsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// D6A4256F-7B83-5BD7-9AC0-72E1FAC05330
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeOpenSearchWhitelistsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchWhitelistsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchWhitelistsResponseBody) GetAccessDeniedDetail() *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	return s.AccessDeniedDetail
}

func (s *DescribeOpenSearchWhitelistsResponseBody) GetData() *DescribeOpenSearchWhitelistsResponseBodyData {
	return s.Data
}

func (s *DescribeOpenSearchWhitelistsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeOpenSearchWhitelistsResponseBody) SetAccessDeniedDetail(v *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) *DescribeOpenSearchWhitelistsResponseBody {
	s.AccessDeniedDetail = v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBody) SetData(v *DescribeOpenSearchWhitelistsResponseBodyData) *DescribeOpenSearchWhitelistsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBody) SetRequestId(v string) *DescribeOpenSearchWhitelistsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBody) Validate() error {
	if s.AccessDeniedDetail != nil {
		if err := s.AccessDeniedDetail.Validate(); err != nil {
			return err
		}
	}
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail struct {
	// The authentication action.
	//
	// example:
	//
	// xxx
	AuthAction *string `json:"AuthAction,omitempty" xml:"AuthAction,omitempty"`
	// The display name of the authentication principal.
	//
	// example:
	//
	// xxx
	AuthPrincipalDisplayName *string `json:"AuthPrincipalDisplayName,omitempty" xml:"AuthPrincipalDisplayName,omitempty"`
	// The owner ID of the authentication principal.
	//
	// example:
	//
	// 111
	AuthPrincipalOwnerId *string `json:"AuthPrincipalOwnerId,omitempty" xml:"AuthPrincipalOwnerId,omitempty"`
	// The authentication principal type.
	//
	// example:
	//
	// 222
	AuthPrincipalType *string `json:"AuthPrincipalType,omitempty" xml:"AuthPrincipalType,omitempty"`
	// The encoded diagnostic message.
	//
	// example:
	//
	// AQEAAAAAaKPfwjY0MzMyODRGLUZCQkQtNTA1RS04MUUxLTc5NTkzODk2MUIzMg==
	EncodedDiagnosticMessage *string `json:"EncodedDiagnosticMessage,omitempty" xml:"EncodedDiagnosticMessage,omitempty"`
	// The type of the missing permission.
	//
	// example:
	//
	// ImplicitDeny
	NoPermissionType *string `json:"NoPermissionType,omitempty" xml:"NoPermissionType,omitempty"`
	// The policy type.
	//
	// example:
	//
	// PRIORITY
	PolicyType *string `json:"PolicyType,omitempty" xml:"PolicyType,omitempty"`
}

func (s DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetAuthAction() *string {
	return s.AuthAction
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetAuthPrincipalDisplayName() *string {
	return s.AuthPrincipalDisplayName
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetAuthPrincipalOwnerId() *string {
	return s.AuthPrincipalOwnerId
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetAuthPrincipalType() *string {
	return s.AuthPrincipalType
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetEncodedDiagnosticMessage() *string {
	return s.EncodedDiagnosticMessage
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetNoPermissionType() *string {
	return s.NoPermissionType
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) GetPolicyType() *string {
	return s.PolicyType
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetAuthAction(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.AuthAction = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetAuthPrincipalDisplayName(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalDisplayName = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetAuthPrincipalOwnerId(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalOwnerId = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetAuthPrincipalType(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.AuthPrincipalType = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetEncodedDiagnosticMessage(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.EncodedDiagnosticMessage = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetNoPermissionType(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.NoPermissionType = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) SetPolicyType(v string) *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail {
	s.PolicyType = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyAccessDeniedDetail) Validate() error {
	return dara.Validate(s)
}

type DescribeOpenSearchWhitelistsResponseBodyData struct {
	// The type of the Internet IPv4 whitelist addresses.
	Whitelists []*DescribeOpenSearchWhitelistsResponseBodyDataWhitelists `json:"Whitelists,omitempty" xml:"Whitelists,omitempty" type:"Repeated"`
}

func (s DescribeOpenSearchWhitelistsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchWhitelistsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchWhitelistsResponseBodyData) GetWhitelists() []*DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	return s.Whitelists
}

func (s *DescribeOpenSearchWhitelistsResponseBodyData) SetWhitelists(v []*DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) *DescribeOpenSearchWhitelistsResponseBodyData {
	s.Whitelists = v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyData) Validate() error {
	if s.Whitelists != nil {
		for _, item := range s.Whitelists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeOpenSearchWhitelistsResponseBodyDataWhitelists struct {
	// The creation time.
	//
	// example:
	//
	// 2026-07-22T02:26:08Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The ID of the group to which the instance belongs.
	//
	// example:
	//
	// GID_QMPRUNTIME_BROADCAST_TASK_CONSUMER_GROUP
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The name of the whitelist group.
	//
	// example:
	//
	// ack_worker_new
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP address list.
	//
	// example:
	//
	// []
	Ips *string `json:"Ips,omitempty" xml:"Ips,omitempty"`
	// The network type. Only VPC is supported.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The policy remarks.
	//
	// example:
	//
	// vpc-t4nt9qxfgbzab587cshhc
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The time when the task was last updated, in timestamp format.
	//
	// example:
	//
	// 0001-01-01T00:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) String() string {
	return dara.Prettify(s)
}

func (s DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GoString() string {
	return s.String()
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetGroupId() *string {
	return s.GroupId
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetIps() *string {
	return s.Ips
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetRemark() *string {
	return s.Remark
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetCreateTime(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.CreateTime = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetGroupId(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.GroupId = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetGroupName(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.GroupName = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetIps(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.Ips = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetNetworkType(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.NetworkType = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetRemark(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.Remark = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) SetUpdateTime(v string) *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists {
	s.UpdateTime = &v
	return s
}

func (s *DescribeOpenSearchWhitelistsResponseBodyDataWhitelists) Validate() error {
	return dara.Validate(s)
}

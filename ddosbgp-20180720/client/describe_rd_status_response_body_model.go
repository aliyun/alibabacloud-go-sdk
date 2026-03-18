// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRdStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCurrentUid(v string) *DescribeRdStatusResponseBody
	GetCurrentUid() *string
	SetCurrentUidType(v string) *DescribeRdStatusResponseBody
	GetCurrentUidType() *string
	SetEnabled(v bool) *DescribeRdStatusResponseBody
	GetEnabled() *bool
	SetLocalEnable(v bool) *DescribeRdStatusResponseBody
	GetLocalEnable() *bool
	SetMasterUid(v string) *DescribeRdStatusResponseBody
	GetMasterUid() *string
	SetRemoteEnable(v bool) *DescribeRdStatusResponseBody
	GetRemoteEnable() *bool
	SetRequestId(v string) *DescribeRdStatusResponseBody
	GetRequestId() *string
	SetRootUid(v string) *DescribeRdStatusResponseBody
	GetRootUid() *string
	SetServicePrincipalEnabled(v bool) *DescribeRdStatusResponseBody
	GetServicePrincipalEnabled() *bool
}

type DescribeRdStatusResponseBody struct {
	// The Alibaba Cloud account ID of the current account.
	//
	// example:
	//
	// 125085778340****
	CurrentUid *string `json:"CurrentUid,omitempty" xml:"CurrentUid,omitempty"`
	// The type of the Alibaba Cloud account. Valid values:
	//
	// 	- **MasterAccount**: management account.
	//
	// 	- **DelegatedAdminAccount**: delegated administrator account.
	//
	// 	- **MemberAccount**: member.
	//
	// example:
	//
	// MemberAccount
	CurrentUidType *string `json:"CurrentUidType,omitempty" xml:"CurrentUidType,omitempty"`
	// Indicates whether the multi-account management feature is enabled for Anti-DDoS Origin.
	//
	// example:
	//
	// false
	Enabled *bool `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	// Indicates whether the multi-account management feature is enabled for the current account in Anti-DDoS Origin.
	//
	// example:
	//
	// false
	LocalEnable *bool `json:"LocalEnable,omitempty" xml:"LocalEnable,omitempty"`
	// The Alibaba Cloud account ID of the management account in the resource directory.
	//
	// example:
	//
	// 125085778340****
	MasterUid *string `json:"MasterUid,omitempty" xml:"MasterUid,omitempty"`
	// Indicates whether Resource Directory is enabled in the [Resource Management console](https://resourcemanager.console.aliyun.com).
	//
	// example:
	//
	// false
	RemoteEnable *bool `json:"RemoteEnable,omitempty" xml:"RemoteEnable,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 1B0F7EC6-51D7-4D70-B0EC-CD8A9E998D86
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The Alibaba Cloud account ID of the management account for which the multi-account management feature is enabled in Anti-DDoS Origin.
	//
	// example:
	//
	// 125085778340****
	RootUid *string `json:"RootUid,omitempty" xml:"RootUid,omitempty"`
	// Indicates whether the trusted service is enabled.
	//
	// example:
	//
	// false
	ServicePrincipalEnabled *bool `json:"ServicePrincipalEnabled,omitempty" xml:"ServicePrincipalEnabled,omitempty"`
}

func (s DescribeRdStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRdStatusResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRdStatusResponseBody) GetCurrentUid() *string {
	return s.CurrentUid
}

func (s *DescribeRdStatusResponseBody) GetCurrentUidType() *string {
	return s.CurrentUidType
}

func (s *DescribeRdStatusResponseBody) GetEnabled() *bool {
	return s.Enabled
}

func (s *DescribeRdStatusResponseBody) GetLocalEnable() *bool {
	return s.LocalEnable
}

func (s *DescribeRdStatusResponseBody) GetMasterUid() *string {
	return s.MasterUid
}

func (s *DescribeRdStatusResponseBody) GetRemoteEnable() *bool {
	return s.RemoteEnable
}

func (s *DescribeRdStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRdStatusResponseBody) GetRootUid() *string {
	return s.RootUid
}

func (s *DescribeRdStatusResponseBody) GetServicePrincipalEnabled() *bool {
	return s.ServicePrincipalEnabled
}

func (s *DescribeRdStatusResponseBody) SetCurrentUid(v string) *DescribeRdStatusResponseBody {
	s.CurrentUid = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetCurrentUidType(v string) *DescribeRdStatusResponseBody {
	s.CurrentUidType = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetEnabled(v bool) *DescribeRdStatusResponseBody {
	s.Enabled = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetLocalEnable(v bool) *DescribeRdStatusResponseBody {
	s.LocalEnable = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetMasterUid(v string) *DescribeRdStatusResponseBody {
	s.MasterUid = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetRemoteEnable(v bool) *DescribeRdStatusResponseBody {
	s.RemoteEnable = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetRequestId(v string) *DescribeRdStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetRootUid(v string) *DescribeRdStatusResponseBody {
	s.RootUid = &v
	return s
}

func (s *DescribeRdStatusResponseBody) SetServicePrincipalEnabled(v bool) *DescribeRdStatusResponseBody {
	s.ServicePrincipalEnabled = &v
	return s
}

func (s *DescribeRdStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

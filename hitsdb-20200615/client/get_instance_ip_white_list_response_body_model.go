// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceIpWhiteListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetInstanceIpWhiteListResponseBody
	GetAccessDeniedDetail() *string
	SetGroupList(v []*GetInstanceIpWhiteListResponseBodyGroupList) *GetInstanceIpWhiteListResponseBody
	GetGroupList() []*GetInstanceIpWhiteListResponseBodyGroupList
	SetInstanceId(v string) *GetInstanceIpWhiteListResponseBody
	GetInstanceId() *string
	SetIpList(v []*string) *GetInstanceIpWhiteListResponseBody
	GetIpList() []*string
	SetRequestId(v string) *GetInstanceIpWhiteListResponseBody
	GetRequestId() *string
}

type GetInstanceIpWhiteListResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// {"AuthAction":"xxx","AuthPrincipalDisplayName":"222","AuthPrincipalOwnerId":"111","AuthPrincipalType":"SubUser",,"NoPermissionType":"ImplicitDeny","PolicyType":"AccountLevelIdentityBasedPolicy","EncodedDiagnosticMessage":"xxxxxx"}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The details about the IP address whitelists.
	GroupList []*GetInstanceIpWhiteListResponseBodyGroupList `json:"GroupList,omitempty" xml:"GroupList,omitempty" type:"Repeated"`
	// The ID of the Lindorm instance.
	//
	// example:
	//
	// ld-bp1z3506imz2f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The list of IP addresses in the whitelist of the instance.
	IpList []*string `json:"IpList,omitempty" xml:"IpList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 1D1F6F4D-9203-53E7-84E9-5376B4657E63
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetInstanceIpWhiteListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhiteListResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetInstanceIpWhiteListResponseBody) GetGroupList() []*GetInstanceIpWhiteListResponseBodyGroupList {
	return s.GroupList
}

func (s *GetInstanceIpWhiteListResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceIpWhiteListResponseBody) GetIpList() []*string {
	return s.IpList
}

func (s *GetInstanceIpWhiteListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceIpWhiteListResponseBody) SetAccessDeniedDetail(v string) *GetInstanceIpWhiteListResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetGroupList(v []*GetInstanceIpWhiteListResponseBodyGroupList) *GetInstanceIpWhiteListResponseBody {
	s.GroupList = v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetInstanceId(v string) *GetInstanceIpWhiteListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetIpList(v []*string) *GetInstanceIpWhiteListResponseBody {
	s.IpList = v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) SetRequestId(v string) *GetInstanceIpWhiteListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBody) Validate() error {
	if s.GroupList != nil {
		for _, item := range s.GroupList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetInstanceIpWhiteListResponseBodyGroupList struct {
	// The name of the IP address whitelist.
	//
	// example:
	//
	// test
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The IP addresses in the whitelist.
	//
	// example:
	//
	// 192.168.1.0/24
	SecurityIpList *string `json:"SecurityIpList,omitempty" xml:"SecurityIpList,omitempty"`
}

func (s GetInstanceIpWhiteListResponseBodyGroupList) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceIpWhiteListResponseBodyGroupList) GoString() string {
	return s.String()
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) GetGroupName() *string {
	return s.GroupName
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) GetSecurityIpList() *string {
	return s.SecurityIpList
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) SetGroupName(v string) *GetInstanceIpWhiteListResponseBodyGroupList {
	s.GroupName = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) SetSecurityIpList(v string) *GetInstanceIpWhiteListResponseBodyGroupList {
	s.SecurityIpList = &v
	return s
}

func (s *GetInstanceIpWhiteListResponseBodyGroupList) Validate() error {
	return dara.Validate(s)
}

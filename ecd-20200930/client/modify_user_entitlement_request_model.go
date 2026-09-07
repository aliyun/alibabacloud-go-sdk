// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserEntitlementRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizeDesktopId(v []*string) *ModifyUserEntitlementRequest
	GetAuthorizeDesktopId() []*string
	SetEndUserId(v []*string) *ModifyUserEntitlementRequest
	GetEndUserId() []*string
	SetRegionId(v string) *ModifyUserEntitlementRequest
	GetRegionId() *string
	SetRevokeDesktopId(v []*string) *ModifyUserEntitlementRequest
	GetRevokeDesktopId() []*string
}

type ModifyUserEntitlementRequest struct {
	// The list of cloud computer IDs for which to add authorized users.
	AuthorizeDesktopId []*string `json:"AuthorizeDesktopId,omitempty" xml:"AuthorizeDesktopId,omitempty" type:"Repeated"`
	// The list of user IDs (usernames).
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The region ID. You can call [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) to query the regions supported by Elastic Desktop Service.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The list of cloud computer IDs for which to remove authorized users.
	RevokeDesktopId []*string `json:"RevokeDesktopId,omitempty" xml:"RevokeDesktopId,omitempty" type:"Repeated"`
}

func (s ModifyUserEntitlementRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserEntitlementRequest) GoString() string {
	return s.String()
}

func (s *ModifyUserEntitlementRequest) GetAuthorizeDesktopId() []*string {
	return s.AuthorizeDesktopId
}

func (s *ModifyUserEntitlementRequest) GetEndUserId() []*string {
	return s.EndUserId
}

func (s *ModifyUserEntitlementRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyUserEntitlementRequest) GetRevokeDesktopId() []*string {
	return s.RevokeDesktopId
}

func (s *ModifyUserEntitlementRequest) SetAuthorizeDesktopId(v []*string) *ModifyUserEntitlementRequest {
	s.AuthorizeDesktopId = v
	return s
}

func (s *ModifyUserEntitlementRequest) SetEndUserId(v []*string) *ModifyUserEntitlementRequest {
	s.EndUserId = v
	return s
}

func (s *ModifyUserEntitlementRequest) SetRegionId(v string) *ModifyUserEntitlementRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyUserEntitlementRequest) SetRevokeDesktopId(v []*string) *ModifyUserEntitlementRequest {
	s.RevokeDesktopId = v
	return s
}

func (s *ModifyUserEntitlementRequest) Validate() error {
	return dara.Validate(s)
}

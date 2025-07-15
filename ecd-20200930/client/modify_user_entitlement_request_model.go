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
	// The IDs of the cloud computers to which you want to add end users.
	AuthorizeDesktopId []*string `json:"AuthorizeDesktopId,omitempty" xml:"AuthorizeDesktopId,omitempty" type:"Repeated"`
	// The ID of the users.
	EndUserId []*string `json:"EndUserId,omitempty" xml:"EndUserId,omitempty" type:"Repeated"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/196646.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the cloud computers whose end users you want to remove.
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

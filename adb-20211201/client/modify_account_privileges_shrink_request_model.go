// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPrivilegesShrinkRequest
	GetAccountName() *string
	SetAccountPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest
	GetAccountPrivilegesShrink() *string
	SetDBClusterId(v string) *ModifyAccountPrivilegesShrinkRequest
	GetDBClusterId() *string
	SetPromqlInsertPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest
	GetPromqlInsertPrivilegesShrink() *string
	SetPromqlSelectNodePercentage(v float64) *ModifyAccountPrivilegesShrinkRequest
	GetPromqlSelectNodePercentage() *float64
	SetPromqlSelectPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest
	GetPromqlSelectPrivilegesShrink() *string
	SetRegionId(v string) *ModifyAccountPrivilegesShrinkRequest
	GetRegionId() *string
	SetResourceGroupName(v string) *ModifyAccountPrivilegesShrinkRequest
	GetResourceGroupName() *string
}

type ModifyAccountPrivilegesShrinkRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The list of granted permissions.
	AccountPrivilegesShrink *string `json:"AccountPrivileges,omitempty" xml:"AccountPrivileges,omitempty"`
	// <props="china">The cluster ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The cluster ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k5p066e1a****
	DBClusterId                  *string  `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PromqlInsertPrivilegesShrink *string  `json:"PromqlInsertPrivileges,omitempty" xml:"PromqlInsertPrivileges,omitempty"`
	PromqlSelectNodePercentage   *float64 `json:"PromqlSelectNodePercentage,omitempty" xml:"PromqlSelectNodePercentage,omitempty"`
	PromqlSelectPrivilegesShrink *string  `json:"PromqlSelectPrivileges,omitempty" xml:"PromqlSelectPrivileges,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId          *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupName *string `json:"ResourceGroupName,omitempty" xml:"ResourceGroupName,omitempty"`
}

func (s ModifyAccountPrivilegesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetAccountPrivilegesShrink() *string {
	return s.AccountPrivilegesShrink
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetPromqlInsertPrivilegesShrink() *string {
	return s.PromqlInsertPrivilegesShrink
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetPromqlSelectNodePercentage() *float64 {
	return s.PromqlSelectNodePercentage
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetPromqlSelectPrivilegesShrink() *string {
	return s.PromqlSelectPrivilegesShrink
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountPrivilegesShrinkRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetAccountName(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetAccountPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.AccountPrivilegesShrink = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetDBClusterId(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetPromqlInsertPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.PromqlInsertPrivilegesShrink = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetPromqlSelectNodePercentage(v float64) *ModifyAccountPrivilegesShrinkRequest {
	s.PromqlSelectNodePercentage = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetPromqlSelectPrivilegesShrink(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.PromqlSelectPrivilegesShrink = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetRegionId(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) SetResourceGroupName(v string) *ModifyAccountPrivilegesShrinkRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ModifyAccountPrivilegesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

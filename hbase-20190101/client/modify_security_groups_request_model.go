// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ModifySecurityGroupsRequest
	GetClusterId() *string
	SetSecurityGroupIds(v string) *ModifySecurityGroupsRequest
	GetSecurityGroupIds() *string
}

type ModifySecurityGroupsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// hb-bp16f1441y6p2kv**
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// sg-t4ng4yyc916o81nu****,sg-x4gg4dyc9d6w********
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
}

func (s ModifySecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ModifySecurityGroupsRequest) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *ModifySecurityGroupsRequest) SetClusterId(v string) *ModifySecurityGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ModifySecurityGroupsRequest) SetSecurityGroupIds(v string) *ModifySecurityGroupsRequest {
	s.SecurityGroupIds = &v
	return s
}

func (s *ModifySecurityGroupsRequest) Validate() error {
	return dara.Validate(s)
}

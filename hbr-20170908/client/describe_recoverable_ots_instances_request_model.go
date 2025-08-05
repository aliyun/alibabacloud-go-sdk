// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecoverableOtsInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCrossAccountRoleName(v string) *DescribeRecoverableOtsInstancesRequest
	GetCrossAccountRoleName() *string
	SetCrossAccountType(v string) *DescribeRecoverableOtsInstancesRequest
	GetCrossAccountType() *string
	SetCrossAccountUserId(v int64) *DescribeRecoverableOtsInstancesRequest
	GetCrossAccountUserId() *int64
}

type DescribeRecoverableOtsInstancesRequest struct {
	// The name of the Resource Access Management (RAM) role that is created within the source Alibaba Cloud account and assigned to the current Alibaba Cloud account to authorize the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// hbrcrossrole
	CrossAccountRoleName *string `json:"CrossAccountRoleName,omitempty" xml:"CrossAccountRoleName,omitempty"`
	// Specifies whether data is backed up within the same Alibaba Cloud account or across Alibaba Cloud accounts. Valid values:
	//
	// 	- SELF_ACCOUNT: Data is backed up within the same Alibaba Cloud account.
	//
	// 	- CROSS_ACCOUNT: Data is backed up across Alibaba Cloud accounts.
	//
	// example:
	//
	// CROSS_ACCOUNT
	CrossAccountType *string `json:"CrossAccountType,omitempty" xml:"CrossAccountType,omitempty"`
	// The ID of the source Alibaba Cloud account that authorizes the current Alibaba Cloud account to back up data across Alibaba Cloud accounts.
	//
	// example:
	//
	// 1440155109798732
	CrossAccountUserId *int64 `json:"CrossAccountUserId,omitempty" xml:"CrossAccountUserId,omitempty"`
}

func (s DescribeRecoverableOtsInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecoverableOtsInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRecoverableOtsInstancesRequest) GetCrossAccountRoleName() *string {
	return s.CrossAccountRoleName
}

func (s *DescribeRecoverableOtsInstancesRequest) GetCrossAccountType() *string {
	return s.CrossAccountType
}

func (s *DescribeRecoverableOtsInstancesRequest) GetCrossAccountUserId() *int64 {
	return s.CrossAccountUserId
}

func (s *DescribeRecoverableOtsInstancesRequest) SetCrossAccountRoleName(v string) *DescribeRecoverableOtsInstancesRequest {
	s.CrossAccountRoleName = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesRequest) SetCrossAccountType(v string) *DescribeRecoverableOtsInstancesRequest {
	s.CrossAccountType = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesRequest) SetCrossAccountUserId(v int64) *DescribeRecoverableOtsInstancesRequest {
	s.CrossAccountUserId = &v
	return s
}

func (s *DescribeRecoverableOtsInstancesRequest) Validate() error {
	return dara.Validate(s)
}

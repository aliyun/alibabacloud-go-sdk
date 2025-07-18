// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBResourceGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *ModifyDBResourceGroupShrinkRequest
	GetDBInstanceId() *string
	SetOwnerId(v int64) *ModifyDBResourceGroupShrinkRequest
	GetOwnerId() *int64
	SetResourceGroupItemsShrink(v string) *ModifyDBResourceGroupShrinkRequest
	GetResourceGroupItemsShrink() *string
}

type ModifyDBResourceGroupShrinkRequest struct {
	// The instance ID.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the information about all AnalyticDB for PostgreSQL instances within a region, including instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-xxxxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The information about the resource group.
	//
	// This parameter is required.
	ResourceGroupItemsShrink *string `json:"ResourceGroupItems,omitempty" xml:"ResourceGroupItems,omitempty"`
}

func (s ModifyDBResourceGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBResourceGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBResourceGroupShrinkRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBResourceGroupShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDBResourceGroupShrinkRequest) GetResourceGroupItemsShrink() *string {
	return s.ResourceGroupItemsShrink
}

func (s *ModifyDBResourceGroupShrinkRequest) SetDBInstanceId(v string) *ModifyDBResourceGroupShrinkRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetOwnerId(v int64) *ModifyDBResourceGroupShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) SetResourceGroupItemsShrink(v string) *ModifyDBResourceGroupShrinkRequest {
	s.ResourceGroupItemsShrink = &v
	return s
}

func (s *ModifyDBResourceGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

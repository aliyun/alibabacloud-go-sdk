// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDBResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DeleteDBResourceGroupRequest
	GetDBClusterId() *string
	SetGroupName(v string) *DeleteDBResourceGroupRequest
	GetGroupName() *string
}

type DeleteDBResourceGroupRequest struct {
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The name of the resource group.
	//
	// >  You can call the [DescribeDBResourceGroup](https://help.aliyun.com/document_detail/612410.html) operation to query the information about resource groups of an AnalyticDB for MySQL cluster, including resource group names.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
}

func (s DeleteDBResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDBResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteDBResourceGroupRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DeleteDBResourceGroupRequest) GetGroupName() *string {
	return s.GroupName
}

func (s *DeleteDBResourceGroupRequest) SetDBClusterId(v string) *DeleteDBResourceGroupRequest {
	s.DBClusterId = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) SetGroupName(v string) *DeleteDBResourceGroupRequest {
	s.GroupName = &v
	return s
}

func (s *DeleteDBResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}

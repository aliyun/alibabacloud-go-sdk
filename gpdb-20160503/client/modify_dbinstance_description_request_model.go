// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDescriptionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceDescription() *string
	SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest
	GetDBInstanceId() *string
	SetResourceGroupId(v string) *ModifyDBInstanceDescriptionRequest
	GetResourceGroupId() *string
}

type ModifyDBInstanceDescriptionRequest struct {
	// The description of the instance.
	//
	// The description must be 2 to 256 characters in length. It cannot start with http:// or https://.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	DBInstanceDescription *string `json:"DBInstanceDescription,omitempty" xml:"DBInstanceDescription,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ModifyDBInstanceDescriptionRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDescriptionRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceDescription() *string {
	return s.DBInstanceDescription
}

func (s *ModifyDBInstanceDescriptionRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceDescriptionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceDescription(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceDescription = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetDBInstanceId(v string) *ModifyDBInstanceDescriptionRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) SetResourceGroupId(v string) *ModifyDBInstanceDescriptionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifyDBInstanceDescriptionRequest) Validate() error {
	return dara.Validate(s)
}

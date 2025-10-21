// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeType(v string) *ModifyDBInstanceAttributeRequest
	GetAttributeType() *string
	SetAttributeValue(v string) *ModifyDBInstanceAttributeRequest
	GetAttributeValue() *string
	SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest
	GetDBInstanceId() *string
	SetProduct(v string) *ModifyDBInstanceAttributeRequest
	GetProduct() *string
	SetRegionId(v string) *ModifyDBInstanceAttributeRequest
	GetRegionId() *string
}

type ModifyDBInstanceAttributeRequest struct {
	// The configuration that you want to modify.
	//
	// 	- MaintainTime: the O\\&M time
	//
	// 	- DBInstanceDescription: the cluster name
	//
	// This parameter is required.
	//
	// example:
	//
	// DBInstanceDescription
	AttributeType *string `json:"AttributeType,omitempty" xml:"AttributeType,omitempty"`
	// The new value of the configuration.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
	// The cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cc-xxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The code of the cloud service.
	//
	// example:
	//
	// clickhouse
	Product *string `json:"Product,omitempty" xml:"Product,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ModifyDBInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceAttributeRequest) GetAttributeType() *string {
	return s.AttributeType
}

func (s *ModifyDBInstanceAttributeRequest) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *ModifyDBInstanceAttributeRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifyDBInstanceAttributeRequest) GetProduct() *string {
	return s.Product
}

func (s *ModifyDBInstanceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDBInstanceAttributeRequest) SetAttributeType(v string) *ModifyDBInstanceAttributeRequest {
	s.AttributeType = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetAttributeValue(v string) *ModifyDBInstanceAttributeRequest {
	s.AttributeValue = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetDBInstanceId(v string) *ModifyDBInstanceAttributeRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetProduct(v string) *ModifyDBInstanceAttributeRequest {
	s.Product = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) SetRegionId(v string) *ModifyDBInstanceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDBInstanceAttributeRequest) Validate() error {
	return dara.Validate(s)
}

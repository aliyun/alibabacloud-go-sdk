// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccountPrivilegesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccountName(v string) *ModifyAccountPrivilegesRequest
	GetAccountName() *string
	SetAccountPrivileges(v []*ModifyAccountPrivilegesRequestAccountPrivileges) *ModifyAccountPrivilegesRequest
	GetAccountPrivileges() []*ModifyAccountPrivilegesRequestAccountPrivileges
	SetDBClusterId(v string) *ModifyAccountPrivilegesRequest
	GetDBClusterId() *string
	SetPromqlInsertPrivileges(v []*string) *ModifyAccountPrivilegesRequest
	GetPromqlInsertPrivileges() []*string
	SetPromqlSelectNodePercentage(v float64) *ModifyAccountPrivilegesRequest
	GetPromqlSelectNodePercentage() *float64
	SetPromqlSelectPrivileges(v []*string) *ModifyAccountPrivilegesRequest
	GetPromqlSelectPrivileges() []*string
	SetRegionId(v string) *ModifyAccountPrivilegesRequest
	GetRegionId() *string
	SetResourceGroupName(v string) *ModifyAccountPrivilegesRequest
	GetResourceGroupName() *string
}

type ModifyAccountPrivilegesRequest struct {
	// The name of the database account.
	//
	// This parameter is required.
	//
	// example:
	//
	// account1
	AccountName *string `json:"AccountName,omitempty" xml:"AccountName,omitempty"`
	// The list of granted permissions.
	AccountPrivileges []*ModifyAccountPrivilegesRequestAccountPrivileges `json:"AccountPrivileges,omitempty" xml:"AccountPrivileges,omitempty" type:"Repeated"`
	// <props="china">The cluster ID of the Enterprise Edition, Basic Edition, or Data Lakehouse Edition cluster.
	//
	// <props="intl">The cluster ID of the Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp1k5p066e1a****
	DBClusterId                *string   `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	PromqlInsertPrivileges     []*string `json:"PromqlInsertPrivileges,omitempty" xml:"PromqlInsertPrivileges,omitempty" type:"Repeated"`
	PromqlSelectNodePercentage *float64  `json:"PromqlSelectNodePercentage,omitempty" xml:"PromqlSelectNodePercentage,omitempty"`
	PromqlSelectPrivileges     []*string `json:"PromqlSelectPrivileges,omitempty" xml:"PromqlSelectPrivileges,omitempty" type:"Repeated"`
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

func (s ModifyAccountPrivilegesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequest) GetAccountName() *string {
	return s.AccountName
}

func (s *ModifyAccountPrivilegesRequest) GetAccountPrivileges() []*ModifyAccountPrivilegesRequestAccountPrivileges {
	return s.AccountPrivileges
}

func (s *ModifyAccountPrivilegesRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyAccountPrivilegesRequest) GetPromqlInsertPrivileges() []*string {
	return s.PromqlInsertPrivileges
}

func (s *ModifyAccountPrivilegesRequest) GetPromqlSelectNodePercentage() *float64 {
	return s.PromqlSelectNodePercentage
}

func (s *ModifyAccountPrivilegesRequest) GetPromqlSelectPrivileges() []*string {
	return s.PromqlSelectPrivileges
}

func (s *ModifyAccountPrivilegesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyAccountPrivilegesRequest) GetResourceGroupName() *string {
	return s.ResourceGroupName
}

func (s *ModifyAccountPrivilegesRequest) SetAccountName(v string) *ModifyAccountPrivilegesRequest {
	s.AccountName = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetAccountPrivileges(v []*ModifyAccountPrivilegesRequestAccountPrivileges) *ModifyAccountPrivilegesRequest {
	s.AccountPrivileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetDBClusterId(v string) *ModifyAccountPrivilegesRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetPromqlInsertPrivileges(v []*string) *ModifyAccountPrivilegesRequest {
	s.PromqlInsertPrivileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetPromqlSelectNodePercentage(v float64) *ModifyAccountPrivilegesRequest {
	s.PromqlSelectNodePercentage = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetPromqlSelectPrivileges(v []*string) *ModifyAccountPrivilegesRequest {
	s.PromqlSelectPrivileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetRegionId(v string) *ModifyAccountPrivilegesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) SetResourceGroupName(v string) *ModifyAccountPrivilegesRequest {
	s.ResourceGroupName = &v
	return s
}

func (s *ModifyAccountPrivilegesRequest) Validate() error {
	if s.AccountPrivileges != nil {
		for _, item := range s.AccountPrivileges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyAccountPrivilegesRequestAccountPrivileges struct {
	// The privilege object, which is a tuple of database, table, and column.
	PrivilegeObject *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The privilege level, obtained from the `DescribeEnabledPrivileges` operation.
	//
	// example:
	//
	// Global
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The list of granted permissions.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s ModifyAccountPrivilegesRequestAccountPrivileges) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesRequestAccountPrivileges) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) GetPrivilegeObject() *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	return s.PrivilegeObject
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) GetPrivileges() []*string {
	return s.Privileges
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivilegeObject(v *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.PrivilegeObject = v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivilegeType(v string) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.PrivilegeType = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) SetPrivileges(v []*string) *ModifyAccountPrivilegesRequestAccountPrivileges {
	s.Privileges = v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivileges) Validate() error {
	if s.PrivilegeObject != nil {
		if err := s.PrivilegeObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject struct {
	// The column to which permissions are granted. This parameter is required when the privilege level is column.
	//
	// example:
	//
	// column1
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The database to which permissions are granted. This parameter is required when the privilege level is database, table, or column.
	//
	// example:
	//
	// tsdb1
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The table to which permissions are granted. This parameter is required when the privilege level is table or column.
	//
	// example:
	//
	// table1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GoString() string {
	return s.String()
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GetColumn() *string {
	return s.Column
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GetDatabase() *string {
	return s.Database
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) GetTable() *string {
	return s.Table
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetColumn(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Column = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetDatabase(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Database = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) SetTable(v string) *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject {
	s.Table = &v
	return s
}

func (s *ModifyAccountPrivilegesRequestAccountPrivilegesPrivilegeObject) Validate() error {
	return dara.Validate(s)
}

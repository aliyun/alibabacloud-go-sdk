// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSQLServerUpgradeVersionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v *DescribeSQLServerUpgradeVersionsResponseBodyItems) *DescribeSQLServerUpgradeVersionsResponseBody
	GetItems() *DescribeSQLServerUpgradeVersionsResponseBodyItems
	SetRequestId(v string) *DescribeSQLServerUpgradeVersionsResponseBody
	GetRequestId() *string
}

type DescribeSQLServerUpgradeVersionsResponseBody struct {
	Items *DescribeSQLServerUpgradeVersionsResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// example:
	//
	// 866F5EB8-4650-4061-87F0-379F6F******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBody) GetItems() *DescribeSQLServerUpgradeVersionsResponseBodyItems {
	return s.Items
}

func (s *DescribeSQLServerUpgradeVersionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSQLServerUpgradeVersionsResponseBody) SetItems(v *DescribeSQLServerUpgradeVersionsResponseBodyItems) *DescribeSQLServerUpgradeVersionsResponseBody {
	s.Items = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBody) SetRequestId(v string) *DescribeSQLServerUpgradeVersionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLServerUpgradeVersionsResponseBodyItems struct {
	Item []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItem `json:"Item,omitempty" xml:"Item,omitempty" type:"Repeated"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItems) GetItem() []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItem {
	return s.Item
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItems) SetItem(v []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) *DescribeSQLServerUpgradeVersionsResponseBodyItems {
	s.Item = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItems) Validate() error {
	if s.Item != nil {
		for _, item := range s.Item {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLServerUpgradeVersionsResponseBodyItemsItem struct {
	// example:
	//
	// 2016_web
	CurrentVersion           *string                                                                        `json:"CurrentVersion,omitempty" xml:"CurrentVersion,omitempty"`
	SQLServerUpgradeVersions *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions `json:"SQLServerUpgradeVersions,omitempty" xml:"SQLServerUpgradeVersions,omitempty" type:"Struct"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) GetCurrentVersion() *string {
	return s.CurrentVersion
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) GetSQLServerUpgradeVersions() *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions {
	return s.SQLServerUpgradeVersions
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) SetCurrentVersion(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem {
	s.CurrentVersion = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) SetSQLServerUpgradeVersions(v *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem {
	s.SQLServerUpgradeVersions = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItem) Validate() error {
	if s.SQLServerUpgradeVersions != nil {
		if err := s.SQLServerUpgradeVersions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions struct {
	SQLServerUpgradeVersion []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion `json:"SQLServerUpgradeVersion,omitempty" xml:"SQLServerUpgradeVersion,omitempty" type:"Repeated"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions) GetSQLServerUpgradeVersion() []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion {
	return s.SQLServerUpgradeVersion
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions) SetSQLServerUpgradeVersion(v []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions {
	s.SQLServerUpgradeVersion = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersions) Validate() error {
	if s.SQLServerUpgradeVersion != nil {
		for _, item := range s.SQLServerUpgradeVersion {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion struct {
	DBInstanceClassItems *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems `json:"DBInstanceClassItems,omitempty" xml:"DBInstanceClassItems,omitempty" type:"Struct"`
	// example:
	//
	// NO/YES
	EnableUpgrade *string `json:"EnableUpgrade,omitempty" xml:"EnableUpgrade,omitempty"`
	// example:
	//
	// 2016_std
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) GetDBInstanceClassItems() *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems {
	return s.DBInstanceClassItems
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) GetEnableUpgrade() *string {
	return s.EnableUpgrade
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) GetVersion() *string {
	return s.Version
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) SetDBInstanceClassItems(v *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion {
	s.DBInstanceClassItems = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) SetEnableUpgrade(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion {
	s.EnableUpgrade = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) SetVersion(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion {
	s.Version = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersion) Validate() error {
	if s.DBInstanceClassItems != nil {
		if err := s.DBInstanceClassItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems struct {
	DBInstanceClassItem []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem `json:"DBInstanceClassItem,omitempty" xml:"DBInstanceClassItem,omitempty" type:"Repeated"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems) GetDBInstanceClassItem() []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem {
	return s.DBInstanceClassItem
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems) SetDBInstanceClassItem(v []*DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems {
	s.DBInstanceClassItem = v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItems) Validate() error {
	if s.DBInstanceClassItem != nil {
		for _, item := range s.DBInstanceClassItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem struct {
	// example:
	//
	// 2
	CPU *string `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// example:
	//
	// mssql.x4.medium.s2
	DBInstanceClass     *string `json:"DBInstanceClass,omitempty" xml:"DBInstanceClass,omitempty"`
	DBInstanceClassType *string `json:"DBInstanceClassType,omitempty" xml:"DBInstanceClassType,omitempty"`
	// example:
	//
	// 2
	Group *string `json:"Group,omitempty" xml:"Group,omitempty"`
	// example:
	//
	// 8GB
	Memory *string `json:"Memory,omitempty" xml:"Memory,omitempty"`
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) GoString() string {
	return s.String()
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) GetCPU() *string {
	return s.CPU
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) GetDBInstanceClass() *string {
	return s.DBInstanceClass
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) GetDBInstanceClassType() *string {
	return s.DBInstanceClassType
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) GetGroup() *string {
	return s.Group
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) GetMemory() *string {
	return s.Memory
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) SetCPU(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem {
	s.CPU = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) SetDBInstanceClass(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem {
	s.DBInstanceClass = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) SetDBInstanceClassType(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem {
	s.DBInstanceClassType = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) SetGroup(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem {
	s.Group = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) SetMemory(v string) *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem {
	s.Memory = &v
	return s
}

func (s *DescribeSQLServerUpgradeVersionsResponseBodyItemsItemSQLServerUpgradeVersionsSQLServerUpgradeVersionDBInstanceClassItemsDBInstanceClassItem) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataBloatResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBInstanceDataBloatResponseBodyItems) *DescribeDBInstanceDataBloatResponseBody
	GetItems() []*DescribeDBInstanceDataBloatResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstanceDataBloatResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeDBInstanceDataBloatResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBInstanceDataBloatResponseBody
	GetTotalCount() *int32
}

type DescribeDBInstanceDataBloatResponseBody struct {
	// The queried data bloat.
	Items []*DescribeDBInstanceDataBloatResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponseBody) GetItems() []*DescribeDBInstanceDataBloatResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceDataBloatResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceDataBloatResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceDataBloatResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetItems(v []*DescribeDBInstanceDataBloatResponseBodyItems) *DescribeDBInstanceDataBloatResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetPageNumber(v int32) *DescribeDBInstanceDataBloatResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetRequestId(v string) *DescribeDBInstanceDataBloatResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) SetTotalCount(v int32) *DescribeDBInstanceDataBloatResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDBInstanceDataBloatResponseBodyItems struct {
	// The coefficient of data bloat. It is calculated by using the following formula:
	//
	// Bloat coefficient = Number of dead rows/Number of active rows.
	//
	// example:
	//
	// 1.03
	BloatCeoff *string `json:"BloatCeoff,omitempty" xml:"BloatCeoff,omitempty"`
	// The bloat size of the table. It indicates the amount of space that can be released.
	//
	// example:
	//
	// 0.2MB
	BloatSize *string `json:"BloatSize,omitempty" xml:"BloatSize,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// adbtest
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The expected size of the table.
	//
	// It indicates the size of the table that has no data bloat.
	//
	// example:
	//
	// 1MB
	ExpectTableSize *string `json:"ExpectTableSize,omitempty" xml:"ExpectTableSize,omitempty"`
	// The actual size of the table.
	//
	// example:
	//
	// 1.2MB
	RealTableSize *string `json:"RealTableSize,omitempty" xml:"RealTableSize,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema1
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sequence number.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The storage type of the table. Valid values:
	//
	// 	- **Heap Table**
	//
	// 	- **Append-Only Heap Table**
	//
	// 	- **Append-Only Columnar Table**
	//
	// example:
	//
	// Heap Table
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
	// This parameter is not returned.
	//
	// example:
	//
	// null
	SuggestedAction *string `json:"SuggestedAction,omitempty" xml:"SuggestedAction,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tab1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The time when the table was last deleted, inserted, or updated.
	//
	// example:
	//
	// 2022-08-08T20:00:00Z
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
	// The time when the table was last vacuumed. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-08-08T16:00:00Z
	TimeLastVacuumed *string `json:"TimeLastVacuumed,omitempty" xml:"TimeLastVacuumed,omitempty"`
}

func (s DescribeDBInstanceDataBloatResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataBloatResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetBloatCeoff() *string {
	return s.BloatCeoff
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetBloatSize() *string {
	return s.BloatSize
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetExpectTableSize() *string {
	return s.ExpectTableSize
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetRealTableSize() *string {
	return s.RealTableSize
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetSequence() *int32 {
	return s.Sequence
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetStorageType() *string {
	return s.StorageType
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetSuggestedAction() *string {
	return s.SuggestedAction
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetTimeLastUpdated() *string {
	return s.TimeLastUpdated
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) GetTimeLastVacuumed() *string {
	return s.TimeLastVacuumed
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetBloatCeoff(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.BloatCeoff = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetBloatSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.BloatSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetExpectTableSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.ExpectTableSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetRealTableSize(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.RealTableSize = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSequence(v int32) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.Sequence = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetStorageType(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.StorageType = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetSuggestedAction(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.SuggestedAction = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTableName(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) SetTimeLastVacuumed(v string) *DescribeDBInstanceDataBloatResponseBodyItems {
	s.TimeLastVacuumed = &v
	return s
}

func (s *DescribeDBInstanceDataBloatResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

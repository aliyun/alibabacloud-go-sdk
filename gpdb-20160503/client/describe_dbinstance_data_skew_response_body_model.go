// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceDataSkewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBInstanceDataSkewResponseBodyItems) *DescribeDBInstanceDataSkewResponseBody
	GetItems() []*DescribeDBInstanceDataSkewResponseBodyItems
	SetPageNumber(v int32) *DescribeDBInstanceDataSkewResponseBody
	GetPageNumber() *int32
	SetRequestId(v string) *DescribeDBInstanceDataSkewResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDBInstanceDataSkewResponseBody
	GetTotalCount() *int32
}

type DescribeDBInstanceDataSkewResponseBody struct {
	// Details about data skew.
	Items []*DescribeDBInstanceDataSkewResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponseBody) GetItems() []*DescribeDBInstanceDataSkewResponseBodyItems {
	return s.Items
}

func (s *DescribeDBInstanceDataSkewResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDBInstanceDataSkewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceDataSkewResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetItems(v []*DescribeDBInstanceDataSkewResponseBodyItems) *DescribeDBInstanceDataSkewResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetPageNumber(v int32) *DescribeDBInstanceDataSkewResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetRequestId(v string) *DescribeDBInstanceDataSkewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) SetTotalCount(v int32) *DescribeDBInstanceDataSkewResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBInstanceDataSkewResponseBodyItems struct {
	// The name of the database.
	//
	// example:
	//
	// adbtest
	DatabaseName *string `json:"DatabaseName,omitempty" xml:"DatabaseName,omitempty"`
	// The distribution key of the table.
	//
	// example:
	//
	// ItemId
	DistributeKey *string `json:"DistributeKey,omitempty" xml:"DistributeKey,omitempty"`
	// The owner of the table.
	//
	// example:
	//
	// adbpguser
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The name of the schema.
	//
	// example:
	//
	// schema1
	SchemaName *string `json:"SchemaName,omitempty" xml:"SchemaName,omitempty"`
	// The sequence number of the data skew case. All data skew cases are sorted by severity in descending order.
	//
	// example:
	//
	// 1
	Sequence *int32 `json:"Sequence,omitempty" xml:"Sequence,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// tab1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// The total number of rows in the table.
	//
	// example:
	//
	// 100000
	TableSize *string `json:"TableSize,omitempty" xml:"TableSize,omitempty"`
	// The skew ratio of the table. Valid values: 0 to 100. Unit: %. A greater value indicates that the table is more severely skewed. A smaller value indicates less impact on query performance. A value of 0 indicates no data skew.
	//
	// example:
	//
	// 10.23
	TableSkew *string `json:"TableSkew,omitempty" xml:"TableSkew,omitempty"`
	// The time when the table was last deleted, inserted, or updated.
	//
	// example:
	//
	// 2020-09-08T20:00:00Z
	TimeLastUpdated *string `json:"TimeLastUpdated,omitempty" xml:"TimeLastUpdated,omitempty"`
}

func (s DescribeDBInstanceDataSkewResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceDataSkewResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetDatabaseName() *string {
	return s.DatabaseName
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetDistributeKey() *string {
	return s.DistributeKey
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetOwner() *string {
	return s.Owner
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetSchemaName() *string {
	return s.SchemaName
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetSequence() *int32 {
	return s.Sequence
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetTableSize() *string {
	return s.TableSize
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetTableSkew() *string {
	return s.TableSkew
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) GetTimeLastUpdated() *string {
	return s.TimeLastUpdated
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetDatabaseName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.DatabaseName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetDistributeKey(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.DistributeKey = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetOwner(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.Owner = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetSchemaName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.SchemaName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetSequence(v int32) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.Sequence = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableName(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableName = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableSize(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableSize = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTableSkew(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TableSkew = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) SetTimeLastUpdated(v string) *DescribeDBInstanceDataSkewResponseBodyItems {
	s.TimeLastUpdated = &v
	return s
}

func (s *DescribeDBInstanceDataSkewResponseBodyItems) Validate() error {
	return dara.Validate(s)
}

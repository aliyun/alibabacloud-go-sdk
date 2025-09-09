// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeTableResponseBodyData) *DescribeTableResponseBody
	GetData() *DescribeTableResponseBodyData
	SetRequestId(v string) *DescribeTableResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTableResponseBody
	GetSuccess() *bool
}

type DescribeTableResponseBody struct {
	// Indicates the returned data.
	Data *DescribeTableResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Indicates the unique ID of the request. If the request fails, provide this ID for technical support to troubleshoot the failure.
	//
	// example:
	//
	// B5644ABB-559A-4A1C-83F2-9E7209******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBody) GetData() *DescribeTableResponseBodyData {
	return s.Data
}

func (s *DescribeTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTableResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTableResponseBody) SetData(v *DescribeTableResponseBodyData) *DescribeTableResponseBody {
	s.Data = v
	return s
}

func (s *DescribeTableResponseBody) SetRequestId(v string) *DescribeTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTableResponseBody) SetSuccess(v bool) *DescribeTableResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTableResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTableResponseBodyData struct {
	// Indicates the details about the table schema.
	List []*DescribeTableResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s DescribeTableResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyData) GetList() []*DescribeTableResponseBodyDataList {
	return s.List
}

func (s *DescribeTableResponseBodyData) SetList(v []*DescribeTableResponseBodyDataList) *DescribeTableResponseBodyData {
	s.List = v
	return s
}

func (s *DescribeTableResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeTableResponseBodyDataList struct {
	// Indicates the name of a column.
	//
	// example:
	//
	// Id
	ColumnName *string `json:"ColumnName,omitempty" xml:"ColumnName,omitempty"`
	// Indicates the type of the column.
	//
	// example:
	//
	// bigint(20)
	ColumnType *string `json:"ColumnType,omitempty" xml:"ColumnType,omitempty"`
	// Extra
	//
	// example:
	//
	// auto_increment
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// Indicates the primary key of the table.
	//
	// example:
	//
	// PRI
	Index *string `json:"Index,omitempty" xml:"Index,omitempty"`
	// Indicates whether the column can be empty.
	//
	// example:
	//
	// NO
	IsAllowNull *string `json:"IsAllowNull,omitempty" xml:"IsAllowNull,omitempty"`
	// Indicates whether the column is the primary key column of the table.
	//
	// example:
	//
	// YES
	IsPk *string `json:"IsPk,omitempty" xml:"IsPk,omitempty"`
}

func (s DescribeTableResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s DescribeTableResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *DescribeTableResponseBodyDataList) GetColumnName() *string {
	return s.ColumnName
}

func (s *DescribeTableResponseBodyDataList) GetColumnType() *string {
	return s.ColumnType
}

func (s *DescribeTableResponseBodyDataList) GetExtra() *string {
	return s.Extra
}

func (s *DescribeTableResponseBodyDataList) GetIndex() *string {
	return s.Index
}

func (s *DescribeTableResponseBodyDataList) GetIsAllowNull() *string {
	return s.IsAllowNull
}

func (s *DescribeTableResponseBodyDataList) GetIsPk() *string {
	return s.IsPk
}

func (s *DescribeTableResponseBodyDataList) SetColumnName(v string) *DescribeTableResponseBodyDataList {
	s.ColumnName = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetColumnType(v string) *DescribeTableResponseBodyDataList {
	s.ColumnType = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetExtra(v string) *DescribeTableResponseBodyDataList {
	s.Extra = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIndex(v string) *DescribeTableResponseBodyDataList {
	s.Index = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsAllowNull(v string) *DescribeTableResponseBodyDataList {
	s.IsAllowNull = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) SetIsPk(v string) *DescribeTableResponseBodyDataList {
	s.IsPk = &v
	return s
}

func (s *DescribeTableResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}

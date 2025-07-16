// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDistributeTableListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDistributeTableListResponseBodyData) *DescribeDistributeTableListResponseBody
	GetData() *DescribeDistributeTableListResponseBodyData
	SetMessage(v string) *DescribeDistributeTableListResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDistributeTableListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeDistributeTableListResponseBody
	GetSuccess() *bool
}

type DescribeDistributeTableListResponseBody struct {
	Data *DescribeDistributeTableListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDistributeTableListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributeTableListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBody) GetData() *DescribeDistributeTableListResponseBodyData {
	return s.Data
}

func (s *DescribeDistributeTableListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDistributeTableListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDistributeTableListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeDistributeTableListResponseBody) SetData(v *DescribeDistributeTableListResponseBodyData) *DescribeDistributeTableListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDistributeTableListResponseBody) SetMessage(v string) *DescribeDistributeTableListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDistributeTableListResponseBody) SetRequestId(v string) *DescribeDistributeTableListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDistributeTableListResponseBody) SetSuccess(v bool) *DescribeDistributeTableListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDistributeTableListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeDistributeTableListResponseBodyData struct {
	Tables []*DescribeDistributeTableListResponseBodyDataTables `json:"Tables,omitempty" xml:"Tables,omitempty" type:"Repeated"`
}

func (s DescribeDistributeTableListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributeTableListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBodyData) GetTables() []*DescribeDistributeTableListResponseBodyDataTables {
	return s.Tables
}

func (s *DescribeDistributeTableListResponseBodyData) SetTables(v []*DescribeDistributeTableListResponseBodyDataTables) *DescribeDistributeTableListResponseBodyData {
	s.Tables = v
	return s
}

func (s *DescribeDistributeTableListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeDistributeTableListResponseBodyDataTables struct {
	// example:
	//
	// id
	DbKey *string `json:"DbKey,omitempty" xml:"DbKey,omitempty"`
	// example:
	//
	// sbtest1
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
	// example:
	//
	// multi
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
	// example:
	//
	// “”
	TbKey *string `json:"TbKey,omitempty" xml:"TbKey,omitempty"`
}

func (s DescribeDistributeTableListResponseBodyDataTables) String() string {
	return dara.Prettify(s)
}

func (s DescribeDistributeTableListResponseBodyDataTables) GoString() string {
	return s.String()
}

func (s *DescribeDistributeTableListResponseBodyDataTables) GetDbKey() *string {
	return s.DbKey
}

func (s *DescribeDistributeTableListResponseBodyDataTables) GetTableName() *string {
	return s.TableName
}

func (s *DescribeDistributeTableListResponseBodyDataTables) GetTableType() *string {
	return s.TableType
}

func (s *DescribeDistributeTableListResponseBodyDataTables) GetTbKey() *string {
	return s.TbKey
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetDbKey(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.DbKey = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetTableName(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.TableName = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetTableType(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.TableType = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) SetTbKey(v string) *DescribeDistributeTableListResponseBodyDataTables {
	s.TbKey = &v
	return s
}

func (s *DescribeDistributeTableListResponseBodyDataTables) Validate() error {
	return dara.Validate(s)
}

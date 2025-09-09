// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeExpandLogicTableInfoListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeExpandLogicTableInfoListResponseBodyData) *DescribeExpandLogicTableInfoListResponseBody
	GetData() *DescribeExpandLogicTableInfoListResponseBodyData
	SetRequestId(v string) *DescribeExpandLogicTableInfoListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeExpandLogicTableInfoListResponseBody
	GetSuccess() *bool
}

type DescribeExpandLogicTableInfoListResponseBody struct {
	// Indicates the result that is returned.
	Data *DescribeExpandLogicTableInfoListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 0B6B7BDC-575D-4A77-A4F8-24B7EF******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBody) GetData() *DescribeExpandLogicTableInfoListResponseBodyData {
	return s.Data
}

func (s *DescribeExpandLogicTableInfoListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeExpandLogicTableInfoListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetData(v *DescribeExpandLogicTableInfoListResponseBodyData) *DescribeExpandLogicTableInfoListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetRequestId(v string) *DescribeExpandLogicTableInfoListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) SetSuccess(v bool) *DescribeExpandLogicTableInfoListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeExpandLogicTableInfoListResponseBodyData struct {
	Data []*DescribeExpandLogicTableInfoListResponseBodyDataData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyData) GetData() []*DescribeExpandLogicTableInfoListResponseBodyDataData {
	return s.Data
}

func (s *DescribeExpandLogicTableInfoListResponseBodyData) SetData(v []*DescribeExpandLogicTableInfoListResponseBodyDataData) *DescribeExpandLogicTableInfoListResponseBodyData {
	s.Data = v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type DescribeExpandLogicTableInfoListResponseBodyDataData struct {
	// Indicates the database sharding key.
	//
	// example:
	//
	// id
	ShardDbKey *string `json:"ShardDbKey,omitempty" xml:"ShardDbKey,omitempty"`
	// Indicates the table sharding key.
	//
	// example:
	//
	// address
	ShardTbKey *string `json:"ShardTbKey,omitempty" xml:"ShardTbKey,omitempty"`
	// Indicates the name of the table.
	//
	// example:
	//
	// employee_split2
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) String() string {
	return dara.Prettify(s)
}

func (s DescribeExpandLogicTableInfoListResponseBodyDataData) GoString() string {
	return s.String()
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) GetShardDbKey() *string {
	return s.ShardDbKey
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) GetShardTbKey() *string {
	return s.ShardTbKey
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) GetTableName() *string {
	return s.TableName
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardDbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardDbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetShardTbKey(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.ShardTbKey = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) SetTableName(v string) *DescribeExpandLogicTableInfoListResponseBodyDataData {
	s.TableName = &v
	return s
}

func (s *DescribeExpandLogicTableInfoListResponseBodyDataData) Validate() error {
	return dara.Validate(s)
}

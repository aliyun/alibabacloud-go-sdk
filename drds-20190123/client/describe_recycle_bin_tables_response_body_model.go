// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecycleBinTablesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribeRecycleBinTablesResponseBodyData) *DescribeRecycleBinTablesResponseBody
	GetData() []*DescribeRecycleBinTablesResponseBodyData
	SetRequestId(v string) *DescribeRecycleBinTablesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeRecycleBinTablesResponseBody
	GetSuccess() *bool
}

type DescribeRecycleBinTablesResponseBody struct {
	// The data object returned.
	Data []*DescribeRecycleBinTablesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 5D64DE5944A1E541E0CB908A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result of the request.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeRecycleBinTablesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinTablesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponseBody) GetData() []*DescribeRecycleBinTablesResponseBodyData {
	return s.Data
}

func (s *DescribeRecycleBinTablesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecycleBinTablesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeRecycleBinTablesResponseBody) SetData(v []*DescribeRecycleBinTablesResponseBodyData) *DescribeRecycleBinTablesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) SetRequestId(v string) *DescribeRecycleBinTablesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) SetSuccess(v bool) *DescribeRecycleBinTablesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRecycleBinTablesResponseBodyData struct {
	// The time when the table was created.
	//
	// example:
	//
	// 2019-09-16 14:42:06
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The original name of the table.
	//
	// example:
	//
	// BIN_T4AG3CY5WWXPKHITCHJY
	OriginalTableName *string `json:"OriginalTableName,omitempty" xml:"OriginalTableName,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeRecycleBinTablesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecycleBinTablesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeRecycleBinTablesResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeRecycleBinTablesResponseBodyData) GetOriginalTableName() *string {
	return s.OriginalTableName
}

func (s *DescribeRecycleBinTablesResponseBodyData) GetTableName() *string {
	return s.TableName
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetCreateTime(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetOriginalTableName(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.OriginalTableName = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) SetTableName(v string) *DescribeRecycleBinTablesResponseBodyData {
	s.TableName = &v
	return s
}

func (s *DescribeRecycleBinTablesResponseBodyData) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAdbMySqlIndexesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIndexCount(v int32) *DescribeAdbMySqlIndexesResponseBody
	GetIndexCount() *int32
	SetIndexes(v []*DescribeAdbMySqlIndexesResponseBodyIndexes) *DescribeAdbMySqlIndexesResponseBody
	GetIndexes() []*DescribeAdbMySqlIndexesResponseBodyIndexes
	SetMessage(v string) *DescribeAdbMySqlIndexesResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeAdbMySqlIndexesResponseBody
	GetRequestId() *string
	SetSchema(v string) *DescribeAdbMySqlIndexesResponseBody
	GetSchema() *string
	SetSuccess(v bool) *DescribeAdbMySqlIndexesResponseBody
	GetSuccess() *bool
	SetTableName(v string) *DescribeAdbMySqlIndexesResponseBody
	GetTableName() *string
}

type DescribeAdbMySqlIndexesResponseBody struct {
	// The number of indexes.````
	//
	// example:
	//
	// 10
	IndexCount *int32 `json:"IndexCount,omitempty" xml:"IndexCount,omitempty"`
	// The queried indexes.
	Indexes []*DescribeAdbMySqlIndexesResponseBodyIndexes `json:"Indexes,omitempty" xml:"Indexes,omitempty" type:"Repeated"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F0983B43-B2EC-536A-8791-142B5CF1E9B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the database.
	//
	// **
	//
	// ****\\
	//
	// example:
	//
	// adb_demo
	Schema *string `json:"Schema,omitempty" xml:"Schema,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s DescribeAdbMySqlIndexesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlIndexesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetIndexCount() *int32 {
	return s.IndexCount
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetIndexes() []*DescribeAdbMySqlIndexesResponseBodyIndexes {
	return s.Indexes
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetSchema() *string {
	return s.Schema
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAdbMySqlIndexesResponseBody) GetTableName() *string {
	return s.TableName
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetIndexCount(v int32) *DescribeAdbMySqlIndexesResponseBody {
	s.IndexCount = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetIndexes(v []*DescribeAdbMySqlIndexesResponseBodyIndexes) *DescribeAdbMySqlIndexesResponseBody {
	s.Indexes = v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetMessage(v string) *DescribeAdbMySqlIndexesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetRequestId(v string) *DescribeAdbMySqlIndexesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetSchema(v string) *DescribeAdbMySqlIndexesResponseBody {
	s.Schema = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetSuccess(v bool) *DescribeAdbMySqlIndexesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) SetTableName(v string) *DescribeAdbMySqlIndexesResponseBody {
	s.TableName = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBody) Validate() error {
	if s.Indexes != nil {
		for _, item := range s.Indexes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAdbMySqlIndexesResponseBodyIndexes struct {
	// The name of the column.
	//
	// example:
	//
	// preclcu
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the index.
	//
	// example:
	//
	// ttl
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The index type.
	//
	// example:
	//
	// normal
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeAdbMySqlIndexesResponseBodyIndexes) String() string {
	return dara.Prettify(s)
}

func (s DescribeAdbMySqlIndexesResponseBodyIndexes) GoString() string {
	return s.String()
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) GetColumn() *string {
	return s.Column
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) GetName() *string {
	return s.Name
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) GetType() *string {
	return s.Type
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) SetColumn(v string) *DescribeAdbMySqlIndexesResponseBodyIndexes {
	s.Column = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) SetName(v string) *DescribeAdbMySqlIndexesResponseBodyIndexes {
	s.Name = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) SetType(v string) *DescribeAdbMySqlIndexesResponseBodyIndexes {
	s.Type = &v
	return s
}

func (s *DescribeAdbMySqlIndexesResponseBodyIndexes) Validate() error {
	return dara.Validate(s)
}

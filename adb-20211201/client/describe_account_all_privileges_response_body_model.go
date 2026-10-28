// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAllPrivilegesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeAccountAllPrivilegesResponseBodyData) *DescribeAccountAllPrivilegesResponseBody
	GetData() *DescribeAccountAllPrivilegesResponseBodyData
	SetRequestId(v string) *DescribeAccountAllPrivilegesResponseBody
	GetRequestId() *string
}

type DescribeAccountAllPrivilegesResponseBody struct {
	// The permission details.
	Data *DescribeAccountAllPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3BB185E9-BB54-1727-B876-13243E4C0EB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBody) GetData() *DescribeAccountAllPrivilegesResponseBodyData {
	return s.Data
}

func (s *DescribeAccountAllPrivilegesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountAllPrivilegesResponseBody) SetData(v *DescribeAccountAllPrivilegesResponseBodyData) *DescribeAccountAllPrivilegesResponseBody {
	s.Data = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBody) SetRequestId(v string) *DescribeAccountAllPrivilegesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountAllPrivilegesResponseBodyData struct {
	// If the `Truncated` field in the response is `true`, specify the value of this field in the next call to retrieve the subsequent entries.
	//
	// example:
	//
	// 0573e74fd1ccb01739993a691e876074db6e1b6ad79f54115f0e98528432ba6a523cfec5780ade5189299cc3396f6ff7
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The list of permissions.
	Result []*DescribeAccountAllPrivilegesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the returned content is truncated. If the value is `true`, the returned content is truncated, and you must continue to call this operation to retrieve the remaining entries until the value is `false`.
	//
	// example:
	//
	// true
	Truncated *bool `json:"Truncated,omitempty" xml:"Truncated,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) GetMarker() *string {
	return s.Marker
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) GetResult() []*DescribeAccountAllPrivilegesResponseBodyDataResult {
	return s.Result
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) GetTruncated() *bool {
	return s.Truncated
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) SetMarker(v string) *DescribeAccountAllPrivilegesResponseBodyData {
	s.Marker = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) SetResult(v []*DescribeAccountAllPrivilegesResponseBodyDataResult) *DescribeAccountAllPrivilegesResponseBodyData {
	s.Result = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) SetTruncated(v bool) *DescribeAccountAllPrivilegesResponseBodyData {
	s.Truncated = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyData) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountAllPrivilegesResponseBodyDataResult struct {
	// The permission object, which is a triplet of the database, table, and column. All fields are empty when the permission level is Global.
	PrivilegeObject *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level, which is obtained from the `DescribeEnabledPrivileges` operation.
	//
	// example:
	//
	// Global
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The list of permissions.
	Privileges []*string `json:"Privileges,omitempty" xml:"Privileges,omitempty" type:"Repeated"`
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResult) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) GetPrivilegeObject() *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	return s.PrivilegeObject
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) GetPrivilegeType() *string {
	return s.PrivilegeType
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) GetPrivileges() []*string {
	return s.Privileges
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) SetPrivilegeObject(v *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) *DescribeAccountAllPrivilegesResponseBodyDataResult {
	s.PrivilegeObject = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) SetPrivilegeType(v string) *DescribeAccountAllPrivilegesResponseBodyDataResult {
	s.PrivilegeType = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) SetPrivileges(v []*string) *DescribeAccountAllPrivilegesResponseBodyDataResult {
	s.Privileges = v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResult) Validate() error {
	if s.PrivilegeObject != nil {
		if err := s.PrivilegeObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject struct {
	// The name of the column.
	//
	// example:
	//
	// id
	Column *string `json:"Column,omitempty" xml:"Column,omitempty"`
	// The name of the database.
	//
	// example:
	//
	// tdb1
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The description of the permission object.
	//
	// example:
	//
	// id of table
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the table.
	//
	// example:
	//
	// table1
	Table *string `json:"Table,omitempty" xml:"Table,omitempty"`
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) GoString() string {
	return s.String()
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) GetColumn() *string {
	return s.Column
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) GetDatabase() *string {
	return s.Database
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) GetTable() *string {
	return s.Table
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetColumn(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Column = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetDatabase(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Database = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetDescription(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Description = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) SetTable(v string) *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject {
	s.Table = &v
	return s
}

func (s *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject) Validate() error {
	return dara.Validate(s)
}

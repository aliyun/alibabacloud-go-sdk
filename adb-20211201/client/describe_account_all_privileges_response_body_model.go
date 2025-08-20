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
	// Details of the permissions.
	Data *DescribeAccountAllPrivilegesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The ID of the request.
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
	return dara.Validate(s)
}

type DescribeAccountAllPrivilegesResponseBodyData struct {
	// Indicates the position where the results are truncated. When a value of `true` is returned for the `Truncated` parameter, this parameter is present and contains the value to use for the Marker parameter in a subsequent call.
	//
	// example:
	//
	// 0573e74fd1ccb01739993a691e876074db6e1b6ad79f54115f0e98528432ba6a523cfec5780ade5189299cc3396f6ff7
	Marker *string `json:"Marker,omitempty" xml:"Marker,omitempty"`
	// The permissions.
	Result []*DescribeAccountAllPrivilegesResponseBodyDataResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the results are truncated. If the results are truncated, a value of `true` is returned. In this case, you must call this operation again to obtain all the results until a value of `false` is returned for this parameter.
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
	return dara.Validate(s)
}

type DescribeAccountAllPrivilegesResponseBodyDataResult struct {
	// The objects on which the permission takes effect, including databases, tables, and columns. If Global is returned for the PrivilegeType parameter, an empty string is returned for this parameter.
	PrivilegeObject *DescribeAccountAllPrivilegesResponseBodyDataResultPrivilegeObject `json:"PrivilegeObject,omitempty" xml:"PrivilegeObject,omitempty" type:"Struct"`
	// The permission level of the database account. You can call the `DescribeEnabledPrivileges` operation to query the permission level of the database account.
	//
	// example:
	//
	// Global
	PrivilegeType *string `json:"PrivilegeType,omitempty" xml:"PrivilegeType,omitempty"`
	// The name of the permission, which is the same as the permission name returned by the `DescribeEnabledPrivileges` operation.
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
	return dara.Validate(s)
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

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLakeStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *UpdateLakeStorageResponseBody
	GetCode() *string
	SetData(v *UpdateLakeStorageResponseBodyData) *UpdateLakeStorageResponseBody
	GetData() *UpdateLakeStorageResponseBodyData
	SetHttpStatusCode(v int32) *UpdateLakeStorageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *UpdateLakeStorageResponseBody
	GetMessage() *string
	SetRequestId(v string) *UpdateLakeStorageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateLakeStorageResponseBody
	GetSuccess() *bool
}

type UpdateLakeStorageResponseBody struct {
	// The status code. A value of 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	//
	// example:
	//
	// -
	Data *UpdateLakeStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code. A value of 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, an OK message is returned.
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******-3EEC-******-9F06-******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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
}

func (s UpdateLakeStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageResponseBody) GetCode() *string {
	return s.Code
}

func (s *UpdateLakeStorageResponseBody) GetData() *UpdateLakeStorageResponseBodyData {
	return s.Data
}

func (s *UpdateLakeStorageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *UpdateLakeStorageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *UpdateLakeStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateLakeStorageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateLakeStorageResponseBody) SetCode(v string) *UpdateLakeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *UpdateLakeStorageResponseBody) SetData(v *UpdateLakeStorageResponseBodyData) *UpdateLakeStorageResponseBody {
	s.Data = v
	return s
}

func (s *UpdateLakeStorageResponseBody) SetHttpStatusCode(v int32) *UpdateLakeStorageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *UpdateLakeStorageResponseBody) SetMessage(v string) *UpdateLakeStorageResponseBody {
	s.Message = &v
	return s
}

func (s *UpdateLakeStorageResponseBody) SetRequestId(v string) *UpdateLakeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateLakeStorageResponseBody) SetSuccess(v bool) *UpdateLakeStorageResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateLakeStorageResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpdateLakeStorageResponseBodyData struct {
	// The time when the lake storage was created.
	//
	// example:
	//
	// 2023-05-15T07:24:58Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The creator UID.
	//
	// example:
	//
	// 123456
	CreatorUid *string `json:"CreatorUid,omitempty" xml:"CreatorUid,omitempty"`
	// The ID of the AnalyticDB for MySQL cluster that is associated with the lake storage.
	//
	// example:
	//
	// amv-23xxxx
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the lake storage.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total storage size.
	//
	// example:
	//
	// 142
	FileSize *string `json:"FileSize,omitempty" xml:"FileSize,omitempty"`
	// The unique identifier of the lake storage.
	//
	// example:
	//
	// -
	LakeStorageId *string `json:"LakeStorageId,omitempty" xml:"LakeStorageId,omitempty"`
	// The operator UID.
	//
	// example:
	//
	// 123456
	OperatorUid *string `json:"OperatorUid,omitempty" xml:"OperatorUid,omitempty"`
	// The owner UID.
	//
	// example:
	//
	// 123456
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The number of partitions.
	//
	// example:
	//
	// 0
	PartitionCount *string `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The permissions on the lake storage.
	//
	// example:
	//
	// -
	Permissions []*UpdateLakeStorageResponseBodyDataPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// >=
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 1234
	TableCount *int32 `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	// The time when the lake storage was last updated.
	//
	// example:
	//
	// 2024-07-01T09:22:17Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s UpdateLakeStorageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *UpdateLakeStorageResponseBodyData) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *UpdateLakeStorageResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *UpdateLakeStorageResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *UpdateLakeStorageResponseBodyData) GetFileSize() *string {
	return s.FileSize
}

func (s *UpdateLakeStorageResponseBodyData) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *UpdateLakeStorageResponseBodyData) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *UpdateLakeStorageResponseBodyData) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *UpdateLakeStorageResponseBodyData) GetPartitionCount() *string {
	return s.PartitionCount
}

func (s *UpdateLakeStorageResponseBodyData) GetPermissions() []*UpdateLakeStorageResponseBodyDataPermissions {
	return s.Permissions
}

func (s *UpdateLakeStorageResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *UpdateLakeStorageResponseBodyData) GetRowCount() *int64 {
	return s.RowCount
}

func (s *UpdateLakeStorageResponseBodyData) GetTableCount() *int32 {
	return s.TableCount
}

func (s *UpdateLakeStorageResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *UpdateLakeStorageResponseBodyData) SetCreateTime(v string) *UpdateLakeStorageResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetCreatorUid(v string) *UpdateLakeStorageResponseBodyData {
	s.CreatorUid = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetDBClusterId(v string) *UpdateLakeStorageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetDescription(v string) *UpdateLakeStorageResponseBodyData {
	s.Description = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetFileSize(v string) *UpdateLakeStorageResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetLakeStorageId(v string) *UpdateLakeStorageResponseBodyData {
	s.LakeStorageId = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetOperatorUid(v string) *UpdateLakeStorageResponseBodyData {
	s.OperatorUid = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetOwnerUid(v string) *UpdateLakeStorageResponseBodyData {
	s.OwnerUid = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetPartitionCount(v string) *UpdateLakeStorageResponseBodyData {
	s.PartitionCount = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetPermissions(v []*UpdateLakeStorageResponseBodyDataPermissions) *UpdateLakeStorageResponseBodyData {
	s.Permissions = v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetRegionId(v string) *UpdateLakeStorageResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetRowCount(v int64) *UpdateLakeStorageResponseBodyData {
	s.RowCount = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetTableCount(v int32) *UpdateLakeStorageResponseBodyData {
	s.TableCount = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) SetUpdateTime(v string) *UpdateLakeStorageResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyData) Validate() error {
	return dara.Validate(s)
}

type UpdateLakeStorageResponseBodyDataPermissions struct {
	// The account ID.
	//
	// example:
	//
	// test
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The read permissions.
	//
	// example:
	//
	// true
	Read *bool `json:"Read,omitempty" xml:"Read,omitempty"`
	// The account type.
	//
	// example:
	//
	// SUB
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write permissions.
	//
	// example:
	//
	// false
	Write *bool `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s UpdateLakeStorageResponseBodyDataPermissions) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageResponseBodyDataPermissions) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) GetAccount() *string {
	return s.Account
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) GetRead() *bool {
	return s.Read
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) GetType() *string {
	return s.Type
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) GetWrite() *bool {
	return s.Write
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) SetAccount(v string) *UpdateLakeStorageResponseBodyDataPermissions {
	s.Account = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) SetRead(v bool) *UpdateLakeStorageResponseBodyDataPermissions {
	s.Read = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) SetType(v string) *UpdateLakeStorageResponseBodyDataPermissions {
	s.Type = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) SetWrite(v bool) *UpdateLakeStorageResponseBodyDataPermissions {
	s.Write = &v
	return s
}

func (s *UpdateLakeStorageResponseBodyDataPermissions) Validate() error {
	return dara.Validate(s)
}

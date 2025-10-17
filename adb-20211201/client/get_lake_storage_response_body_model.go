// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLakeStorageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetLakeStorageResponseBody
	GetCode() *string
	SetData(v *GetLakeStorageResponseBodyData) *GetLakeStorageResponseBody
	GetData() *GetLakeStorageResponseBodyData
	SetHttpStatusCode(v int32) *GetLakeStorageResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetLakeStorageResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetLakeStorageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetLakeStorageResponseBody
	GetSuccess() *bool
}

type GetLakeStorageResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The queried lake storage.
	//
	// example:
	//
	// -
	Data *GetLakeStorageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// SUCCESS
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

func (s GetLakeStorageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetLakeStorageResponseBody) GoString() string {
	return s.String()
}

func (s *GetLakeStorageResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetLakeStorageResponseBody) GetData() *GetLakeStorageResponseBodyData {
	return s.Data
}

func (s *GetLakeStorageResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetLakeStorageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetLakeStorageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetLakeStorageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetLakeStorageResponseBody) SetCode(v string) *GetLakeStorageResponseBody {
	s.Code = &v
	return s
}

func (s *GetLakeStorageResponseBody) SetData(v *GetLakeStorageResponseBodyData) *GetLakeStorageResponseBody {
	s.Data = v
	return s
}

func (s *GetLakeStorageResponseBody) SetHttpStatusCode(v int32) *GetLakeStorageResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetLakeStorageResponseBody) SetMessage(v string) *GetLakeStorageResponseBody {
	s.Message = &v
	return s
}

func (s *GetLakeStorageResponseBody) SetRequestId(v string) *GetLakeStorageResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLakeStorageResponseBody) SetSuccess(v bool) *GetLakeStorageResponseBody {
	s.Success = &v
	return s
}

func (s *GetLakeStorageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetLakeStorageResponseBodyData struct {
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
	// The ID of the AnalyticDB for MySQL cluster.
	//
	// example:
	//
	// amv-*******
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The description of the lake storage.
	//
	// example:
	//
	// a test db
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total storage size.
	//
	// example:
	//
	// 1
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
	// -
	PartitionCount *string `json:"PartitionCount,omitempty" xml:"PartitionCount,omitempty"`
	// The permissions on the lake storage.
	//
	// example:
	//
	// -
	Permissions []*GetLakeStorageResponseBodyDataPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
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
	// -
	RowCount *int64 `json:"RowCount,omitempty" xml:"RowCount,omitempty"`
	// The number of the tables.
	//
	// example:
	//
	// -
	TableCount *int32 `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	// The time when the lake storage was last updated.
	//
	// example:
	//
	// 2024-10-14T02:28:41Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetLakeStorageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetLakeStorageResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetLakeStorageResponseBodyData) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetLakeStorageResponseBodyData) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *GetLakeStorageResponseBodyData) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *GetLakeStorageResponseBodyData) GetDescription() *string {
	return s.Description
}

func (s *GetLakeStorageResponseBodyData) GetFileSize() *string {
	return s.FileSize
}

func (s *GetLakeStorageResponseBodyData) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *GetLakeStorageResponseBodyData) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *GetLakeStorageResponseBodyData) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *GetLakeStorageResponseBodyData) GetPartitionCount() *string {
	return s.PartitionCount
}

func (s *GetLakeStorageResponseBodyData) GetPermissions() []*GetLakeStorageResponseBodyDataPermissions {
	return s.Permissions
}

func (s *GetLakeStorageResponseBodyData) GetRegionId() *string {
	return s.RegionId
}

func (s *GetLakeStorageResponseBodyData) GetRowCount() *int64 {
	return s.RowCount
}

func (s *GetLakeStorageResponseBodyData) GetTableCount() *int32 {
	return s.TableCount
}

func (s *GetLakeStorageResponseBodyData) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *GetLakeStorageResponseBodyData) SetCreateTime(v string) *GetLakeStorageResponseBodyData {
	s.CreateTime = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetCreatorUid(v string) *GetLakeStorageResponseBodyData {
	s.CreatorUid = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetDBClusterId(v string) *GetLakeStorageResponseBodyData {
	s.DBClusterId = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetDescription(v string) *GetLakeStorageResponseBodyData {
	s.Description = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetFileSize(v string) *GetLakeStorageResponseBodyData {
	s.FileSize = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetLakeStorageId(v string) *GetLakeStorageResponseBodyData {
	s.LakeStorageId = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetOperatorUid(v string) *GetLakeStorageResponseBodyData {
	s.OperatorUid = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetOwnerUid(v string) *GetLakeStorageResponseBodyData {
	s.OwnerUid = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetPartitionCount(v string) *GetLakeStorageResponseBodyData {
	s.PartitionCount = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetPermissions(v []*GetLakeStorageResponseBodyDataPermissions) *GetLakeStorageResponseBodyData {
	s.Permissions = v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetRegionId(v string) *GetLakeStorageResponseBodyData {
	s.RegionId = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetRowCount(v int64) *GetLakeStorageResponseBodyData {
	s.RowCount = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetTableCount(v int32) *GetLakeStorageResponseBodyData {
	s.TableCount = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) SetUpdateTime(v string) *GetLakeStorageResponseBodyData {
	s.UpdateTime = &v
	return s
}

func (s *GetLakeStorageResponseBodyData) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetLakeStorageResponseBodyDataPermissions struct {
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
	// -
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write permissions.
	//
	// example:
	//
	// true
	Write *bool `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s GetLakeStorageResponseBodyDataPermissions) String() string {
	return dara.Prettify(s)
}

func (s GetLakeStorageResponseBodyDataPermissions) GoString() string {
	return s.String()
}

func (s *GetLakeStorageResponseBodyDataPermissions) GetAccount() *string {
	return s.Account
}

func (s *GetLakeStorageResponseBodyDataPermissions) GetRead() *bool {
	return s.Read
}

func (s *GetLakeStorageResponseBodyDataPermissions) GetType() *string {
	return s.Type
}

func (s *GetLakeStorageResponseBodyDataPermissions) GetWrite() *bool {
	return s.Write
}

func (s *GetLakeStorageResponseBodyDataPermissions) SetAccount(v string) *GetLakeStorageResponseBodyDataPermissions {
	s.Account = &v
	return s
}

func (s *GetLakeStorageResponseBodyDataPermissions) SetRead(v bool) *GetLakeStorageResponseBodyDataPermissions {
	s.Read = &v
	return s
}

func (s *GetLakeStorageResponseBodyDataPermissions) SetType(v string) *GetLakeStorageResponseBodyDataPermissions {
	s.Type = &v
	return s
}

func (s *GetLakeStorageResponseBodyDataPermissions) SetWrite(v bool) *GetLakeStorageResponseBodyDataPermissions {
	s.Write = &v
	return s
}

func (s *GetLakeStorageResponseBodyDataPermissions) Validate() error {
	return dara.Validate(s)
}

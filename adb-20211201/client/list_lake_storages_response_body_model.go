// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLakeStoragesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListLakeStoragesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *ListLakeStoragesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*ListLakeStoragesResponseBodyItems) *ListLakeStoragesResponseBody
	GetItems() []*ListLakeStoragesResponseBodyItems
	SetMessage(v string) *ListLakeStoragesResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *ListLakeStoragesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListLakeStoragesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListLakeStoragesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListLakeStoragesResponseBody
	GetSuccess() *bool
	SetTotalCount(v int32) *ListLakeStoragesResponseBody
	GetTotalCount() *int32
}

type ListLakeStoragesResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response code. The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The queried lake storages.
	//
	// example:
	//
	// -
	Items []*ListLakeStoragesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
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
	// The token that is used for paging when the number of results is greater than the value of MaxResults.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ****-964D-****-9D31-****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the dry run succeeds. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListLakeStoragesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListLakeStoragesResponseBody) GoString() string {
	return s.String()
}

func (s *ListLakeStoragesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListLakeStoragesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListLakeStoragesResponseBody) GetItems() []*ListLakeStoragesResponseBodyItems {
	return s.Items
}

func (s *ListLakeStoragesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListLakeStoragesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListLakeStoragesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListLakeStoragesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListLakeStoragesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListLakeStoragesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListLakeStoragesResponseBody) SetCode(v string) *ListLakeStoragesResponseBody {
	s.Code = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetHttpStatusCode(v int32) *ListLakeStoragesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetItems(v []*ListLakeStoragesResponseBodyItems) *ListLakeStoragesResponseBody {
	s.Items = v
	return s
}

func (s *ListLakeStoragesResponseBody) SetMessage(v string) *ListLakeStoragesResponseBody {
	s.Message = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetPageNumber(v int32) *ListLakeStoragesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetPageSize(v int32) *ListLakeStoragesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetRequestId(v string) *ListLakeStoragesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetSuccess(v bool) *ListLakeStoragesResponseBody {
	s.Success = &v
	return s
}

func (s *ListLakeStoragesResponseBody) SetTotalCount(v int32) *ListLakeStoragesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListLakeStoragesResponseBody) Validate() error {
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

type ListLakeStoragesResponseBodyItems struct {
	// The time when the lake storage was created.
	//
	// example:
	//
	// 2021-04-01T09:50:18Z
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
	// Create a role to run ROS StackGroups.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The size of data files.
	//
	// example:
	//
	// 651
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
	// The queried lake storage.
	//
	// example:
	//
	// 123456
	OwnerUid *string `json:"OwnerUid,omitempty" xml:"OwnerUid,omitempty"`
	// The permissions on the lake storage.
	//
	// example:
	//
	// -
	Permissions []*ListLakeStoragesResponseBodyItemsPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The number of tables.
	//
	// example:
	//
	// 30
	TableCount *int32 `json:"TableCount,omitempty" xml:"TableCount,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 30
	TotalRows *int64 `json:"TotalRows,omitempty" xml:"TotalRows,omitempty"`
	// The total storage size.
	//
	// example:
	//
	// 111333
	TotalStorage *string `json:"TotalStorage,omitempty" xml:"TotalStorage,omitempty"`
	// The time when the lake storage was last updated.
	//
	// example:
	//
	// 2024-03-15T02:24:32
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s ListLakeStoragesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListLakeStoragesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListLakeStoragesResponseBodyItems) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListLakeStoragesResponseBodyItems) GetCreatorUid() *string {
	return s.CreatorUid
}

func (s *ListLakeStoragesResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ListLakeStoragesResponseBodyItems) GetDescription() *string {
	return s.Description
}

func (s *ListLakeStoragesResponseBodyItems) GetFileSize() *string {
	return s.FileSize
}

func (s *ListLakeStoragesResponseBodyItems) GetLakeStorageId() *string {
	return s.LakeStorageId
}

func (s *ListLakeStoragesResponseBodyItems) GetOperatorUid() *string {
	return s.OperatorUid
}

func (s *ListLakeStoragesResponseBodyItems) GetOwnerUid() *string {
	return s.OwnerUid
}

func (s *ListLakeStoragesResponseBodyItems) GetPermissions() []*ListLakeStoragesResponseBodyItemsPermissions {
	return s.Permissions
}

func (s *ListLakeStoragesResponseBodyItems) GetRegionId() *string {
	return s.RegionId
}

func (s *ListLakeStoragesResponseBodyItems) GetTableCount() *int32 {
	return s.TableCount
}

func (s *ListLakeStoragesResponseBodyItems) GetTotalRows() *int64 {
	return s.TotalRows
}

func (s *ListLakeStoragesResponseBodyItems) GetTotalStorage() *string {
	return s.TotalStorage
}

func (s *ListLakeStoragesResponseBodyItems) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListLakeStoragesResponseBodyItems) SetCreateTime(v string) *ListLakeStoragesResponseBodyItems {
	s.CreateTime = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetCreatorUid(v string) *ListLakeStoragesResponseBodyItems {
	s.CreatorUid = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetDBClusterId(v string) *ListLakeStoragesResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetDescription(v string) *ListLakeStoragesResponseBodyItems {
	s.Description = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetFileSize(v string) *ListLakeStoragesResponseBodyItems {
	s.FileSize = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetLakeStorageId(v string) *ListLakeStoragesResponseBodyItems {
	s.LakeStorageId = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetOperatorUid(v string) *ListLakeStoragesResponseBodyItems {
	s.OperatorUid = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetOwnerUid(v string) *ListLakeStoragesResponseBodyItems {
	s.OwnerUid = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetPermissions(v []*ListLakeStoragesResponseBodyItemsPermissions) *ListLakeStoragesResponseBodyItems {
	s.Permissions = v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetRegionId(v string) *ListLakeStoragesResponseBodyItems {
	s.RegionId = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetTableCount(v int32) *ListLakeStoragesResponseBodyItems {
	s.TableCount = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetTotalRows(v int64) *ListLakeStoragesResponseBodyItems {
	s.TotalRows = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetTotalStorage(v string) *ListLakeStoragesResponseBodyItems {
	s.TotalStorage = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) SetUpdateTime(v string) *ListLakeStoragesResponseBodyItems {
	s.UpdateTime = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItems) Validate() error {
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

type ListLakeStoragesResponseBodyItemsPermissions struct {
	// The database account ID.
	//
	// example:
	//
	// -
	Account *string `json:"Account,omitempty" xml:"Account,omitempty"`
	// The read permissions.
	//
	// example:
	//
	// true
	Read *bool `json:"Read,omitempty" xml:"Read,omitempty"`
	// The type of the database account.
	//
	// example:
	//
	// -
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The write permissions.
	//
	// example:
	//
	// false
	Write *bool `json:"Write,omitempty" xml:"Write,omitempty"`
}

func (s ListLakeStoragesResponseBodyItemsPermissions) String() string {
	return dara.Prettify(s)
}

func (s ListLakeStoragesResponseBodyItemsPermissions) GoString() string {
	return s.String()
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) GetAccount() *string {
	return s.Account
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) GetRead() *bool {
	return s.Read
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) GetType() *string {
	return s.Type
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) GetWrite() *bool {
	return s.Write
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) SetAccount(v string) *ListLakeStoragesResponseBodyItemsPermissions {
	s.Account = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) SetRead(v bool) *ListLakeStoragesResponseBodyItemsPermissions {
	s.Read = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) SetType(v string) *ListLakeStoragesResponseBodyItemsPermissions {
	s.Type = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) SetWrite(v bool) *ListLakeStoragesResponseBodyItemsPermissions {
	s.Write = &v
	return s
}

func (s *ListLakeStoragesResponseBodyItemsPermissions) Validate() error {
	return dara.Validate(s)
}

// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *GetStorageListResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetStorageListResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetStorageListResponseBody
	GetRequestId() *string
	SetStorageList(v *GetStorageListResponseBodyStorageList) *GetStorageListResponseBody
	GetStorageList() *GetStorageListResponseBodyStorageList
	SetTotalCount(v int32) *GetStorageListResponseBody
	GetTotalCount() *int32
}

type GetStorageListResponseBody struct {
	PageNumber  *int32                                 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageList *GetStorageListResponseBodyStorageList `json:"StorageList,omitempty" xml:"StorageList,omitempty" type:"Struct"`
	TotalCount  *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s GetStorageListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageListResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetStorageListResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetStorageListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageListResponseBody) GetStorageList() *GetStorageListResponseBodyStorageList {
	return s.StorageList
}

func (s *GetStorageListResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *GetStorageListResponseBody) SetPageNumber(v int32) *GetStorageListResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetStorageListResponseBody) SetPageSize(v int32) *GetStorageListResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetStorageListResponseBody) SetRequestId(v string) *GetStorageListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageListResponseBody) SetStorageList(v *GetStorageListResponseBodyStorageList) *GetStorageListResponseBody {
	s.StorageList = v
	return s
}

func (s *GetStorageListResponseBody) SetTotalCount(v int32) *GetStorageListResponseBody {
	s.TotalCount = &v
	return s
}

func (s *GetStorageListResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStorageListResponseBodyStorageList struct {
	Storage []*GetStorageListResponseBodyStorageListStorage `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Repeated"`
}

func (s GetStorageListResponseBodyStorageList) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListResponseBodyStorageList) GoString() string {
	return s.String()
}

func (s *GetStorageListResponseBodyStorageList) GetStorage() []*GetStorageListResponseBodyStorageListStorage {
	return s.Storage
}

func (s *GetStorageListResponseBodyStorageList) SetStorage(v []*GetStorageListResponseBodyStorageListStorage) *GetStorageListResponseBodyStorageList {
	s.Storage = v
	return s
}

func (s *GetStorageListResponseBodyStorageList) Validate() error {
	return dara.Validate(s)
}

type GetStorageListResponseBodyStorageListStorage struct {
	AppId           *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	DefaultUpload   *bool   `json:"DefaultUpload,omitempty" xml:"DefaultUpload,omitempty"`
	GmtCreate       *string `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	GmtModified     *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	GroupId         *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Region          *string `json:"Region,omitempty" xml:"Region,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *int32  `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageUsage    *int32  `json:"StorageUsage,omitempty" xml:"StorageUsage,omitempty"`
	Type            *int32  `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetStorageListResponseBodyStorageListStorage) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListResponseBodyStorageListStorage) GoString() string {
	return s.String()
}

func (s *GetStorageListResponseBodyStorageListStorage) GetAppId() *string {
	return s.AppId
}

func (s *GetStorageListResponseBodyStorageListStorage) GetDefaultUpload() *bool {
	return s.DefaultUpload
}

func (s *GetStorageListResponseBodyStorageListStorage) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetStorageListResponseBodyStorageListStorage) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetStorageListResponseBodyStorageListStorage) GetGroupId() *string {
	return s.GroupId
}

func (s *GetStorageListResponseBodyStorageListStorage) GetLocation() *string {
	return s.Location
}

func (s *GetStorageListResponseBodyStorageListStorage) GetRegion() *string {
	return s.Region
}

func (s *GetStorageListResponseBodyStorageListStorage) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetStorageListResponseBodyStorageListStorage) GetStatus() *int32 {
	return s.Status
}

func (s *GetStorageListResponseBodyStorageListStorage) GetStorageUsage() *int32 {
	return s.StorageUsage
}

func (s *GetStorageListResponseBodyStorageListStorage) GetType() *int32 {
	return s.Type
}

func (s *GetStorageListResponseBodyStorageListStorage) SetAppId(v string) *GetStorageListResponseBodyStorageListStorage {
	s.AppId = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetDefaultUpload(v bool) *GetStorageListResponseBodyStorageListStorage {
	s.DefaultUpload = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetGmtCreate(v string) *GetStorageListResponseBodyStorageListStorage {
	s.GmtCreate = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetGmtModified(v string) *GetStorageListResponseBodyStorageListStorage {
	s.GmtModified = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetGroupId(v string) *GetStorageListResponseBodyStorageListStorage {
	s.GroupId = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetLocation(v string) *GetStorageListResponseBodyStorageListStorage {
	s.Location = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetRegion(v string) *GetStorageListResponseBodyStorageListStorage {
	s.Region = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetResourceGroupId(v string) *GetStorageListResponseBodyStorageListStorage {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetStatus(v int32) *GetStorageListResponseBodyStorageListStorage {
	s.Status = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetStorageUsage(v int32) *GetStorageListResponseBodyStorageListStorage {
	s.StorageUsage = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) SetType(v int32) *GetStorageListResponseBodyStorageListStorage {
	s.Type = &v
	return s
}

func (s *GetStorageListResponseBodyStorageListStorage) Validate() error {
	return dara.Validate(s)
}

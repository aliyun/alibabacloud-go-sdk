// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDomainInfoList(v *GetStorageInfoResponseBodyDomainInfoList) *GetStorageInfoResponseBody
	GetDomainInfoList() *GetStorageInfoResponseBodyDomainInfoList
	SetRequestId(v string) *GetStorageInfoResponseBody
	GetRequestId() *string
	SetStorage(v *GetStorageInfoResponseBodyStorage) *GetStorageInfoResponseBody
	GetStorage() *GetStorageInfoResponseBodyStorage
	SetStorageACL(v string) *GetStorageInfoResponseBody
	GetStorageACL() *string
}

type GetStorageInfoResponseBody struct {
	DomainInfoList *GetStorageInfoResponseBodyDomainInfoList `json:"DomainInfoList,omitempty" xml:"DomainInfoList,omitempty" type:"Struct"`
	RequestId      *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Storage        *GetStorageInfoResponseBodyStorage        `json:"Storage,omitempty" xml:"Storage,omitempty" type:"Struct"`
	StorageACL     *string                                   `json:"StorageACL,omitempty" xml:"StorageACL,omitempty"`
}

func (s GetStorageInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetStorageInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetStorageInfoResponseBody) GetDomainInfoList() *GetStorageInfoResponseBodyDomainInfoList {
	return s.DomainInfoList
}

func (s *GetStorageInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetStorageInfoResponseBody) GetStorage() *GetStorageInfoResponseBodyStorage {
	return s.Storage
}

func (s *GetStorageInfoResponseBody) GetStorageACL() *string {
	return s.StorageACL
}

func (s *GetStorageInfoResponseBody) SetDomainInfoList(v *GetStorageInfoResponseBodyDomainInfoList) *GetStorageInfoResponseBody {
	s.DomainInfoList = v
	return s
}

func (s *GetStorageInfoResponseBody) SetRequestId(v string) *GetStorageInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetStorageInfoResponseBody) SetStorage(v *GetStorageInfoResponseBodyStorage) *GetStorageInfoResponseBody {
	s.Storage = v
	return s
}

func (s *GetStorageInfoResponseBody) SetStorageACL(v string) *GetStorageInfoResponseBody {
	s.StorageACL = &v
	return s
}

func (s *GetStorageInfoResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetStorageInfoResponseBodyDomainInfoList struct {
	DomainInfo []*GetStorageInfoResponseBodyDomainInfoListDomainInfo `json:"DomainInfo,omitempty" xml:"DomainInfo,omitempty" type:"Repeated"`
}

func (s GetStorageInfoResponseBodyDomainInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetStorageInfoResponseBodyDomainInfoList) GoString() string {
	return s.String()
}

func (s *GetStorageInfoResponseBodyDomainInfoList) GetDomainInfo() []*GetStorageInfoResponseBodyDomainInfoListDomainInfo {
	return s.DomainInfo
}

func (s *GetStorageInfoResponseBodyDomainInfoList) SetDomainInfo(v []*GetStorageInfoResponseBodyDomainInfoListDomainInfo) *GetStorageInfoResponseBodyDomainInfoList {
	s.DomainInfo = v
	return s
}

func (s *GetStorageInfoResponseBodyDomainInfoList) Validate() error {
	return dara.Validate(s)
}

type GetStorageInfoResponseBodyDomainInfoListDomainInfo struct {
	DefaultPlay  *bool   `json:"DefaultPlay,omitempty" xml:"DefaultPlay,omitempty"`
	DomainCname  *string `json:"DomainCname,omitempty" xml:"DomainCname,omitempty"`
	DomainName   *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	DomainStatus *string `json:"DomainStatus,omitempty" xml:"DomainStatus,omitempty"`
}

func (s GetStorageInfoResponseBodyDomainInfoListDomainInfo) String() string {
	return dara.Prettify(s)
}

func (s GetStorageInfoResponseBodyDomainInfoListDomainInfo) GoString() string {
	return s.String()
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) GetDefaultPlay() *bool {
	return s.DefaultPlay
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) GetDomainCname() *string {
	return s.DomainCname
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) GetDomainName() *string {
	return s.DomainName
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) GetDomainStatus() *string {
	return s.DomainStatus
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) SetDefaultPlay(v bool) *GetStorageInfoResponseBodyDomainInfoListDomainInfo {
	s.DefaultPlay = &v
	return s
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) SetDomainCname(v string) *GetStorageInfoResponseBodyDomainInfoListDomainInfo {
	s.DomainCname = &v
	return s
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) SetDomainName(v string) *GetStorageInfoResponseBodyDomainInfoListDomainInfo {
	s.DomainName = &v
	return s
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) SetDomainStatus(v string) *GetStorageInfoResponseBodyDomainInfoListDomainInfo {
	s.DomainStatus = &v
	return s
}

func (s *GetStorageInfoResponseBodyDomainInfoListDomainInfo) Validate() error {
	return dara.Validate(s)
}

type GetStorageInfoResponseBodyStorage struct {
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

func (s GetStorageInfoResponseBodyStorage) String() string {
	return dara.Prettify(s)
}

func (s GetStorageInfoResponseBodyStorage) GoString() string {
	return s.String()
}

func (s *GetStorageInfoResponseBodyStorage) GetDefaultUpload() *bool {
	return s.DefaultUpload
}

func (s *GetStorageInfoResponseBodyStorage) GetGmtCreate() *string {
	return s.GmtCreate
}

func (s *GetStorageInfoResponseBodyStorage) GetGmtModified() *string {
	return s.GmtModified
}

func (s *GetStorageInfoResponseBodyStorage) GetGroupId() *string {
	return s.GroupId
}

func (s *GetStorageInfoResponseBodyStorage) GetLocation() *string {
	return s.Location
}

func (s *GetStorageInfoResponseBodyStorage) GetRegion() *string {
	return s.Region
}

func (s *GetStorageInfoResponseBodyStorage) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *GetStorageInfoResponseBodyStorage) GetStatus() *int32 {
	return s.Status
}

func (s *GetStorageInfoResponseBodyStorage) GetStorageUsage() *int32 {
	return s.StorageUsage
}

func (s *GetStorageInfoResponseBodyStorage) GetType() *int32 {
	return s.Type
}

func (s *GetStorageInfoResponseBodyStorage) SetDefaultUpload(v bool) *GetStorageInfoResponseBodyStorage {
	s.DefaultUpload = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetGmtCreate(v string) *GetStorageInfoResponseBodyStorage {
	s.GmtCreate = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetGmtModified(v string) *GetStorageInfoResponseBodyStorage {
	s.GmtModified = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetGroupId(v string) *GetStorageInfoResponseBodyStorage {
	s.GroupId = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetLocation(v string) *GetStorageInfoResponseBodyStorage {
	s.Location = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetRegion(v string) *GetStorageInfoResponseBodyStorage {
	s.Region = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetResourceGroupId(v string) *GetStorageInfoResponseBodyStorage {
	s.ResourceGroupId = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetStatus(v int32) *GetStorageInfoResponseBodyStorage {
	s.Status = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetStorageUsage(v int32) *GetStorageInfoResponseBodyStorage {
	s.StorageUsage = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) SetType(v int32) *GetStorageInfoResponseBodyStorage {
	s.Type = &v
	return s
}

func (s *GetStorageInfoResponseBodyStorage) Validate() error {
	return dara.Validate(s)
}

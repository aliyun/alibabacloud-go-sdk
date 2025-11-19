// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAttachedMediaInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAttachedMediaList(v []*GetAttachedMediaInfoResponseBodyAttachedMediaList) *GetAttachedMediaInfoResponseBody
	GetAttachedMediaList() []*GetAttachedMediaInfoResponseBodyAttachedMediaList
	SetNonExistMediaIds(v []*string) *GetAttachedMediaInfoResponseBody
	GetNonExistMediaIds() []*string
	SetRequestId(v string) *GetAttachedMediaInfoResponseBody
	GetRequestId() *string
}

type GetAttachedMediaInfoResponseBody struct {
	// The information about the media assets.
	AttachedMediaList []*GetAttachedMediaInfoResponseBodyAttachedMediaList `json:"AttachedMediaList,omitempty" xml:"AttachedMediaList,omitempty" type:"Repeated"`
	// The IDs of the auxiliary media assets that do not exist.
	NonExistMediaIds []*string `json:"NonExistMediaIds,omitempty" xml:"NonExistMediaIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// 221BCB57-B217-42BF-619BD13378F9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAttachedMediaInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAttachedMediaInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponseBody) GetAttachedMediaList() []*GetAttachedMediaInfoResponseBodyAttachedMediaList {
	return s.AttachedMediaList
}

func (s *GetAttachedMediaInfoResponseBody) GetNonExistMediaIds() []*string {
	return s.NonExistMediaIds
}

func (s *GetAttachedMediaInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAttachedMediaInfoResponseBody) SetAttachedMediaList(v []*GetAttachedMediaInfoResponseBodyAttachedMediaList) *GetAttachedMediaInfoResponseBody {
	s.AttachedMediaList = v
	return s
}

func (s *GetAttachedMediaInfoResponseBody) SetNonExistMediaIds(v []*string) *GetAttachedMediaInfoResponseBody {
	s.NonExistMediaIds = v
	return s
}

func (s *GetAttachedMediaInfoResponseBody) SetRequestId(v string) *GetAttachedMediaInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBody) Validate() error {
	if s.AttachedMediaList != nil {
		for _, item := range s.AttachedMediaList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAttachedMediaInfoResponseBodyAttachedMediaList struct {
	// The ID of the application.
	//
	// example:
	//
	// app-*****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The categories.
	Categories []*GetAttachedMediaInfoResponseBodyAttachedMediaListCategories `json:"Categories,omitempty" xml:"Categories,omitempty" type:"Repeated"`
	// The time when the auxiliary media asset was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-01T10:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the auxiliary media asset.
	//
	// >  This parameter is returned only when a description is specified for the auxiliary media asset.
	//
	// example:
	//
	// description test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the auxiliary media asset.
	//
	// example:
	//
	// 0222e203cf80f9c22870a4d2c****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The time when the auxiliary media asset was last updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2020-05-31T11:42:20Z
	ModificationTime *string `json:"ModificationTime,omitempty" xml:"ModificationTime,omitempty"`
	// The status of the auxiliary media asset. Valid values:
	//
	// 	- **Uploading**
	//
	// 	- **Normal**
	//
	// 	- **UploadFail**
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The storage address of the auxiliary media asset.
	//
	// example:
	//
	// outin-bfefbb9*****c7426.oss-cn-shanghai.aliyuncs.com
	StorageLocation *string `json:"StorageLocation,omitempty" xml:"StorageLocation,omitempty"`
	// The tags of the auxiliary media asset.
	//
	// >  This parameter is returned only when tags are specified for the auxiliary media asset.
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the auxiliary media asset.
	//
	// example:
	//
	// subtitle test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The type of the auxiliary media asset.
	//
	// 	- **watermark**
	//
	// 	- **subtitle**
	//
	// 	- **material**
	//
	// example:
	//
	// subtitle
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The URL of the auxiliary media asset.
	//
	// >  If a CDN domain name is specified, a CDN URL is returned. Otherwise, an OSS URL is returned.
	//
	// example:
	//
	// https://al*****.cn/subtitle/9843C2*****4E186F19B6.vtt?auth_key=159099f60e0b7fd59****
	URL *string `json:"URL,omitempty" xml:"URL,omitempty"`
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaList) String() string {
	return dara.Prettify(s)
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaList) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetAppId() *string {
	return s.AppId
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetCategories() []*GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	return s.Categories
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetDescription() *string {
	return s.Description
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetMediaId() *string {
	return s.MediaId
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetModificationTime() *string {
	return s.ModificationTime
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetStatus() *string {
	return s.Status
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetStorageLocation() *string {
	return s.StorageLocation
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetTags() *string {
	return s.Tags
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetTitle() *string {
	return s.Title
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetType() *string {
	return s.Type
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) GetURL() *string {
	return s.URL
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetAppId(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.AppId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetCategories(v []*GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Categories = v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetCreationTime(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.CreationTime = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetDescription(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Description = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetMediaId(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.MediaId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetModificationTime(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.ModificationTime = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetStatus(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Status = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetStorageLocation(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.StorageLocation = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetTags(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Tags = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetTitle(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Title = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetType(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.Type = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) SetURL(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaList {
	s.URL = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaList) Validate() error {
	if s.Categories != nil {
		for _, item := range s.Categories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAttachedMediaInfoResponseBodyAttachedMediaListCategories struct {
	// The ID of the category.
	//
	// example:
	//
	// 1000224338
	CateId *int64 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the category.
	//
	// example:
	//
	// category test
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The level of the category.
	//
	// example:
	//
	// 0
	Level *int64 `json:"Level,omitempty" xml:"Level,omitempty"`
	// The ID of the parent category.
	//
	// example:
	//
	// 1000224336
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) String() string {
	return dara.Prettify(s)
}

func (s GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) GoString() string {
	return s.String()
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) GetCateId() *int64 {
	return s.CateId
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) GetCateName() *string {
	return s.CateName
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) GetLevel() *int64 {
	return s.Level
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) GetParentId() *int64 {
	return s.ParentId
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetCateId(v int64) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.CateId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetCateName(v string) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.CateName = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetLevel(v int64) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.Level = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) SetParentId(v int64) *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories {
	s.ParentId = &v
	return s
}

func (s *GetAttachedMediaInfoResponseBodyAttachedMediaListCategories) Validate() error {
	return dara.Validate(s)
}

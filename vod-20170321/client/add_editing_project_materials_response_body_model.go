// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddEditingProjectMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialList(v []*AddEditingProjectMaterialsResponseBodyMaterialList) *AddEditingProjectMaterialsResponseBody
	GetMaterialList() []*AddEditingProjectMaterialsResponseBodyMaterialList
	SetRequestId(v string) *AddEditingProjectMaterialsResponseBody
	GetRequestId() *string
}

type AddEditingProjectMaterialsResponseBody struct {
	// The list of media assets.
	MaterialList []*AddEditingProjectMaterialsResponseBodyMaterialList `json:"MaterialList,omitempty" xml:"MaterialList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 85237CDA-0B54-5CED-BA10-A8A71AA13C1A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBody) GetMaterialList() []*AddEditingProjectMaterialsResponseBodyMaterialList {
	return s.MaterialList
}

func (s *AddEditingProjectMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddEditingProjectMaterialsResponseBody) SetMaterialList(v []*AddEditingProjectMaterialsResponseBodyMaterialList) *AddEditingProjectMaterialsResponseBody {
	s.MaterialList = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) SetRequestId(v string) *AddEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBody) Validate() error {
	if s.MaterialList != nil {
		for _, item := range s.MaterialList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddEditingProjectMaterialsResponseBodyMaterialList struct {
	// The category ID.
	//
	// example:
	//
	// 1000487543
	CateId *int32 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The name of the media asset category.
	//
	// example:
	//
	// Category 1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The URL of the media asset thumbnail.
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the media asset was created. The time follows the <i>yyyy-mm-dd</i>T<i>hh:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-10-02T08:26Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1234751840694470
	CustomerId *int64 `json:"CustomerId,omitempty" xml:"CustomerId,omitempty"`
	// The description of the media asset.
	//
	// example:
	//
	// Media asset description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the media asset. Unit: seconds. The value is accurate to four decimal places.
	//
	// example:
	//
	// 3690.2332
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// 85befc4118b84c6723e53b80b1****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	// The type of the media asset. Valid values:
	//
	// - **video**: video
	//
	// - **audio**: audio-only
	//
	// - **image**: image
	//
	// example:
	//
	// video
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// The time when the media asset was updated. The time follows the <i>yyyy-mm-dd</i>T<i>hh:mm:ss</i>Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2022-11-25T07:28:34Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The file size.
	//
	// example:
	//
	// 1682694
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The array of snapshot URLs for the media asset.
	Snapshots []*string `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Repeated"`
	// The sprite configuration.
	//
	// example:
	//
	// xxx
	SpriteConfig *string `json:"SpriteConfig,omitempty" xml:"SpriteConfig,omitempty"`
	// The array of sprite URLs for the media asset.
	Sprites []*string `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Repeated"`
	// The status of the media asset. Valid values:
	//
	// - **Normal**: Normal.
	//
	// - **Uploading**: Uploading.
	//
	// - **UploadFail**: Upload failed.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the media asset. Separate multiple tags with commas (,).
	//
	// example:
	//
	// tag1,tag2
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the media asset.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s AddEditingProjectMaterialsResponseBodyMaterialList) String() string {
	return dara.Prettify(s)
}

func (s AddEditingProjectMaterialsResponseBodyMaterialList) GoString() string {
	return s.String()
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetCateId() *int32 {
	return s.CateId
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetCateName() *string {
	return s.CateName
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetCoverURL() *string {
	return s.CoverURL
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetCreateTime() *string {
	return s.CreateTime
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetCustomerId() *int64 {
	return s.CustomerId
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetDescription() *string {
	return s.Description
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetDuration() *float32 {
	return s.Duration
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetMaterialId() *string {
	return s.MaterialId
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetMaterialType() *string {
	return s.MaterialType
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetSize() *int64 {
	return s.Size
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetSnapshots() []*string {
	return s.Snapshots
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetSpriteConfig() *string {
	return s.SpriteConfig
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetSprites() []*string {
	return s.Sprites
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetStatus() *string {
	return s.Status
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetTags() *string {
	return s.Tags
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) GetTitle() *string {
	return s.Title
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetCateId(v int32) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.CateId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetCateName(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.CateName = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetCoverURL(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.CoverURL = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetCreateTime(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.CreateTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetCustomerId(v int64) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.CustomerId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetDescription(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Description = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetDuration(v float32) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Duration = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetMaterialId(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.MaterialId = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetMaterialType(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.MaterialType = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetModifyTime(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.ModifyTime = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetSize(v int64) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Size = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetSnapshots(v []*string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Snapshots = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetSpriteConfig(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.SpriteConfig = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetSprites(v []*string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Sprites = v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetStatus(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Status = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetTags(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Tags = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) SetTitle(v string) *AddEditingProjectMaterialsResponseBodyMaterialList {
	s.Title = &v
	return s
}

func (s *AddEditingProjectMaterialsResponseBodyMaterialList) Validate() error {
	return dara.Validate(s)
}

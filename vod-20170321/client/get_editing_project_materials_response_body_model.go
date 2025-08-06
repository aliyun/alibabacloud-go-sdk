// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEditingProjectMaterialsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaterialList(v *GetEditingProjectMaterialsResponseBodyMaterialList) *GetEditingProjectMaterialsResponseBody
	GetMaterialList() *GetEditingProjectMaterialsResponseBodyMaterialList
	SetRequestId(v string) *GetEditingProjectMaterialsResponseBody
	GetRequestId() *string
}

type GetEditingProjectMaterialsResponseBody struct {
	// The materials.
	MaterialList *GetEditingProjectMaterialsResponseBodyMaterialList `json:"MaterialList,omitempty" xml:"MaterialList,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 746FFA07-8BBB-46B1-3E94E3B2915E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBody) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBody) GetMaterialList() *GetEditingProjectMaterialsResponseBodyMaterialList {
	return s.MaterialList
}

func (s *GetEditingProjectMaterialsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEditingProjectMaterialsResponseBody) SetMaterialList(v *GetEditingProjectMaterialsResponseBodyMaterialList) *GetEditingProjectMaterialsResponseBody {
	s.MaterialList = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) SetRequestId(v string) *GetEditingProjectMaterialsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMaterialList struct {
	Material []*GetEditingProjectMaterialsResponseBodyMaterialListMaterial `json:"Material,omitempty" xml:"Material,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialList) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialList) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialList) GetMaterial() []*GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	return s.Material
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialList) SetMaterial(v []*GetEditingProjectMaterialsResponseBodyMaterialListMaterial) *GetEditingProjectMaterialsResponseBodyMaterialList {
	s.Material = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialList) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterial struct {
	// The category ID of the material.
	//
	// example:
	//
	// 100005****
	CateId *int32 `json:"CateId,omitempty" xml:"CateId,omitempty"`
	// The category name of the material.
	//
	// example:
	//
	// test1
	CateName *string `json:"CateName,omitempty" xml:"CateName,omitempty"`
	// The thumbnail URL of the material.
	CoverURL *string `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	// The time when the material was created. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-10-19 10:48:17
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the material.
	//
	// example:
	//
	// test2
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The duration of the material. The value is rounded to four decimal places. Unit: seconds.
	//
	// example:
	//
	// 15.16
	Duration *float32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the material.
	//
	// example:
	//
	// 85befc4118b84c6723e53b80b1****
	MaterialId *string `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	// The type of the material. Valid values:
	//
	// 	- **video**
	//
	// 	- **audio**
	//
	// 	- **image**
	//
	// example:
	//
	// video
	MaterialType *string `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	// The time when the material was last updated. The time follows the ISO 8601 standard in the *yyyy-MM-dd*T*HH:mm:ss*Z format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-10-19 10:48:17
	ModifiedTime *string `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	// The size of the mezzanine file. Unit: byte.
	//
	// example:
	//
	// 1682694
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The URLs of material snapshots. The value is an array.
	Snapshots *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	// The source of the sprite.
	//
	// example:
	//
	// xxx
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The configuration of the sprite.
	//
	// example:
	//
	// xxx
	SpriteConfig *string `json:"SpriteConfig,omitempty" xml:"SpriteConfig,omitempty"`
	// The URLs of material sprites. The value is an array.
	Sprites *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Struct"`
	// The status of the material. Valid values:
	//
	// 	- **Normal**: The material is in draft.
	//
	// 	- **Producing**: The material is being produced.
	//
	// 	- **Produced**: The material was produced.
	//
	// 	- **ProduceFailed**: The material failed to be produced.
	//
	// example:
	//
	// Normal
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tag of the material. Multiple tags are separated by commas (,).
	//
	// example:
	//
	// editing test
	Tags *string `json:"Tags,omitempty" xml:"Tags,omitempty"`
	// The title of the material.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterial) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetCateId() *int32 {
	return s.CateId
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetCateName() *string {
	return s.CateName
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetCoverURL() *string {
	return s.CoverURL
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetCreationTime() *string {
	return s.CreationTime
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetDescription() *string {
	return s.Description
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetDuration() *float32 {
	return s.Duration
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetMaterialId() *string {
	return s.MaterialId
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetMaterialType() *string {
	return s.MaterialType
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetModifiedTime() *string {
	return s.ModifiedTime
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetSize() *int64 {
	return s.Size
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetSnapshots() *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots {
	return s.Snapshots
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetSource() *string {
	return s.Source
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetSpriteConfig() *string {
	return s.SpriteConfig
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetSprites() *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites {
	return s.Sprites
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetStatus() *string {
	return s.Status
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetTags() *string {
	return s.Tags
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) GetTitle() *string {
	return s.Title
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCateId(v int32) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CateId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCateName(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CateName = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCoverURL(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CoverURL = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetCreationTime(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.CreationTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetDescription(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Description = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetDuration(v float32) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Duration = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetMaterialId(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.MaterialId = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetMaterialType(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.MaterialType = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetModifiedTime(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.ModifiedTime = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSize(v int64) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Size = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSnapshots(v *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Snapshots = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSource(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Source = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSpriteConfig(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.SpriteConfig = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetSprites(v *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Sprites = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetStatus(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Status = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetTags(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Tags = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) SetTitle(v string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterial {
	s.Title = &v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterial) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots struct {
	Snapshot []*string `json:"Snapshot,omitempty" xml:"Snapshot,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) GetSnapshot() []*string {
	return s.Snapshot
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) SetSnapshot(v []*string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots {
	s.Snapshot = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots) Validate() error {
	return dara.Validate(s)
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites struct {
	Sprite []*string `json:"Sprite,omitempty" xml:"Sprite,omitempty" type:"Repeated"`
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) String() string {
	return dara.Prettify(s)
}

func (s GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) GoString() string {
	return s.String()
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) GetSprite() []*string {
	return s.Sprite
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) SetSprite(v []*string) *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites {
	s.Sprite = v
	return s
}

func (s *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites) Validate() error {
	return dara.Validate(s)
}

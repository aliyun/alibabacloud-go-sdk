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
	if s.MaterialList != nil {
		if err := s.MaterialList.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.Material != nil {
		for _, item := range s.Material {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetEditingProjectMaterialsResponseBodyMaterialListMaterial struct {
	CateId       *int32                                                               `json:"CateId,omitempty" xml:"CateId,omitempty"`
	CateName     *string                                                              `json:"CateName,omitempty" xml:"CateName,omitempty"`
	CoverURL     *string                                                              `json:"CoverURL,omitempty" xml:"CoverURL,omitempty"`
	CreationTime *string                                                              `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string                                                              `json:"Description,omitempty" xml:"Description,omitempty"`
	Duration     *float32                                                             `json:"Duration,omitempty" xml:"Duration,omitempty"`
	MaterialId   *string                                                              `json:"MaterialId,omitempty" xml:"MaterialId,omitempty"`
	MaterialType *string                                                              `json:"MaterialType,omitempty" xml:"MaterialType,omitempty"`
	ModifiedTime *string                                                              `json:"ModifiedTime,omitempty" xml:"ModifiedTime,omitempty"`
	Size         *int64                                                               `json:"Size,omitempty" xml:"Size,omitempty"`
	Snapshots    *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSnapshots `json:"Snapshots,omitempty" xml:"Snapshots,omitempty" type:"Struct"`
	Source       *string                                                              `json:"Source,omitempty" xml:"Source,omitempty"`
	SpriteConfig *string                                                              `json:"SpriteConfig,omitempty" xml:"SpriteConfig,omitempty"`
	Sprites      *GetEditingProjectMaterialsResponseBodyMaterialListMaterialSprites   `json:"Sprites,omitempty" xml:"Sprites,omitempty" type:"Struct"`
	Status       *string                                                              `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags         *string                                                              `json:"Tags,omitempty" xml:"Tags,omitempty"`
	Title        *string                                                              `json:"Title,omitempty" xml:"Title,omitempty"`
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
	if s.Snapshots != nil {
		if err := s.Snapshots.Validate(); err != nil {
			return err
		}
	}
	if s.Sprites != nil {
		if err := s.Sprites.Validate(); err != nil {
			return err
		}
	}
	return nil
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
